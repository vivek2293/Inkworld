package repository

import (
	"context"

	"github.com/vivek2293/Inkworld/store"
)

var bookdetails = []store.Book{
	{
		ID:     "1",
		Title:  "The Richest Man in Babylon",
		Genre:  "Personal Finance",
		Author: "George S. Clason",
		Price:  "Rs. 120",
	},
	{
		ID:     "2",
		Title:  "The Richest Man in Babylon",
		Genre:  "Personal Finance",
		Author: "George S. Clason",
		Price:  "Rs. 120",
	},
}

// GetAllBookDetails provides repo layer for get all book details
func GetAllBookDetails(ctx context.Context) (*[]store.Book, error) {

	return &bookdetails, nil
}

// GetBookDetailsByID provides repo layer for get book details by ID
func GetBookDetailsByID(ctx context.Context, id string) (*store.Book, error) {
	for _, book := range bookdetails {
		if book.ID == id {
			return &book, nil
		}
	}

	return nil, nil
}
