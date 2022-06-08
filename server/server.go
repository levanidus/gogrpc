package server

import (
	"context"
	"gogrpc/api"
	"gogrpc/repository"

	"github.com/jmoiron/sqlx"
)

type GRPCServer struct {
	repo *repository.Repository
}

func NewGRPCServer(db *sqlx.DB) *GRPCServer {
	return &GRPCServer{repo: repository.NewRepository(db)}
}

func (s *GRPCServer) FindBooksByAuthor(ctx context.Context, req *api.AuthorRequest) (*api.BooksResponseList, error) {
	var result = []*api.BookResponse{}
	books, err := s.repo.FindBooksByAuthor(req.GetAuthor())
	if err != nil {
		return &api.BooksResponseList{Books: result}, nil
	}

	for _, book := range books {
		result = append(result, &api.BookResponse{Name: book})
	}

	return &api.BooksResponseList{Books: result}, nil
}

func (s *GRPCServer) FindAuthorsByBook(ctx context.Context, req *api.BookRequest) (*api.AuthorsResponseList, error) {
	var result = []*api.AuthorResponse{}
	authors, err := s.repo.FindAuthorsByBook(req.GetBook())
	if err != nil {
		return &api.AuthorsResponseList{Authors: result}, nil
	}

	for _, author := range authors {
		result = append(result, &api.AuthorResponse{Name: author})
	}

	return &api.AuthorsResponseList{Authors: result}, nil
}
