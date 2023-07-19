package book

type Book struct {
	BookID      string `json:"id"`
	BookTitle   string `json:"title"`
	BookContent string `json:"content"`
	Author      string `json:"author"`
}
