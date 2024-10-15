package grpcclient

import (
	"context"
	"encoding/json"
	"log"
	"time"

	pb "github.com/achmad-dev/internal/pkg/common/genproto"
	"github.com/achmad-dev/internal/user/api/dto"
	"github.com/achmad-dev/internal/user/internal/middleware"
	"github.com/achmad-dev/internal/user/internal/service"
	"github.com/gofiber/fiber/v2"
	"github.com/redis/go-redis/v9"
	"google.golang.org/grpc"
)

type BookClient struct {
	client         pb.BookServiceClient
	authorClient   pb.AuthorServiceClient
	categoryClient pb.CategoryServiceClient
	ubbService     service.UserBorrowedBookService
	userService    service.UserService
	redisClient    *redis.Client
}

func NewBookClient(conn *grpc.ClientConn, atConn *grpc.ClientConn, ctConn *grpc.ClientConn, ubbService service.UserBorrowedBookService, userService service.UserService, redisClient *redis.Client) *BookClient {
	return &BookClient{
		client:         pb.NewBookServiceClient(conn),
		authorClient:   pb.NewAuthorServiceClient(atConn),
		categoryClient: pb.NewCategoryServiceClient(ctConn),
		ubbService:     ubbService,
		userService:    userService,
		redisClient:    redisClient,
	}
}
func (bc *BookClient) GetBook(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	req := &pb.GetBookRequest{Id: id}
	res, err := bc.client.GetBook(context.Background(), req)
	if err != nil {
		log.Printf("Error calling GetBook: %v", err)
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Internal Server Error"})
	}
	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{"book": res.Book})
}

func (bc *BookClient) GetBookByTitle(ctx *fiber.Ctx) error {
	title := ctx.Params("title")
	req := &pb.GetBookByTitleRequest{Title: title}
	res, err := bc.client.GetBookByTitle(context.Background(), req)
	if err != nil {
		log.Printf("Error calling GetBookByTitle: %v", err)
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Internal Server Error"})
	}
	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{"book": res.Book})
}

func (bc *BookClient) GetBooksByAuthorName(ctx *fiber.Ctx) error {
	author := ctx.Params("author")
	req := &pb.GetBooksByAuthorNameRequest{Author: author}
	res, err := bc.client.GetBooksByAuthorName(context.Background(), req)
	if err != nil {
		log.Printf("Error calling GetBooksByAuthorName: %v", err)
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Internal Server Error"})
	}
	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{"books": res.Books})
}

func (bc *BookClient) GetPopularBooksByCategory(ctx *fiber.Ctx) error {
	category := ctx.Params("category")
	req := &pb.GetPopularBooksByCategoryRequest{Category: category}
	res, err := bc.client.GetPopularBooksByCategory(context.Background(), req)
	if err != nil {
		log.Printf("Error calling GetPopularBooksByCategory: %v", err)
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Internal Server Error"})
	}
	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{"books": res.Books})
}

func (bc *BookClient) BorrowBook(ctx *fiber.Ctx) error {
	var req dto.BorrowBookRequest
	if err := ctx.BodyParser(&req); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Bad Request"})
	}

	// Check if the book exist first
	book, err := bc.client.GetBook(ctx.Context(), &pb.GetBookRequest{Id: req.BookId})
	if err != nil {
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{"message": "not found"})
	}
	pbReq := &pb.BorrowBookRequest{
		Id:       req.BookId,
		Quantity: int32(req.Quantity),
	}
	_, err = bc.client.BorrowBook(context.Background(), pbReq)
	if err != nil {
		log.Printf("Error calling BorrowBook: %v", err)
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Internal Server Error"})
	}
	username := ctx.Locals("username").(string)
	user, err := bc.userService.GetUserByUsername(ctx.Context(), username)
	if err != nil {
		log.Printf("Error fetching user: %v", err)
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Internal Server Error"})
	}
	err = bc.ubbService.BorrowBook(user.ID, book.Book.Id, book.Book.Title, req.Quantity)
	if err != nil {
		log.Printf("Error updating UserBorrowedBookService: %v", err)
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Internal Server Error"})
	}

	// delete from redis cache
	err = bc.redisClient.Del(ctx.Context(), "books").Err()
	if err != nil {
		log.Printf("Error deleting data from Redis: %v", err)
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{"message": "Book borrowed successfully"})
}

