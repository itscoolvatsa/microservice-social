package controller

import (
	"bytes"
	"context"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"microservice/posts/pkg/model"
)

// postRepository it gives methods that can be implemented by more than one handler.
type postRepository interface {
	AddPost(ctx context.Context, post *model.Post) error
	AddMedia(name string, file []byte) (primitive.ObjectID, error)
	GetMedia(id primitive.ObjectID) (*bytes.Buffer, error)
	GetPostsForUser(id primitive.ObjectID) ([]model.Post, error)
}

// Controller defines user service controller.
type Controller struct {
	repo postRepository
}

// New creates a new user service controller.
func New(repo postRepository) *Controller {
	return &Controller{
		repo,
	}
}

// AddPost add post to the database
func (ctrl *Controller) AddPost(ctx context.Context, post *model.Post) error {
	return ctrl.repo.AddPost(ctx, post)
}

// AddMedia add media for creating a post
func (ctrl *Controller) AddMedia(name string, file []byte) (primitive.ObjectID, error) {
	return ctrl.repo.AddMedia(name, file)
}

// GetPostsForUser returns the posts for a user based on their id
func (ctrl *Controller) GetPostsForUser(id primitive.ObjectID) ([]model.Post, error) {
	return ctrl.repo.GetPostsForUser(id)
}

func (ctrl *Controller) GetMedia(id primitive.ObjectID) (*bytes.Buffer, error) {
	return ctrl.repo.GetMedia(id)
}
