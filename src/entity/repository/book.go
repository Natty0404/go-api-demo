package repository

import "go-api-demo/entity/model/book"

type Book interface {
	Get(BookID string) (*book.Book, error)
}
