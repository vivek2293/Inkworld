package database

import (
	"context"
	"database/sql"
	"time"

	_ "github.com/lib/pq"
	"github.com/vivek2293/Inkworld/utils/logger"
	"go.uber.org/zap"
)

// Config is set of configurable parameters of a database.
type DbConfig struct {
	DriverName            string
	URL                   string
	MaxOpenConnections    int
	MaxIdleConnections    int
	ConnectionMaxLifeTime time.Duration
	ConnectionMaxIdleTime time.Duration
}

// DbClient is a wrapper around sql.DB to provide a singleton instance of the database client.
// It is used to manage the database connection pool and provide a single point of access to the database.
type dbClient struct {
	*sql.DB
}

// DbTxClient is a wrapper around sql.Tx to provide a transaction client for the database.
// It is used to manage transactions and provide a single point of access to the database transaction.
type dbTxClient struct {
	*sql.Tx
}

var dbClientInstance *dbClient

func InitDatabase(config DbConfig) error {
	conn, err := sql.Open(config.DriverName, config.URL)
	if err != nil {
		return err
	}

	// Check if the connection is alive
	if err = conn.Ping(); err != nil {
		return err
	}

	conn.SetMaxOpenConns(config.MaxOpenConnections)
	conn.SetMaxIdleConns(config.MaxIdleConnections)
	conn.SetConnMaxLifetime(config.ConnectionMaxLifeTime)
	conn.SetConnMaxIdleTime(config.ConnectionMaxIdleTime)

	dbClientInstance = &dbClient{conn}

	return nil
}

// GetDbClient returns the singleton instance of the database client.
func GetDbClient() *dbClient {
	if dbClientInstance == nil {
		logger.Panic("Database client is not initialized")
	}
	return dbClientInstance
}

// GetTxClient returns a new transaction client for the database.
// It is used to manage transactions and provide a single point of access to the database transaction.
func GetTxClient(ctx context.Context, options *sql.TxOptions) (*dbTxClient, error) {
	tx, err := dbClientInstance.BeginTx(ctx, options)
	if err != nil {
		return nil, err
	}
	return &dbTxClient{tx}, nil
}

// CommitTx commits the transaction and releases the resources associated with it.
// It is used to commit the changes made in the transaction to the database.
func CommitTx(tx *dbTxClient) error {
	if tx == nil {
		return nil
	}

	if err := tx.Commit(); err != nil {
		return err
	}
	return nil
}

// RollbackTx rolls back the transaction and releases the resources associated with it.
// It is used to discard the changes made in the transaction to the database.
func RollbackTx(tx *dbTxClient) error {
	if tx == nil {
		return nil
	}

	if err := tx.Rollback(); err != nil {
		return err
	}
	return nil
}

// CloseDatabase closes the database connection and release the resources associated with it.
func CloseDB() {
	if dbClientInstance == nil {
		return
	}

	if err := dbClientInstance.Close(); err != nil {
		logger.Error("Error closing database connection", zap.Error(err))
		return
	}

	dbClientInstance = nil
}