func (bc *BookClient) ReturnBook(ctx *fiber.Ctx) error {
	var req dto.ReturnBookRequest
	if err := ctx.BodyParser(&req); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Bad Request"})
	}

	// Check if the book exist first
	book, err := bc.client.GetBook(ctx.Context(), &pb.GetBookRequest{Id: req.BookId})
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "Bad Request"})
	}

	pbReq := &pb.ReturnBookRequest{
		Id:       book.Book.Id,
		Quantity: int32(req.Quantity),
	}
	_, err = bc.client.ReturnBook(context.Background(), pbReq)
	if err != nil {
		log.Printf("Error calling ReturnBook: %v", err)
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Internal Server Error"})
	}
	err = bc.ubbService.ReturnBook(book.Book.Id, req.Quantity)
	if err != nil {
		log.Printf("Error updating UserBorrowedBookService: %v", err)
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Internal Server Error"})
	}

	// delete from redis cache
	err = bc.redisClient.Del(ctx.Context(), "books").Err()
	if err != nil {
		log.Printf("Error deleting data from Redis: %v", err)
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{"message": "Book returned successfully"})
}

func (bc *BookClient) CreateBook(ctx *fiber.Ctx) error {
	var req dto.CreateBookRequest
	if err := ctx.BodyParser(&req); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Bad Request"})
	}
	author, err := bc.authorClient.GetAuthor(ctx.Context(), &pb.GetAuthorRequest{Id: req.AuthorId})
	if err != nil {
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{"message": "author not found"})
	}
	category, err := bc.categoryClient.GetCategory(ctx.Context(), &pb.GetCategoryRequest{Id: req.CategoryId})
	if err != nil {
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{"message": "category not found"})
	}
	pbReq := &pb.CreateBookRequest{
		AuthorId:   author.Author.Id,
		CategoryId: category.Category.Id,
		Title:      req.Title,
		Author:     author.Author.Name,
		Category:   category.Category.Name,
		Stock:      int32(req.Stock),
	}
	_, err = bc.client.CreateBook(context.Background(), pbReq)
	if err != nil {
		log.Printf("Error calling CreateBook: %v", err)
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Internal Server Error"})
	}

	// delete from redis cache
	err = bc.redisClient.Del(ctx.Context(), "books").Err()
	if err != nil {
		log.Printf("Error deleting data from Redis: %v", err)
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{"message": "Book added succesfully"})
}

func (bc *BookClient) UpdateBook(ctx *fiber.Ctx) error {
	var req dto.UpdateBookRequest
	if err := ctx.BodyParser(&req); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Bad Request"})
	}
	pbReq := &pb.UpdateBookRequest{
		Id:        req.Id,
		Title:     req.Title,
		Author:    req.Author,
		Category:  req.Category,
		Stock:     int32(req.Stock),
		Borrowed:  int32(req.Borrowed),
		IsPopular: req.IsPopular,
	}
	_, err := bc.client.UpdateBook(context.Background(), pbReq)
	if err != nil {
		log.Printf("Error calling UpdateBook: %v", err)
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Internal Server Error"})
	}
	// delete from redis cache
	err = bc.redisClient.Del(ctx.Context(), "books").Err()
	if err != nil {
		log.Printf("Error deleting data from Redis: %v", err)
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{"message": "Book updated succesfully"})
}

