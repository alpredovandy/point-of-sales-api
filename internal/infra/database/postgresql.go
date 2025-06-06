package database

import (
	"database/sql"
	"fmt"
	"point-of-sales-api/internal/config"
	"time"

	_ "github.com/lib/pq" // PostgreSQL driver
)

func ConnectPostgreSQL(cfg config.DBConfig) (db *sql.DB, err error) {
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		cfg.Host,
		cfg.Port,
		cfg.Username,
		cfg.Password,
		cfg.Database,
		cfg.SSLMode,
	)

	db, err = sql.Open("postgres", dsn)
	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		return nil, err
	}

	// Set connection pool parameters
	db.SetMaxOpenConns(cfg.MaxOpenConns)
	db.SetMaxIdleConns(cfg.MaxIdleConns)
	db.SetConnMaxLifetime(time.Duration(cfg.MaxLifetime) * time.Second)
	db.SetConnMaxIdleTime(time.Duration(cfg.MaxIdleTime) * time.Second)

	return
}
