package main

import (
	"fmt"
	"time"

	"github.com/vivek2293/Inkworld/constants"
	"github.com/vivek2293/Inkworld/database"
	router "github.com/vivek2293/Inkworld/routes"
	"github.com/vivek2293/Inkworld/utils/env"
	"github.com/vivek2293/Inkworld/utils/logger"
	timeutils "github.com/vivek2293/Inkworld/utils/time"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func main() {
	// ctx := context.Background()
	initEnv()
	initLogger()
	defer logger.Sync()
	initTime()
	initDB()
	defer database.CloseDB()
	initRouter()
}

func initEnv() {
	// Common environment variables
	err := env.ReadEnv(constants.GenENVPath)
	if err != nil {
		panic(err)
	}

	currentMode := env.GetEnv(constants.GetEnvModeKey)
	switch currentMode {
	case constants.Production:
		err = env.ReadEnv(constants.ProdENVPath)
	case constants.Development:
		err = env.ReadEnv(constants.DevENVPath)
	default:
		err = fmt.Errorf("unknown environment mode: %s", currentMode)
	}

	if err != nil {
		panic(err)
	}
}

func initLogger() {
	currentMode := env.GetEnv(constants.GetEnvModeKey)
	var logger *zap.Logger

	switch currentMode {
	case constants.Production:
		logger = zap.Must(zap.NewProduction())
	case constants.Development:
		config := zap.NewDevelopmentConfig()
		config.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
		logger = zap.Must(config.Build())
	default:
		err := fmt.Errorf("unknown environment mode: %s", currentMode)
		panic(err)
	}

	// zap.L() returns the global Logger, which can be reconfigured with ReplaceGlobals. It's safe for concurrent use.
	zap.ReplaceGlobals(logger)
	logger.Info("Logger initialized successfully")
}

func initTime() {
	err := timeutils.InitLocation(constants.DefaultLocation)
	if err != nil {
		logger.Error("Error initializing time location", zap.Error(err))
		panic(err)
	}

	logger.Info("Time location initialized successfully")
}

func initDB() {
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
		logger.Error("Error initializing database", zap.Error(err))
		panic(err)
	}

	logger.Info("Database initialized successfully")
}

func initRouter() {
	router, err := router.GetRouter()
	if err != nil {
		logger.Error("Error initializing router", zap.Error(err))
		panic(err)
	}

	if err = router.Run(":" + constants.RouterPort); err != nil {
		logger.Error("Error running router", zap.Error(err))
		panic(err)
	}
}
