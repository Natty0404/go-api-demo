package controller

import (
	"go-api-demo/entity/repository"
)

type GoApiDemo struct {
	bookRepo repository.Book
}

func NewGoApiDemo(
	bookRepo repository.Book,
) *GoApiDemo {
	return &GoApiDemo{
		bookRepo: bookRepo,
	}
}
