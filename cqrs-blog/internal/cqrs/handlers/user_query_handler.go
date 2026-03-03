package handlers

import (
	"context"
	"cqrs-blog/internal/cqrs/queries"
	"cqrs-blog/internal/domain/readmodel"
	"cqrs-blog/internal/domain/user"
	"errors"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

// UserQueryHandler handles all user-related queries (Read → MongoDB)
type UserQueryHandler struct {
	mongoDB *mongo.Database
}

// NewUserQueryHandler creates a new UserQueryHandler
func NewUserQueryHandler(mongoDB *mongo.Database) *UserQueryHandler {
	return &UserQueryHandler{mongoDB: mongoDB}
}

// HandleGetByID handles the GetUserByIDQuery
func (h *UserQueryHandler) HandleGetByID(query queries.GetUserByIDQuery) (*user.UserResponse, error) {
	collection := h.mongoDB.Collection("users")

	var result readmodel.UserReadModel
	filter := bson.M{"_id": query.ID}

	err := collection.FindOne(context.Background(), filter).Decode(&result)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, errors.New("user not found")
		}
		return nil, err
	}

	return result.ToUserResponse(), nil
}

// HandleGetAll handles the GetAllUsersQuery
func (h *UserQueryHandler) HandleGetAll(query queries.GetAllUsersQuery) ([]user.UserResponse, error) {
	collection := h.mongoDB.Collection("users")

	cursor, err := collection.Find(context.Background(), bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.Background())

	var results []readmodel.UserReadModel
	if err := cursor.All(context.Background(), &results); err != nil {
		return nil, err
	}

	if results == nil {
		return []user.UserResponse{}, nil
	}

	return readmodel.ToUserResponses(results), nil
}
