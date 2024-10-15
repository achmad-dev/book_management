package dto

// CreateBookRequest represents the request to create a book.
type CreateBookRequest struct {
	AuthorId   string `json:"author_id"`
	CategoryId string `json:"category_id"`
	Title      string `json:"title"`
	Stock      int    `json:"stock"`
}

type BorrowBookRequest struct {
	BookId   string `json:"book_id"`
	Title    string `json:"title"`
	Quantity int    `json:"quantity"`
}

type ReturnBookRequest struct {
	BookId   string `json:"book_id"`
	Quantity int    `json:"quantity"`
}

// UpdateBookRequest represents the request to update a book.
type UpdateBookRequest struct {
	Id        string `json:"id"`
	Title     string `json:"title"`
	Author    string `json:"author"`
	Category  string `json:"category"`
	Stock     int    `json:"stock"`
	Borrowed  int    `json:"borrowed"`
	IsPopular bool   `json:"is_popular"`
}
