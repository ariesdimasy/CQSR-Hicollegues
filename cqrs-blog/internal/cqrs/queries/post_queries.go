package queries

// GetPostByIDQuery represents a query to get a post by ID
type GetPostByIDQuery struct {
	ID uint
}

// GetAllPostsQuery represents a query to get all posts
type GetAllPostsQuery struct{}

// GetPostsByUserIDQuery represents a query to get all posts by a user
type GetPostsByUserIDQuery struct {
	UserID uint
}
