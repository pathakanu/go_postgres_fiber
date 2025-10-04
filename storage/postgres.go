package storage

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// Config is a struct that holds the database connection configuration.
type Config struct {
	Host     string
	Port     string
	User     string
	DBName   string
	Password string
	SSLMode  string
}

// NewConnection creates a new database connection using the provided configuration.
// It returns a gorm.DB instance and an error if the connection fails.
func NewConnection(config *Config) (*gorm.DB, error) {
	// Create the data source name (DSN) string from the configuration
	dsn := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
		config.Host,
		config.Port,
		config.User,
		config.DBName,
		config.Password,
		config.SSLMode,
	)

	// Open a new database connection
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	return db, nil
}
