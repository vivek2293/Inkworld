package apiv1

import (
	"net/http"

	"github.com/gin-gonic/gin"
	businessv1 "github.com/vivek2293/Inkworld/business/v1"
)

//go:generate mockgen -source=books.go -destination=mocks/books.go -package=mockhandler

// HandleGetAllBookDetails is api handler for get all book details
func HandleGetAllBookDetails(ctx *gin.Context) {
	response, err := businessv1.GetAllBookDetails(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, "internal server error")
		return
	}

	ctx.JSON(http.StatusOK, response)
}

// HandleGetBookDetailsByID is api handler for get book details by ID
func HandleGetBookDetailsByID(ctx *gin.Context) {
	id := ctx.Param("id")

	response, err := businessv1.GetBookDetailsByID(ctx, id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, "internal server error")
		return
	}

	ctx.JSON(http.StatusOK, response)
}
