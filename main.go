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

	currentMode := env.GetEnv(constants.GetEnvModeKey)

	if currentMode == constants.Production {
		err = env.InitEnv(constants.ProdENVPath)
	} else if currentMode == constants.Development {
		err = env.InitEnv(constants.DevENVPath)
	} else {
		err = fmt.Errorf("unknown environment mode: %s", currentMode)
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
		DriverName:            env.GetEnv(constants.DriverName),
		URL:                   env.GetEnv(constants.URL),
		MaxOpenConnections:    env.GetEnvInt(constants.MaxOpenConnections),
		MaxIdleConnections:    env.GetEnvInt(constants.MaxIdleConnections),
		ConnectionMaxLifeTime: time.Duration(env.GetEnvInt(constants.ConnectionMaxLifeTime)) * time.Second,
		ConnectionMaxIdleTime: time.Duration(env.GetEnvInt(constants.ConnectionMaxIdleTime)) * time.Second,
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
