package router

import (
	"github.com/gin-gonic/gin"
	apiv1 "github.com/vivek2293/Inkworld/api/v1"
	"github.com/vivek2293/Inkworld/constants"
)

// GetRouter initializes the router and sets up the routes for the application.
func GetRouter() (*gin.Engine, error) {
	router := gin.New()

	v1AuthRouter := router.Group(constants.Version1)
	{
		addBookRoutes(v1AuthRouter)
	}

	return router, nil
}

/*
	*gin.RouterGroup

	RouterGroup is used internally to configure router, a RouterGroup is associated with a prefix and an array of handlers (middleware).

	func (group *gin.RouterGroup) Any(relativePath string, handlers ...gin.HandlerFunc) gin.IRoutes
	func (group *gin.RouterGroup) BasePath() string
	func (group *gin.RouterGroup) DELETE(relativePath string, handlers ...gin.HandlerFunc) gin.IRoutes
	func (group *gin.RouterGroup) GET(relativePath string, handlers ...gin.HandlerFunc) gin.IRoutes
	func (group *gin.RouterGroup) Group(relativePath string, handlers ...gin.HandlerFunc) *gin.RouterGroup
*/

func addBookRoutes(group *gin.RouterGroup) {
	book := group.Group(constants.BookRoute)
	{
		book.GET(constants.GetAllBookDetails, apiv1.HandleGetAllBookDetails)
		book.GET(constants.GetBookDetailsByID, apiv1.HandleGetBookDetailsByID)
	}
}
