package main

import (
	"context"
	"matemo.net/go-grpc-gateway/proto"
	"log"
)

type BookService struct {}

func NewBookService() *BookService {
	return &BookService{}
}

func (s *BookService) Add(ctx context.Context, request *proto.BookRequest) (*proto.BookResponse, error) {
	log.Printf("Add Book with Id %d", request.GetBook().BookId)

	return &proto.BookResponse{}, nil
}

func (s *BookService) Get(ctx context.Context, request *proto.BookRequest) (*proto.BookResponse, error) {
	log.Printf("Get Book with Id %d", request.GetBookId())

	return &proto.BookResponse{}, nil
}

func (s *BookService) Update(ctx context.Context, request *proto.BookRequest) (*proto.BookResponse, error) {
	log.Printf("Update Book with Id %d", request.GetBook().BookId)

	return &proto.BookResponse{}, nil
}

func (s *BookService) Delete(ctx context.Context, request *proto.BookRequest) (*proto.BookResponse, error) {
	log.Printf("Delete Book with Id %d", request.GetBook().BookId)

	return &proto.BookResponse{}, nil
}

func (s *BookService) Filter(ctx context.Context, request *proto.FilterRequest) (*proto.FilterResponse, error) {
	log.Printf("Filter Books with page %d size %d", request.Page, request.Size)

	return &proto.FilterResponse{}, nil
}