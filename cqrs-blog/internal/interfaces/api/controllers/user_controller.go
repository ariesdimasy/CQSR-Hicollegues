package controllers

import (
	"cqrs-blog/internal/cqrs/commands"
	"cqrs-blog/internal/cqrs/handlers"
	"cqrs-blog/internal/cqrs/queries"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	commandHandler *handlers.UserCommandHandler
	queryHandler   *handlers.UserQueryHandler
}

func NewUserController(commandHandler *handlers.UserCommandHandler, queryHandler *handlers.UserQueryHandler) *UserController {
	return &UserController{
		commandHandler: commandHandler,
		queryHandler:   queryHandler,
	}
}

// Create dispatches a CreateUserCommand
func (c *UserController) Create(ctx *gin.Context) {
	var cmd commands.CreateUserCommand
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

// GetByID dispatches a GetUserByIDQuery
func (c *UserController) GetByID(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	query := queries.GetUserByIDQuery{ID: uint(id)}
	response, err := c.queryHandler.HandleGetByID(query)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, response)
}

// GetAll dispatches a GetAllUsersQuery
func (c *UserController) GetAll(ctx *gin.Context) {
	query := queries.GetAllUsersQuery{}
	responses, err := c.queryHandler.HandleGetAll(query)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, responses)
}

// Update dispatches an UpdateUserCommand
func (c *UserController) Update(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	var cmd commands.UpdateUserCommand
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

// Delete dispatches a DeleteUserCommand
func (c *UserController) Delete(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	cmd := commands.DeleteUserCommand{ID: uint(id)}
	if err := c.commandHandler.HandleDelete(cmd); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusNoContent, nil)
}
