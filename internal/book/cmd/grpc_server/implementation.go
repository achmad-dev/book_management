package grpcserver

import (
	"context"

	"github.com/achmad-dev/internal/book/internal/domain"
	pb "github.com/achmad-dev/internal/pkg/common/genproto"
)

func (s *server) GetBook(ctx context.Context, req *pb.GetBookRequest) (*pb.GetBookResponse, error) {
	book, err := s.bookService.GetBook(req.Id)
	if err != nil {
		return nil, err
	}
	return &pb.GetBookResponse{Book: &pb.Book{
		Id:        book.ID,
		Title:     book.Title,
		Author:    book.Author,
		Category:  book.Category,
		Stock:     int32(book.Stock),
		Borrowed:  int32(book.Borrowed),
		IsPopular: book.IsPopular,
	}}, nil
}
func (s *server) GetBookByTitle(ctx context.Context, req *pb.GetBookByTitleRequest) (*pb.GetBookByTitleResponse, error) {
	book, err := s.bookService.GetBookByTitle(req.Title)
	if err != nil {
		return nil, err
	}
	return &pb.GetBookByTitleResponse{Book: &pb.Book{
		Id:        book.ID,
		Title:     book.Title,
		Author:    book.Author,
		Category:  book.Category,
		Stock:     int32(book.Stock),
		Borrowed:  int32(book.Borrowed),
		IsPopular: book.IsPopular,
	}}, nil
}

func (s *server) GetBooksByAuthorName(ctx context.Context, req *pb.GetBooksByAuthorNameRequest) (*pb.GetBooksByAuthorNameResponse, error) {
	books, err := s.bookService.GetBooksByAuthorName(req.Author)
	if err != nil {
		return nil, err
	}
	var pbBooks []*pb.Book
	for _, book := range books {
		pbBooks = append(pbBooks, &pb.Book{
			Id:        book.ID,
			Title:     book.Title,
			Author:    book.Author,
			Category:  book.Category,
			Stock:     int32(book.Stock),
			Borrowed:  int32(book.Borrowed),
			IsPopular: book.IsPopular,
		})
	}
	return &pb.GetBooksByAuthorNameResponse{Books: pbBooks}, nil
}

func (s *server) GetPopularBooksByCategory(ctx context.Context, req *pb.GetPopularBooksByCategoryRequest) (*pb.GetPopularBooksByCategoryResponse, error) {
	books, err := s.bookService.GetPopularBooksByCategory(req.Category)
	if err != nil {
		return nil, err
	}
	var pbBooks []*pb.Book
	for _, book := range books {
		pbBooks = append(pbBooks, &pb.Book{
			Id:        book.ID,
			Title:     book.Title,
			Author:    book.Author,
			Category:  book.Category,
			Stock:     int32(book.Stock),
			Borrowed:  int32(book.Borrowed),
			IsPopular: book.IsPopular,
		})
	}
	return &pb.GetPopularBooksByCategoryResponse{Books: pbBooks}, nil
}

func (s *server) BorrowBook(ctx context.Context, req *pb.BorrowBookRequest) (*pb.BorrowBookResponse, error) {
	err := s.bookService.BorrowBook(req.Id, int(req.Quantity))
	if err != nil {
		return nil, err
	}
	return &pb.BorrowBookResponse{Book: &pb.Book{}}, nil
}

func (s *server) ReturnBook(ctx context.Context, req *pb.ReturnBookRequest) (*pb.ReturnBookResponse, error) {
	err := s.bookService.ReturnBook(req.Id, int(req.Quantity))
	if err != nil {
		return nil, err
	}
	return &pb.ReturnBookResponse{Book: &pb.Book{}}, nil
}

func (s *server) CreateBook(ctx context.Context, req *pb.CreateBookRequest) (*pb.CreateBookResponse, error) {
	book := &domain.Book{
		Author_id:   req.AuthorId,
		Category_id: req.CategoryId,
		Title:       req.Title,
		Author:      req.Author,
		Category:    req.Category,
		Stock:       int(req.Stock),
		Borrowed:    int(req.Borrowed),
		IsPopular:   req.IsPopular,
	}
	err := s.bookService.CreateBook(book)
	if err != nil {
		return nil, err
	}
	return &pb.CreateBookResponse{Book: &pb.Book{}}, nil
}

func (s *server) UpdateBook(ctx context.Context, req *pb.UpdateBookRequest) (*pb.UpdateBookResponse, error) {
	book := &domain.Book{
		ID:        req.Id,
		Title:     req.Title,
		Author:    req.Author,
		Category:  req.Category,
		Stock:     int(req.Stock),
		Borrowed:  int(req.Borrowed),
		IsPopular: req.IsPopular,
	}
	err := s.bookService.UpdateBook(req.Id, book)
	if err != nil {
		return nil, err
	}
	return &pb.UpdateBookResponse{Book: &pb.Book{}}, nil
}

func (s *server) DeleteBook(ctx context.Context, req *pb.DeleteBookRequest) (*pb.DeleteBookResponse, error) {
	err := s.bookService.DeleteBook(req.Id)
	if err != nil {
		return nil, err
	}
	return &pb.DeleteBookResponse{Success: true}, nil
}

func (s *server) ListBooks(ctx context.Context, req *pb.ListBooksRequest) (*pb.ListBooksResponse, error) {
	books, err := s.bookService.ListBooks()
	if err != nil {
		return nil, err
	}
	var pbBooks []*pb.Book
	for _, book := range books {
		pbBooks = append(pbBooks, &pb.Book{
			Id:        book.ID,
			Title:     book.Title,
			Author:    book.Author,
			Category:  book.Category,
			Stock:     int32(book.Stock),
			Borrowed:  int32(book.Borrowed),
			IsPopular: book.IsPopular,
		})
	}
	return &pb.ListBooksResponse{Books: pbBooks}, nil
}
