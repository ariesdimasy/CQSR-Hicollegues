package main

import (
	"cqrs-blog/internal/cqrs/handlers"
	"cqrs-blog/internal/infrastructure/database"
	"cqrs-blog/internal/infrastructure/sync"
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

	// ==========================================
	// Initialize Write Database (PostgreSQL)
	// ==========================================
	pgDB, err := database.NewPostgresConnection()
	if err != nil {
		log.Fatal("PostgreSQL connection failed:", err)
	}

	// Auto-migrate models (PostgreSQL only)
	if err := database.AutoMigrate(pgDB); err != nil {
		log.Fatal("PostgreSQL migration failed:", err)
	}

	// ==========================================
	// Initialize Read Database (MongoDB)
	// ==========================================
	mongoDB, err := database.NewMongoConnection()
	if err != nil {
		log.Fatal("MongoDB connection failed:", err)
	}

	// ==========================================
	// Initialize Sync Service (PG → MongoDB)
	// ==========================================
	syncService := sync.NewSyncService(mongoDB)

	// Initialize validator
	validate := validator.New()

	// ==========================================
	// CQRS: Initialize Command Handlers (→ PostgreSQL)
	// ==========================================
	userCommandHandler := handlers.NewUserCommandHandler(pgDB, validate, syncService)
	postCommandHandler := handlers.NewPostCommandHandler(pgDB, validate, syncService)
	roleCommandHandler := handlers.NewRoleCommandHandler(pgDB, validate, syncService)

	// ==========================================
	// CQRS: Initialize Query Handlers (→ MongoDB)
	// ==========================================
	userQueryHandler := handlers.NewUserQueryHandler(mongoDB)
	postQueryHandler := handlers.NewPostQueryHandler(mongoDB)
	roleQueryHandler := handlers.NewRoleQueryHandler(mongoDB)

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
	log.Println("Write DB: PostgreSQL | Read DB: MongoDB")
	if err := engine.Run(":" + port); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}
