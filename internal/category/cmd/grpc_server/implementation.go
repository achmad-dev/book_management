package grpcserver

import (
	"context"

	pb "github.com/achmad-dev/internal/pkg/common/genproto"
)

func (s *server) CreateCategory(ctx context.Context, req *pb.CreateCategoryRequest) (*pb.CreateCategoryResponse, error) {
	err := s.categoryService.CreateCategory(req.Name)
	if err != nil {
		return nil, err
	}
	return &pb.CreateCategoryResponse{
		Category: &pb.Category{
			Name: req.Name,
		},
	}, nil
}

func (s *server) GetCategory(ctx context.Context, req *pb.GetCategoryRequest) (*pb.GetCategoryResponse, error) {
	category, err := s.categoryService.GetCategory(req.Id)
	if err != nil {
		return nil, err
	}
	return &pb.GetCategoryResponse{
		Category: &pb.Category{
			Id:   category.ID,
			Name: category.Name,
		},
	}, nil
}

func (s *server) GetCategoryByName(ctx context.Context, req *pb.GetCategoryByNameRequest) (*pb.GetCategoryByNameResponse, error) {
	category, err := s.categoryService.GetCategoryByName(req.Name)
	if err != nil {
		return nil, err
	}
	return &pb.GetCategoryByNameResponse{
		Category: &pb.Category{
			Id:   category.ID,
			Name: category.Name,
		},
	}, nil
}

func (s *server) UpdateCategory(ctx context.Context, req *pb.UpdateCategoryRequest) (*pb.UpdateCategoryResponse, error) {
	err := s.categoryService.UpdateCategory(req.Id, req.Name)
	if err != nil {
		return nil, err
	}
	return &pb.UpdateCategoryResponse{}, nil
}

func (s *server) DeleteCategory(ctx context.Context, req *pb.DeleteCategoryRequest) (*pb.DeleteCategoryResponse, error) {
	err := s.categoryService.DeleteCategory(req.Id)
	if err != nil {
		return nil, err
	}
	return &pb.DeleteCategoryResponse{}, nil
}

func (s *server) ListCategories(ctx context.Context, req *pb.ListCategoriesRequest) (*pb.ListCategoriesResponse, error) {
	categories, err := s.categoryService.ListCategories()
	if err != nil {
		return nil, err
	}
	var pbCategories []*pb.Category
	for _, category := range categories {
		pbCategories = append(pbCategories, &pb.Category{
			Id:   category.ID,
			Name: category.Name,
		})
	}
	return &pb.ListCategoriesResponse{
		Categories: pbCategories,
	}, nil
}
