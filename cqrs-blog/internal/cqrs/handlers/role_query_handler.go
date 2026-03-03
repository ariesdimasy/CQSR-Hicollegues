package handlers

import (
	"context"
	"cqrs-blog/internal/cqrs/queries"
	"cqrs-blog/internal/domain/readmodel"
	"cqrs-blog/internal/domain/role"
	"errors"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

// RoleQueryHandler handles all role-related queries (Read → MongoDB)
type RoleQueryHandler struct {
	mongoDB *mongo.Database
}

// NewRoleQueryHandler creates a new RoleQueryHandler
func NewRoleQueryHandler(mongoDB *mongo.Database) *RoleQueryHandler {
	return &RoleQueryHandler{mongoDB: mongoDB}
}

// HandleGetByID handles the GetRoleByIDQuery
func (h *RoleQueryHandler) HandleGetByID(query queries.GetRoleByIDQuery) (*role.RoleResponse, error) {
	collection := h.mongoDB.Collection("roles")

	var result readmodel.RoleReadModel
	filter := bson.M{"_id": query.ID}

	err := collection.FindOne(context.Background(), filter).Decode(&result)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, errors.New("role not found")
		}
		return nil, err
	}

	return result.ToRoleResponse(), nil
}

// HandleGetAll handles the GetAllRolesQuery
func (h *RoleQueryHandler) HandleGetAll(query queries.GetAllRolesQuery) ([]role.RoleResponse, error) {
	collection := h.mongoDB.Collection("roles")

	cursor, err := collection.Find(context.Background(), bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.Background())

	var results []readmodel.RoleReadModel
	if err := cursor.All(context.Background(), &results); err != nil {
		return nil, err
	}

	if results == nil {
		return []role.RoleResponse{}, nil
	}

	return readmodel.ToRoleResponses(results), nil
}
