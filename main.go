package main

import (
	"fmt"

	"github.com/vivek2293/Inkworld/constants"
	router "github.com/vivek2293/Inkworld/routes"
)

func main() {
	fmt.Println("Bookstore application running")
	// ctx := context.Background()
	/*
		In Go, the context package provides a way to manage concurrent operations by passing deadlines, cancellation signals, and request-scoped values across API boundaries and between goroutines

		Background returns a non-nil, empty [Context]. It is never canceled, has no values, and has no deadline. It is typically used by the main function, initialization, and tests, and as the top-level Context for incoming requests.
	*/

	initRouter()
}

func initRouter() {
	router, err := router.GetRouter()
	if err != nil {
		panic(err)
	}

	err = router.Run(":" + constants.RouterPort)
	if err != nil {
		panic("Unable to start router")
	}

	fmt.Println("Server running ...")
}
