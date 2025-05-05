package router

import (
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	apiv1 "github.com/vivek2293/Inkworld/api/v1"
	"github.com/vivek2293/Inkworld/constants"
	"github.com/vivek2293/Inkworld/utils/env"
	"github.com/vivek2293/Inkworld/utils/logger"
	"github.com/vivek2293/Inkworld/utils/middlewares"
	"go.opentelemetry.io/contrib/instrumentation/github.com/gin-gonic/gin/otelgin"
)

// GetRouter initializes the router and sets up the routes for the application.
func GetRouter() (*gin.Engine, error) {
	router := gin.New()

	currentMode := env.GetEnv(constants.GetEnvModeKey)
	switch currentMode {
	case constants.Production:
		gin.SetMode(gin.ReleaseMode)
	default:
		gin.SetMode(gin.DebugMode)
	}

	// logger middleware
	router.Use(middlewares.GinZapLogger(logger.GetLogger()))
	// tracing middleware
	router.Use(otelgin.Middleware(""))
	router.Use(gin.Recovery()) // Recovers from any panics with status code 500

	router.GET("/metrics", gin.WrapH(promhttp.Handler()))
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

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
