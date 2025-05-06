package main

import (
	"context"
	"fmt"
	"time"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracehttp"
	"go.opentelemetry.io/otel/sdk/resource"
	"go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.25.0"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"

	"github.com/vivek2293/Inkworld/constants"
	"github.com/vivek2293/Inkworld/database"
	router "github.com/vivek2293/Inkworld/routes"
	"github.com/vivek2293/Inkworld/utils/env"
	"github.com/vivek2293/Inkworld/utils/logger"
	"github.com/vivek2293/Inkworld/utils/monitoring"
	timeutils "github.com/vivek2293/Inkworld/utils/time"
)

func main() {
	ctx := context.Background()
	initEnv()
	initLogger()
	defer logger.Sync()
	initTime()
	tp := initTracer(ctx)
	defer tp.Shutdown(ctx)
	initMonitoring()
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

func initTracer(ctx context.Context) *trace.TracerProvider {
	headers := map[string]string{
		"content-type": "application/json",
	}
	exporter, err := otlptrace.New(
		ctx,
		otlptracehttp.NewClient(
			otlptracehttp.WithEndpoint(env.GetEnv(constants.OtlpEndpoint)),
			otlptracehttp.WithHeaders(headers),
			otlptracehttp.WithInsecure(), // Not recommended for production
		),
	)

	if err != nil {
		logger.Panic("Error initializing OTLP trace exporter", zap.Error(err))
	}

	tp := trace.NewTracerProvider(
		trace.WithBatcher(
			exporter,
			trace.WithMaxExportBatchSize(trace.DefaultMaxExportBatchSize),
			trace.WithBatchTimeout(trace.DefaultScheduleDelay*time.Millisecond),
		),
		trace.WithResource(
			resource.NewWithAttributes(
				semconv.SchemaURL,
				semconv.ServiceNameKey.String(constants.TracerName),
				attribute.String("environment", env.GetEnv(constants.GetEnvModeKey)),
				attribute.String("monitor_name", constants.MonitorName),
			),
		),
	)

	otel.SetTracerProvider(tp)
	logger.Info("Tracer provider initialized successfully")
	return tp
}

func initMonitoring() {
	// Register Prometheus metrics
	err := monitoring.RegisterPrometheus()

	if err != nil {
		logger.Panic("Error registering Prometheus metrics", zap.Error(err))
	}
	logger.Info("Monitoring initialized successfully")
}

func initTime() {
	err := timeutils.InitLocation(constants.DefaultLocation)
	if err != nil {
		logger.Panic("Error initializing time location", zap.Error(err))
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
		logger.Panic("Error initializing database", zap.Error(err))
	}

	logger.Info("Database initialized successfully")
}

func initRouter() {
	router, err := router.GetRouter()
	if err != nil {
		logger.Panic("Error initializing router", zap.Error(err))
	}

	if err = router.Run(":" + constants.RouterPort); err != nil {
		logger.Panic("Error running router", zap.Error(err))
	}
}
