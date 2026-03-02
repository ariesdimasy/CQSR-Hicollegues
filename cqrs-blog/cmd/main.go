package main

import (
	"cqrs-blog/internal/cqrs/handlers"
	"cqrs-blog/internal/infrastructure/database"
	"cqrs-blog/internal/interfaces/api/controllers"
	"cqrs-blog/internal/interfaces/api/routes"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/joho/godotenv"
)

func main() {
	// Load environment variables
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}

	// Initialize database
	db, err := database.NewPostgresConnection()
	if err != nil {
		log.Fatal(err)
	}

	// Auto-migrate models
	if err := database.AutoMigrate(db); err != nil {
		log.Fatal(err)
	}

	// Initialize validator
	validate := validator.New()

	// ==========================================
	// CQRS: Initialize Command & Query Handlers
	// ==========================================

	// User handlers
	userCommandHandler := handlers.NewUserCommandHandler(db, validate)
	userQueryHandler := handlers.NewUserQueryHandler(db)

	// Post handlers
	postCommandHandler := handlers.NewPostCommandHandler(db, validate)
	postQueryHandler := handlers.NewPostQueryHandler(db)

	// Role handlers
	roleCommandHandler := handlers.NewRoleCommandHandler(db, validate)
	roleQueryHandler := handlers.NewRoleQueryHandler(db)

	// ==========================================
	// Initialize Controllers (with CQRS handlers)
	// ==========================================
	userController := controllers.NewUserController(userCommandHandler, userQueryHandler)
	postController := controllers.NewPostController(postCommandHandler, postQueryHandler)
	roleController := controllers.NewRoleController(roleCommandHandler, roleQueryHandler)

	// Setup router
	router := routes.NewRouter(userController, postController, roleController)

	// Create gin engine
	engine := gin.Default()
	router.SetupRoutes(engine)

	// Start server
	port := os.Getenv("APP_PORT")
	if port == "" {
		port = "8080"
	}
	log.Printf("Server starting on :%s\n", port)
	if err := engine.Run(":" + port); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}
