package handlers

import (
	"context"
	"cqrs-blog/internal/cqrs/queries"
	"cqrs-blog/internal/domain/post"
	"cqrs-blog/internal/domain/readmodel"
	"errors"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

// PostQueryHandler handles all post-related queries (Read → MongoDB)
type PostQueryHandler struct {
	mongoDB *mongo.Database
}

// NewPostQueryHandler creates a new PostQueryHandler
func NewPostQueryHandler(mongoDB *mongo.Database) *PostQueryHandler {
	return &PostQueryHandler{mongoDB: mongoDB}
}

// HandleGetByID handles the GetPostByIDQuery
func (h *PostQueryHandler) HandleGetByID(query queries.GetPostByIDQuery) (*post.PostResponse, error) {
	collection := h.mongoDB.Collection("posts")

	var result readmodel.PostReadModel
	filter := bson.M{"_id": query.ID}

	err := collection.FindOne(context.Background(), filter).Decode(&result)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, errors.New("post not found")
		}
		return nil, err
	}

	return result.ToPostResponse(), nil
}

// HandleGetAll handles the GetAllPostsQuery
func (h *PostQueryHandler) HandleGetAll(query queries.GetAllPostsQuery) ([]post.PostResponse, error) {
	collection := h.mongoDB.Collection("posts")

	cursor, err := collection.Find(context.Background(), bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.Background())

	var results []readmodel.PostReadModel
	if err := cursor.All(context.Background(), &results); err != nil {
		return nil, err
	}

	if results == nil {
		return []post.PostResponse{}, nil
	}

	return readmodel.ToPostResponses(results), nil
}

// HandleGetByUserID handles the GetPostsByUserIDQuery
func (h *PostQueryHandler) HandleGetByUserID(query queries.GetPostsByUserIDQuery) ([]post.PostResponse, error) {
	collection := h.mongoDB.Collection("posts")

	filter := bson.M{"user_id": query.UserID}
	cursor, err := collection.Find(context.Background(), filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.Background())

	var results []readmodel.PostReadModel
	if err := cursor.All(context.Background(), &results); err != nil {
		return nil, err
	}

	if results == nil {
		return []post.PostResponse{}, nil
	}

	return readmodel.ToPostResponses(results), nil
}
