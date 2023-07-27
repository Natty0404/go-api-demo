package controller

import (
	"context"
)

func (g *GoApiDemo) GetBook(ctx context.Context, BookID string) (interface{}, error) {
	book, err := g.bookRepo.Get(BookID)
	if err != nil {
		return nil, err
	}
	return book, nil
}
