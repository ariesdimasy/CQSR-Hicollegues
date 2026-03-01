package routes

import (
	"event-driven-blog/internal/interfaces/api/controllers"
	"event-driven-blog/internal/interfaces/api/middlewares"

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

	// User routes
	users := api.Group("/users")
	{
		users.POST("/", r.userController.Create)
		users.GET("/", r.userController.GetAll)
		users.GET("/:id", r.userController.GetByID)
		users.PUT("/:id", r.userController.Update)
		users.DELETE("/:id", r.userController.Delete)
		users.GET("/:id/posts", r.postController.GetByUserID)
	}

	// Post routes
	posts := api.Group("/posts")
	{
		posts.POST("/", r.postController.Create)
		posts.GET("/", r.postController.GetAll)
		posts.GET("/:id", r.postController.GetByID)
		posts.PUT("/:id", r.postController.Update)
		posts.DELETE("/:id", r.postController.Delete)
	}

	// Role routes
	roles := api.Group("/roles")
	{
		roles.POST("/", r.roleController.Create)
		roles.GET("/", r.roleController.GetAll)
		roles.GET("/:id", r.roleController.GetByID)
		roles.PUT("/:id", r.roleController.Update)
		roles.DELETE("/:id", r.roleController.Delete)
	}
}
