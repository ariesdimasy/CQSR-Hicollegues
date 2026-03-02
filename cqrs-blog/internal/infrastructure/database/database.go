package database

import (
	"cqrs-blog/internal/domain/post"
	"cqrs-blog/internal/domain/role"
	"cqrs-blog/internal/domain/user"
	"fmt"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewPostgresConnection() (*gorm.DB, error) {
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")

	// Validate required environment variables
	if host == "" || port == "" || dbUser == "" || dbName == "" {
		return nil, fmt.Errorf("missing required database environment variables (DB_HOST, DB_PORT, DB_USER, DB_NAME)")
	}

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Jakarta",
		host, dbUser, password, dbName, port)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}

	log.Println("Database connected successfully")
	return db, nil
}

func AutoMigrate(db *gorm.DB) error {
	err := db.AutoMigrate(
		&role.Role{},
		&user.User{},
		&post.Post{},
	)
	if err != nil {
		return fmt.Errorf("failed to auto-migrate: %w", err)
	}
	log.Println("Database migrated successfully")
	return nil
}
