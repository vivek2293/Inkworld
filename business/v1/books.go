package businessv1

import (
	"context"

	business_utils "github.com/vivek2293/Inkworld/business/utils"
	models "github.com/vivek2293/Inkworld/models"
	"github.com/vivek2293/Inkworld/repository"
)

// GetAllBookDetails retrieves all book details from the repository and returns the response in a specific format.
func GetAllBookDetails(ctx context.Context) (*[]models.BookDetailsResponse, error) {
	bookList, err := repository.GetAllBookDetails(ctx)
	if err != nil {
		return nil, err
	}

	return business_utils.SetBookListResponse(bookList), nil
}

// GetBookDetailsByID retrieves the details of a book by its ID from the repository and returns the response in a specific format.
func GetBookDetailsByID(ctx context.Context, id string) (*models.BookDetailsResponse, error) {
	bookDetails, err := repository.GetBookDetailsByID(ctx, id)
	if err != nil {
		return nil, err
	}

	return business_utils.SetBookResponse(bookDetails), nil
}
