package businessutils

import (
	models "github.com/vivek2293/Inkworld/models"
	"github.com/vivek2293/Inkworld/store"
)

// SetBookListResponse converts a list of book objects to the response format.
func SetBookListResponse(bookList *[]store.Book) *[]models.BookDetailsResponse {
	if bookList == nil {
		return nil
	}

	convertedBookList := []models.BookDetailsResponse{}

	for _, book := range *bookList {
		bookDetails := models.BookDetailsResponse{}

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
func SetBookResponse(book *store.Book) *models.BookDetailsResponse {
	if book == nil {
		return nil
	}
	convertedBook := models.BookDetailsResponse{}

	convertedBook.ID = book.ID
	convertedBook.Author = book.Author
	convertedBook.Genre = book.Genre
	convertedBook.Price = book.Price
	convertedBook.Title = book.Title

	return &convertedBook
}
