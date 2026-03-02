package controllers

import (
	"cqrs-blog/internal/cqrs/commands"
	"cqrs-blog/internal/cqrs/handlers"
	"cqrs-blog/internal/cqrs/queries"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type RoleController struct {
	commandHandler *handlers.RoleCommandHandler
	queryHandler   *handlers.RoleQueryHandler
}

func NewRoleController(commandHandler *handlers.RoleCommandHandler, queryHandler *handlers.RoleQueryHandler) *RoleController {
	return &RoleController{
		commandHandler: commandHandler,
		queryHandler:   queryHandler,
	}
}

// Create dispatches a CreateRoleCommand
func (c *RoleController) Create(ctx *gin.Context) {
	var cmd commands.CreateRoleCommand
	if err := ctx.ShouldBindJSON(&cmd); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	response, err := c.commandHandler.HandleCreate(cmd)
	if err != nil {
		ctx.JSON(classifyError(err), gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, response)
}

// GetByID dispatches a GetRoleByIDQuery
func (c *RoleController) GetByID(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	query := queries.GetRoleByIDQuery{ID: uint(id)}
	response, err := c.queryHandler.HandleGetByID(query)
	if err != nil {
		ctx.JSON(classifyError(err), gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, response)
}

// GetAll dispatches a GetAllRolesQuery
func (c *RoleController) GetAll(ctx *gin.Context) {
	query := queries.GetAllRolesQuery{}
	responses, err := c.queryHandler.HandleGetAll(query)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, responses)
}

// Update dispatches an UpdateRoleCommand
func (c *RoleController) Update(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	var cmd commands.UpdateRoleCommand
	if err := ctx.ShouldBindJSON(&cmd); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	cmd.ID = uint(id)

	response, err := c.commandHandler.HandleUpdate(cmd)
	if err != nil {
		ctx.JSON(classifyError(err), gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, response)
}

// Delete dispatches a DeleteRoleCommand
func (c *RoleController) Delete(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	cmd := commands.DeleteRoleCommand{ID: uint(id)}
	if err := c.commandHandler.HandleDelete(cmd); err != nil {
		ctx.JSON(classifyError(err), gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusNoContent, nil)
}