func (bc *BookClient) DeleteBook(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	req := &pb.DeleteBookRequest{Id: id}
	_, err := bc.client.DeleteBook(context.Background(), req)
	if err != nil {
		log.Printf("Error calling DeleteBook: %v", err)
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Internal Server Error"})
	}
	// delete from redis cache
	err = bc.redisClient.Del(ctx.Context(), "books").Err()
	if err != nil {
		log.Printf("Error deleting data from Redis: %v", err)
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{"message": "Book deleted successfully"})
}

func (bc *BookClient) ListBooks(ctx *fiber.Ctx) error {
	// return from redis cache if available
	booksJSON, err := bc.redisClient.Get(ctx.Context(), "books").Result()
	if err == nil {
		var books []*pb.Book
		err = json.Unmarshal([]byte(booksJSON), &books)
		if err != nil {
			log.Printf("Error unmarshaling data from Redis: %v", err)
		}
		return ctx.Status(fiber.StatusOK).JSON(fiber.Map{"books": books})
	}

	req := &pb.ListBooksRequest{}
	res, err := bc.client.ListBooks(context.Background(), req)
	if err != nil {
		log.Printf("Error calling ListBooks: %v", err)
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Internal Server Error"})
	}

	// store in redis cache
	booksJSONBytes, err := json.Marshal(res.Books)
	if err != nil {
		log.Printf("Error marshaling data to JSON: %v", err)
	}
	booksJSON = string(booksJSONBytes)

	err = bc.redisClient.Set(ctx.Context(), "books", booksJSON, 2*time.Hour).Err()
	if err != nil {
		log.Printf("Error storing data in Redis: %v", err)
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{"books": res.Books})
}

// TODO: implement create book and delete book and get book by title, borrow book, return book
func RegisterBookRoutes(
	api fiber.Router,
	conn *grpc.ClientConn,
	atConn *grpc.ClientConn,
	ctConn *grpc.ClientConn,
	ubbService service.UserBorrowedBookService,
	secret string,
	userService service.UserService,
	redisClient *redis.Client,
) {

	client := NewBookClient(conn, atConn, ctConn, ubbService, userService, redisClient)

	// Book routes
	api.Get("/book/:id", middleware.AuthMiddleware(secret, userService), client.GetBook)
	api.Get("/book/title/:title", middleware.AuthMiddleware(secret, userService), client.GetBookByTitle)
	api.Get("/books/author/:author", middleware.AuthMiddleware(secret, userService), client.GetBooksByAuthorName)
	api.Get("/books/popular/:category", middleware.AuthMiddleware(secret, userService), client.GetPopularBooksByCategory)
	api.Get("/books", middleware.AuthMiddleware(secret, userService), client.ListBooks)

	// Book actions
	api.Post("/book/borrow", middleware.AuthMiddleware(secret, userService), client.BorrowBook)
	api.Post("/book/return", middleware.AuthMiddleware(secret, userService), client.ReturnBook)

	//List user borrowed books
	api.Get("/book/user/borrowed", middleware.AuthMiddleware(secret, userService), func(c *fiber.Ctx) error {
		username := c.Locals("username").(string)
		log.Printf("Fetching user borrowed books for: %s", username)
		user, err := userService.GetUserByUsername(c.Context(), username)
		if err != nil {
			log.Printf("Error fetching user: %v", err)
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"message": "Unauthorized"})
		}
		books, err := ubbService.GetBorrowedBooksByUserID(user.ID)
		if err != nil {
			log.Printf("Error fetching user borrowed books: %v", err)
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Internal Server Error"})
		}
		return c.Status(fiber.StatusOK).JSON(fiber.Map{"books": books})
	})

	// for admin
	api.Post("/book", middleware.AuthAdminMiddleware(secret, userService), client.CreateBook)
	api.Put("/book", middleware.AuthAdminMiddleware(secret, userService), client.UpdateBook)
	api.Delete("/book/:id", middleware.AuthAdminMiddleware(secret, userService), client.DeleteBook)
}
