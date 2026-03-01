package controllers

import (
	"cqrs-blog/internal/cqrs/commands"
	"cqrs-blog/internal/cqrs/handlers"
	"cqrs-blog/internal/cqrs/queries"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type PostController struct {
	commandHandler *handlers.PostCommandHandler
	queryHandler   *handlers.PostQueryHandler
}

func NewPostController(commandHandler *handlers.PostCommandHandler, queryHandler *handlers.PostQueryHandler) *PostController {
	return &PostController{
		commandHandler: commandHandler,
		queryHandler:   queryHandler,
	}
}

// Create dispatches a CreatePostCommand
func (c *PostController) Create(ctx *gin.Context) {
	var cmd commands.CreatePostCommand
	if err := ctx.ShouldBindJSON(&cmd); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	response, err := c.commandHandler.HandleCreate(cmd)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, response)
}

// GetByID dispatches a GetPostByIDQuery
func (c *PostController) GetByID(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	query := queries.GetPostByIDQuery{ID: uint(id)}
	response, err := c.queryHandler.HandleGetByID(query)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, response)
}

// GetAll dispatches a GetAllPostsQuery
func (c *PostController) GetAll(ctx *gin.Context) {
	query := queries.GetAllPostsQuery{}
	responses, err := c.queryHandler.HandleGetAll(query)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, responses)
}

// GetByUserID dispatches a GetPostsByUserIDQuery
func (c *PostController) GetByUserID(ctx *gin.Context) {
	userID, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid User ID"})
		return
	}

	query := queries.GetPostsByUserIDQuery{UserID: uint(userID)}
	responses, err := c.queryHandler.HandleGetByUserID(query)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, responses)
}

// Update dispatches an UpdatePostCommand
func (c *PostController) Update(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	var cmd commands.UpdatePostCommand
	if err := ctx.ShouldBindJSON(&cmd); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	cmd.ID = uint(id)

	response, err := c.commandHandler.HandleUpdate(cmd)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, response)
}

// Delete dispatches a DeletePostCommand
func (c *PostController) Delete(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	cmd := commands.DeletePostCommand{ID: uint(id)}
	if err := c.commandHandler.HandleDelete(cmd); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusNoContent, nil)
}
