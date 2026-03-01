package main

import (
	"event-driven-blog/internal/application/services"
	"event-driven-blog/internal/infrastructure/database"
	"event-driven-blog/internal/infrastructure/eventbus"
	"event-driven-blog/internal/infrastructure/validator"
	"event-driven-blog/internal/interfaces/api/controllers"
	"event-driven-blog/internal/interfaces/api/routes"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	// Load environment variables
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}

	// Initialize database
	db := database.NewMySQLConnection()

	// Initialize event bus
	eventBus := eventbus.NewEventBus()

	// Subscribe to events
	eventBus.Subscribe("user.created", eventbus.LogEventHandler)
	eventBus.Subscribe("user.created", eventbus.SendNotificationHandler)
	eventBus.Subscribe("post.created", eventbus.LogEventHandler)

	// Initialize validator
	validate := validator.NewValidator()

	// Initialize services
	userService := services.NewUserService(db, eventBus, validate)
	postService := services.NewPostService(db, eventBus, validate)
	roleService := services.NewRoleService(db, eventBus, validate)

	// Initialize controllers
	userController := controllers.NewUserController(userService)
	postController := controllers.NewPostController(postService)
	roleController := controllers.NewRoleController(roleService)

	// Setup router
	router := routes.NewRouter(userController, postController, roleController)

	// Create gin engine
	engine := gin.Default()
	router.SetupRoutes(engine)

	// Start server
	if err := engine.Run(":8080"); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}
