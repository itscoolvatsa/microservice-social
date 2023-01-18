package controller

import (
	"context"
	"microservice/posts/pkg/model"
)

// postRepository it gives methods that can be implemented by more than one handler.
type postRepository interface {
	AddPost(ctx context.Context, post *model.Post) error
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
