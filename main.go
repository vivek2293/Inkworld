package main

import (
	"context"
	"fmt"
	"time"

	"github.com/vivek2293/Inkworld/constants"
	"github.com/vivek2293/Inkworld/database"
	router "github.com/vivek2293/Inkworld/routes"
	"github.com/vivek2293/Inkworld/utils/env"
	timeutils "github.com/vivek2293/Inkworld/utils/time"
)

func main() {
	fmt.Println("Bookstore application running")
	ctx := context.Background()

	initEnv(ctx)
	initTime(ctx)
	initDB(ctx)
	defer database.CloseDB()
	initRouter(ctx)
}

func initEnv(ctx context.Context) {
	err := env.InitEnv(constants.GenENVPath)
	if err != nil {
		panic(err)
	}

	currentMode := env.GetEnv(constants.GetCurrentModeKey)

	if currentMode == constants.Production {
		err = env.InitEnv(constants.ProdENVPath)
	} else {
		err = env.InitEnv(constants.DevENVPath)
	}

	if err != nil {
		panic(err)
	}
}

func initTime(ctx context.Context) {
	err := timeutils.InitLocation(constants.DefaultLocation)
	if err != nil {
		panic(err)
	}
}

func initDB(ctx context.Context) {
	configSettings := database.DbConfig{
		DriverName:            env.GetEnv("DB_DRIVER_NAME"),
		URL:                   env.GetEnv("DB_URL"),
		MaxOpenConnections:    env.GetEnvInt("DB_MAX_OPEN_CONNECTIONS"),
		MaxIdleConnections:    env.GetEnvInt("DB_MAX_IDLE_CONNECTIONS"),
		ConnectionMaxLifeTime: time.Duration(env.GetEnvInt("DB_CONNECTION_MAX_LIFE_TIME")) * time.Second,
		ConnectionMaxIdleTime: time.Duration(env.GetEnvInt("DB_CONNECTION_MAX_IDLE_TIME")) * time.Second,
	}

	err := database.InitDatabase(configSettings)
	if err != nil {
		panic(err)
	}

	fmt.Println("Database initialized successfully")
}

func initRouter(ctx context.Context) {
	router, err := router.GetRouter()
	if err != nil {
		panic(err)
	}

	if err = router.Run(":" + constants.RouterPort); err != nil {
		panic(err)
	}
}
