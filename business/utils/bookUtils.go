package businessutils

import (
	modelsv1 "github.com/vivek2293/Inkworld/models/v1"
	"github.com/vivek2293/Inkworld/store"
)

// SetBookListResponse converts a list of book objects to the response format.
func SetBookListResponse(bookList *[]store.Book) *[]modelsv1.BookDetailsResponse {
	if bookList == nil {
		return nil
	}

	convertedBookList := []modelsv1.BookDetailsResponse{}

	for _, book := range *bookList {
		bookDetails := modelsv1.BookDetailsResponse{}

		bookDetails.ID = book.ID
		bookDetails.Author = book.Author
		bookDetails.Genre = book.Genre
		bookDetails.Price = book.Price
		bookDetails.Title = book.Title

		convertedBookList = append(convertedBookList, bookDetails)
	}

	return &convertedBookList
}

// SetBookResponse converts a single book object to the response format.
func SetBookResponse(book *store.Book) *modelsv1.BookDetailsResponse {
	if book == nil {
		return nil
	}
	convertedBook := modelsv1.BookDetailsResponse{}

	convertedBook.ID = book.ID
	convertedBook.Author = book.Author
	convertedBook.Genre = book.Genre
	convertedBook.Price = book.Price
	convertedBook.Title = book.Title

	return &convertedBook
}
