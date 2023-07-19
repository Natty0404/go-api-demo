package gateway

import (
	"go-api-demo/entity/model/book"
)

type Book struct{}

func NewBook() *Book {
	return &Book{}
}

func (b *Book) Get(BookID string) (*book.Book, error) {
	// successならbook.Bookを返す
	// errorならnilとエラーを返す
	return &book.Book{
		BookID:      "1",
		BookTitle:   "test_title",
		BookContent: "test_content",
		Author:      "test_author",
	}, nil

}
