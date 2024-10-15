package grpcserver

import (
	"github.com/achmad-dev/internal/author/internal/service"
	pb "github.com/achmad-dev/internal/pkg/common/genproto"
)

type server struct {
	pb.UnimplementedAuthorServiceServer
	authorService service.AuthorService
}

func NewServer(authorService service.AuthorService) *server {
	return &server{
		authorService: authorService,
	}
}
