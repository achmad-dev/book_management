package grpcserver

import (
	"context"

	"github.com/achmad-dev/internal/author/internal/domain"
	pb "github.com/achmad-dev/internal/pkg/common/genproto"
)

func (s *server) GetAuthor(ctx context.Context, req *pb.GetAuthorRequest) (*pb.GetAuthorResponse, error) {
	author, err := s.authorService.GetAuthor(req.Id)
	if err != nil {
		return nil, err
	}
	return &pb.GetAuthorResponse{
		Author: &pb.Author{
			Id:   author.ID,
			Name: author.Name,
		},
	}, nil
}

func (s *server) GetAuthorByName(ctx context.Context, req *pb.GetAuthorByNameRequest) (*pb.GetAuthorByNameResponse, error) {
	author, err := s.authorService.GetAuthorByName(req.Name)
	if err != nil {
		return nil, err
	}
	return &pb.GetAuthorByNameResponse{
		Author: &pb.Author{
			Id:   author.ID,
			Name: author.Name,
		},
	}, nil
}

func (s *server) CreateAuthor(ctx context.Context, req *pb.CreateAuthorRequest) (*pb.CreateAuthorResponse, error) {
	author := &domain.Author{
		Name: req.Name,
	}
	err := s.authorService.CreateAuthor(author)
	if err != nil {
		return nil, err
	}
	return &pb.CreateAuthorResponse{
		Author: &pb.Author{
			Name: author.Name,
		},
	}, nil
}

func (s *server) UpdateAuthor(ctx context.Context, req *pb.UpdateAuthorRequest) (*pb.UpdateAuthorResponse, error) {
	author := &domain.Author{
		ID:   req.Id,
		Name: req.Name,
	}
	err := s.authorService.UpdateAuthor(req.Id, author)
	if err != nil {
		return nil, err
	}
	return &pb.UpdateAuthorResponse{
		Author: &pb.Author{
			Id:   author.ID,
			Name: author.Name,
		},
	}, nil
}

func (s *server) DeleteAuthor(ctx context.Context, req *pb.DeleteAuthorRequest) (*pb.DeleteAuthorResponse, error) {
	err := s.authorService.DeleteAuthor(req.Id)
	if err != nil {
		return nil, err
	}
	return &pb.DeleteAuthorResponse{}, nil
}

func (s *server) DeleteAuthorByName(ctx context.Context, req *pb.DeleteAuthorByNameRequest) (*pb.DeleteAuthorByNameResponse, error) {
	err := s.authorService.DeleteAuthorByName(req.Name)
	if err != nil {
		return nil, err
	}
	return &pb.DeleteAuthorByNameResponse{}, nil
}

func (s *server) ListAuthors(ctx context.Context, req *pb.ListAuthorsRequest) (*pb.ListAuthorsResponse, error) {
	authors, err := s.authorService.ListAuthors()
	if err != nil {
		return nil, err
	}
	var pbAuthors []*pb.Author
	for _, author := range authors {
		pbAuthors = append(pbAuthors, &pb.Author{
			Id:   author.ID,
			Name: author.Name,
		})
	}
	return &pb.ListAuthorsResponse{
		Authors: pbAuthors,
	}, nil
}
