package sync

import (
	"context"
	"cqrs-blog/internal/domain/post"
	"cqrs-blog/internal/domain/readmodel"
	"cqrs-blog/internal/domain/role"
	"cqrs-blog/internal/domain/user"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

// SyncService handles synchronization from PostgreSQL to MongoDB
type SyncService struct {
	mongoDB *mongo.Database
}

// NewSyncService creates a new SyncService
func NewSyncService(mongoDB *mongo.Database) *SyncService {
	return &SyncService{mongoDB: mongoDB}
}

// SyncRole upserts a Role into MongoDB
func (s *SyncService) SyncRole(r *role.Role) error {
	collection := s.mongoDB.Collection("roles")
	readModel := readmodel.NewRoleReadModel(r)

	filter := bson.M{"_id": readModel.ID}
	update := bson.M{"$set": readModel}
	opts := options.UpdateOne().SetUpsert(true)

	_, err := collection.UpdateOne(context.Background(), filter, update, opts)
	if err != nil {
		return fmt.Errorf("failed to sync role to MongoDB: %w", err)
	}

	log.Printf("Role (ID: %d) synced to MongoDB", r.ID)
	return nil
}

// DeleteRole removes a Role from MongoDB
func (s *SyncService) DeleteRole(id uint) error {
	collection := s.mongoDB.Collection("roles")
	filter := bson.M{"_id": id}

	_, err := collection.DeleteOne(context.Background(), filter)
	if err != nil {
		return fmt.Errorf("failed to delete role from MongoDB: %w", err)
	}

	log.Printf("Role (ID: %d) deleted from MongoDB", id)
	return nil
}

// SyncUser upserts a User into MongoDB
func (s *SyncService) SyncUser(u *user.User) error {
	collection := s.mongoDB.Collection("users")
	readModel := readmodel.NewUserReadModel(u)

	filter := bson.M{"_id": readModel.ID}
	update := bson.M{"$set": readModel}
	opts := options.UpdateOne().SetUpsert(true)

	_, err := collection.UpdateOne(context.Background(), filter, update, opts)
	if err != nil {
		return fmt.Errorf("failed to sync user to MongoDB: %w", err)
	}

	log.Printf("User (ID: %d) synced to MongoDB", u.ID)
	return nil
}

// DeleteUser removes a User from MongoDB
func (s *SyncService) DeleteUser(id uint) error {
	collection := s.mongoDB.Collection("users")
	filter := bson.M{"_id": id}

	_, err := collection.DeleteOne(context.Background(), filter)
	if err != nil {
		return fmt.Errorf("failed to delete user from MongoDB: %w", err)
	}

	log.Printf("User (ID: %d) deleted from MongoDB", id)
	return nil
}

// SyncPost upserts a Post into MongoDB
func (s *SyncService) SyncPost(p *post.Post) error {
	collection := s.mongoDB.Collection("posts")
	readModel := readmodel.NewPostReadModel(p)

	filter := bson.M{"_id": readModel.ID}
	update := bson.M{"$set": readModel}
	opts := options.UpdateOne().SetUpsert(true)

	_, err := collection.UpdateOne(context.Background(), filter, update, opts)
	if err != nil {
		return fmt.Errorf("failed to sync post to MongoDB: %w", err)
	}

	log.Printf("Post (ID: %d) synced to MongoDB", p.ID)
	return nil
}

// DeletePost removes a Post from MongoDB
func (s *SyncService) DeletePost(id uint) error {
	collection := s.mongoDB.Collection("posts")
	filter := bson.M{"_id": id}

	_, err := collection.DeleteOne(context.Background(), filter)
	if err != nil {
		return fmt.Errorf("failed to delete post from MongoDB: %w", err)
	}

	log.Printf("Post (ID: %d) deleted from MongoDB", id)
	return nil
}
