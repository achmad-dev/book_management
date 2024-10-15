package grpcclient

import (
	"context"
	"log"

	pb "github.com/achmad-dev/internal/pkg/common/genproto"
	"github.com/achmad-dev/internal/user/internal/middleware"
	"github.com/achmad-dev/internal/user/internal/service"
	"github.com/gofiber/fiber/v2"
	"google.golang.org/grpc"
)

type AuthorClient struct {
	client pb.AuthorServiceClient
}

func NewAuthorClient(conn *grpc.ClientConn) *AuthorClient {
	return &AuthorClient{
		client: pb.NewAuthorServiceClient(conn),
	}
}

func (ac *AuthorClient) GetAuthor(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	req := &pb.GetAuthorRequest{Id: id}
	res, err := ac.client.GetAuthor(context.Background(), req)
	if err != nil {
		log.Printf("Error calling GetAuthor: %v", err)
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Internal Server Error"})
	}
	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{"author": res.Author})
}

func (ac *AuthorClient) GetAuthorByName(ctx *fiber.Ctx) error {
	name := ctx.Params("name")
	req := &pb.GetAuthorByNameRequest{Name: name}
	res, err := ac.client.GetAuthorByName(context.Background(), req)
	if err != nil {
		log.Printf("Error calling GetAuthorByName: %v", err)
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Internal Server Error"})
	}
	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{"author": res.Author})
}

func (ac *AuthorClient) CreateAuthor(ctx *fiber.Ctx) error {
	var req pb.CreateAuthorRequest
	if err := ctx.BodyParser(&req); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Bad Request"})
	}
	res, err := ac.client.CreateAuthor(context.Background(), &req)
	if err != nil {
		log.Printf("Error calling CreateAuthor: %v", err)
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Internal Server Error"})
	}
	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{"author": res.Author})
}

func (ac *AuthorClient) UpdateAuthor(ctx *fiber.Ctx) error {
	var req pb.UpdateAuthorRequest
	if err := ctx.BodyParser(&req); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Bad Request"})
	}
	res, err := ac.client.UpdateAuthor(context.Background(), &req)
	if err != nil {
		log.Printf("Error calling UpdateAuthor: %v", err)
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Internal Server Error"})
	}
	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{"author": res.Author})
}

func (ac *AuthorClient) DeleteAuthor(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	req := &pb.DeleteAuthorRequest{Id: id}
	_, err := ac.client.DeleteAuthor(context.Background(), req)
	if err != nil {
		log.Printf("Error calling DeleteAuthor: %v", err)
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Internal Server Error"})
	}
	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{"message": "Author deleted successfully"})
}

func (ac *AuthorClient) DeleteAuthorByName(ctx *fiber.Ctx) error {
	name := ctx.Params("name")
	req := &pb.DeleteAuthorByNameRequest{Name: name}
	_, err := ac.client.DeleteAuthorByName(context.Background(), req)
	if err != nil {
		log.Printf("Error calling DeleteAuthorByName: %v", err)
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Internal Server Error"})
	}
	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{"message": "Author deleted successfully"})
}

func (ac *AuthorClient) ListAuthors(ctx *fiber.Ctx) error {
	req := &pb.ListAuthorsRequest{}
	res, err := ac.client.ListAuthors(context.Background(), req)
	if err != nil {
		log.Printf("Error calling ListAuthors: %v", err)
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Internal Server Error"})
	}
	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{"authors": res.Authors})
}

// TODO: implement create author and delete author and get author by name
func RegisterAuthorRoutes(api fiber.Router, conn *grpc.ClientConn, secret string, userService service.UserService) {
	authorClient := NewAuthorClient(conn)
	api.Get("/author/:id", middleware.AuthAdminMiddleware(secret, userService), authorClient.GetAuthor)
	api.Get("/author/name/:name", middleware.AuthAdminMiddleware(secret, userService), authorClient.GetAuthorByName)
	api.Post("/author", middleware.AuthAdminMiddleware(secret, userService), authorClient.CreateAuthor)
	api.Put("/author", middleware.AuthAdminMiddleware(secret, userService), authorClient.UpdateAuthor)
	api.Delete("/author/:id", middleware.AuthAdminMiddleware(secret, userService), authorClient.DeleteAuthor)
	api.Delete("/author/name/:name", middleware.AuthAdminMiddleware(secret, userService), authorClient.DeleteAuthorByName)
	api.Get("/authors", middleware.AuthAdminMiddleware(secret, userService), authorClient.ListAuthors)
}
