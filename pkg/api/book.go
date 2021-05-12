package api

// BookService contains methods of the books service
type BookService interface{}

// BookRepository is what lets our service do db operations without knowing anything about the implementation
type BookRepository interface{}

type bookService struct {
	storage BookRepository
}

func NewBookService(bookRepo BookRepository) BookService {
	return &bookService{
		storage: bookRepo,
	}
}
