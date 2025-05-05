package constants

// Environment variables configuration keys
const (
	DriverName            = "DB_DRIVER_NAME"
	URL                   = "DB_URL"
	MaxOpenConnections    = "DB_MAX_OPEN_CONNECTIONS"
	MaxIdleConnections    = "DB_MAX_IDLE_CONNECTIONS"
	ConnectionMaxLifeTime = "DB_CONNECTION_MAX_LIFE_TIME"
	ConnectionMaxIdleTime = "DB_CONNECTION_MAX_IDLE_TIME"
)

// Environment variable file paths
const (
	GenENVPath  = ".env"
	ProdENVPath = ".env.prod"
	DevENVPath  = ".env.dev"
)

// Current environment mode
const GetEnvModeKey = "ENV_MODE"
const OtlpEndpoint = "OTLP_ENDPOINT"
