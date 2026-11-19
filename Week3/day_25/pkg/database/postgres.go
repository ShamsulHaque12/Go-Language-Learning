package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	gormlogger "gorm.io/gorm/logger"
)

type Database struct {
	DB    *gorm.DB
	SQLDB *sql.DB
}

// ConnectPostgres initializes a GORM PostgreSQL connection with connection pooling
func ConnectPostgres(databaseURL string, maxOpen, maxIdle int, maxLifetime time.Duration, isProduction bool) (*Database, error) {
	logLevel := gormlogger.Warn
	if !isProduction {
		logLevel = gormlogger.Info
	}

	newLogger := gormlogger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		gormlogger.Config{
			SlowThreshold:             2 * time.Second, // 2s threshold for cross-continental cloud DBs (e.g. Neon US)
			LogLevel:                  logLevel,
			IgnoreRecordNotFoundError: true,
			Colorful:                  true,
		},
	)

	gormConfig := &gorm.Config{
		Logger: newLogger,
	}

	db, err := gorm.Open(postgres.Open(databaseURL), gormConfig)
	if err != nil {
		return nil, fmt.Errorf("failed to open database connection: %w", err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve sql.DB: %w", err)
	}

	// Configure connection pool
	sqlDB.SetMaxOpenConns(maxOpen)
	sqlDB.SetMaxIdleConns(maxIdle)
	sqlDB.SetConnMaxLifetime(maxLifetime)

	// Verify connection
	if err := sqlDB.Ping(); err != nil {
		return nil, fmt.Errorf("database ping failed: %w", err)
	}

	log.Printf("[INFO] Database connected successfully (Pool: MaxOpen=%d, MaxIdle=%d)", maxOpen, maxIdle)
	return &Database{
		DB:    db,
		SQLDB: sqlDB,
	}, nil
}

// Close closes the underlying sql.DB connection safely
func (d *Database) Close() error {
	if d.SQLDB != nil {
		return d.SQLDB.Close()
	}
	return nil
}
