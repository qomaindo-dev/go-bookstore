package graph

import (
	"context"

	"github.com/qomaindo-dev/go-bookstore/graph/model"
)

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.34

// Mutation returns generated.MutationResolver implementation.
// func (r *ResolverBook) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
// func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

func (r *mutationResolverBook) CreateBook(ctx context.Context, input model.BookInput) (*model.Book, error) {
	book, err := r.BookRepository.CreateBook(&input)
	bookCreated := &model.Book{
		Title:     book.Title,
		Synopsis:  book.Synopsis,
		Authore:   book.Author,
		Publisher: book.Publisher,
		Stock:     book.Stock,
		Price:     book.Price,
		Status:    book.Status,
		ID:        book.ID,
	}
	if err != nil {
		return nil, err
	}

	return bookCreated, nil
}

func (r *mutationResolverBook) UpdateBook(ctx context.Context, id int, input model.BookInput) (string, error) {
	err := r.BookRepository.UpdateBook(&input, id)
	if err != nil {
		return "null", err
	}

	message := "Success update data!"

	return message, nil
}

func (r *mutationResolverBook) DeleteBook(ctx context.Context, id int) (string, error) {
	err := r.BookRepository.DeleteBook(id)
	if err != nil {
		return "", err
	}

	message := "Success delete data"

	return message, nil
}

func (r *queryResolverBook) GetOneBook(ctx context.Context, id int) (*model.Book, error) {
	book, err := r.BookRepository.GetOneBook(id)
	selectedBook := &model.Book{
		ID:        book.ID,
		Title:     book.Title,
		Synopsis:  book.Synopsis,
		Authore:   book.Author,
		Publisher: book.Publisher,
		Stock:     book.Stock,
		Price:     book.Price,
		Status:    book.Status,
	}
	if err != nil {
		return nil, err
	}

	return selectedBook, nil
}

func (r *queryResolverBook) GetAllBooks(ctx context.Context) ([]*model.Book, error) {
	books, err := r.BookRepository.GetAllBooks()
	if err != nil {
		return nil, err
	}

	return books, err
}

type mutationResolverBook struct {
	*Resolver
}
type queryResolverBook struct {
	*Resolver
}