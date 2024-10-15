package grpcserver

import (
	"github.com/achmad-dev/internal/book/internal/service"
	pb "github.com/achmad-dev/internal/pkg/common/genproto"
)

type server struct {
	pb.UnimplementedBookServiceServer
	bookService service.BookService
}

func NewServer(bookService service.BookService) *server {
	return &server{bookService: bookService}
}
