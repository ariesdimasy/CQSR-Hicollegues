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

func NewPostgresConnection() *gorm.DB {
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Jakarta",
		host, dbUser, password, dbName, port)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	log.Println("Database connected successfully")
	return db
}

func AutoMigrate(db *gorm.DB) {
	err := db.AutoMigrate(
		&role.Role{},
		&user.User{},
		&post.Post{},
	)
	if err != nil {
		log.Fatal("Failed to auto-migrate:", err)
	}
	log.Println("Database migrated successfully")
}
