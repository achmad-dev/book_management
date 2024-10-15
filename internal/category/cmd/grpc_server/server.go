package grpcserver

import (
	"github.com/achmad-dev/internal/category/internal/service"
	pb "github.com/achmad-dev/internal/pkg/common/genproto"
)

type server struct {
	pb.UnimplementedCategoryServiceServer
	categoryService service.CategoryService
}

func NewServer(categoryService service.CategoryService) *server {
	return &server{categoryService: categoryService}
}
