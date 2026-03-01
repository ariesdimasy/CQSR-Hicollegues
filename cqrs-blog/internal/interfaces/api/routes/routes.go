package routes

import (
	"cqrs-blog/internal/interfaces/api/controllers"
	"cqrs-blog/internal/interfaces/api/middlewares"

	"github.com/gin-gonic/gin"
)

type Router struct {
	userController *controllers.UserController
	postController *controllers.PostController
	roleController *controllers.RoleController
}

func NewRouter(
	userController *controllers.UserController,
	postController *controllers.PostController,
	roleController *controllers.RoleController,
) *Router {
	return &Router{
		userController: userController,
		postController: postController,
		roleController: roleController,
	}
}

func (r *Router) SetupRoutes(engine *gin.Engine) {
	api := engine.Group("/api/v1")
	api.Use(middlewares.Logger())

	// User routes (Commands + Queries)
	users := api.Group("/users")
	{
		users.POST("/", r.userController.Create)              // Command
		users.GET("/", r.userController.GetAll)               // Query
		users.GET("/:id", r.userController.GetByID)           // Query
		users.PUT("/:id", r.userController.Update)            // Command
		users.DELETE("/:id", r.userController.Delete)         // Command
		users.GET("/:id/posts", r.postController.GetByUserID) // Query
	}

	// Post routes (Commands + Queries)
	posts := api.Group("/posts")
	{
		posts.POST("/", r.postController.Create)      // Command
		posts.GET("/", r.postController.GetAll)       // Query
		posts.GET("/:id", r.postController.GetByID)   // Query
		posts.PUT("/:id", r.postController.Update)    // Command
		posts.DELETE("/:id", r.postController.Delete) // Command
	}

	// Role routes (Commands + Queries)
	roles := api.Group("/roles")
	{
		roles.POST("/", r.roleController.Create)      // Command
		roles.GET("/", r.roleController.GetAll)       // Query
		roles.GET("/:id", r.roleController.GetByID)   // Query
		roles.PUT("/:id", r.roleController.Update)    // Command
		roles.DELETE("/:id", r.roleController.Delete) // Command
	}
}
