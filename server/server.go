package server

import (
	"context"
	"gogrpc/api"
)

type GRPCServer struct{}

func (s *GRPCServer) FindBooksByAuthor(ctx context.Context, req *api.AuthorRequest) (*api.BooksResponseList, error) {
	var result = []*api.BookResponse{}
	result = append(result, &api.BookResponse{Name: req.GetAuthor()})
	return &api.BooksResponseList{Books: result}, nil
}

func (s *GRPCServer) FindAuthorsByBook(ctx context.Context, req *api.BookRequest) (*api.AuthorsResponseList, error) {
	var result = []*api.AuthorResponse{}
	result = append(result, &api.AuthorResponse{Name: req.GetBook()})
	return &api.AuthorsResponseList{Authors: result}, nil
}
