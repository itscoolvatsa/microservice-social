package controller

import (
	"context"
)

// mediaRepository it gives methods that can be implemented by more than one handler.
type mediaRepository interface {
	AddMedia(ctx context.Context, name string, metadata map[string]any, file []byte) error
}

// Controller defines user service controller.
type Controller struct {
	repo mediaRepository
}

// New creates a new user service controller.
func New(repo mediaRepository) *Controller {
	return &Controller{
		repo,
	}
}

// AddMedia add post to the database
func (ctrl *Controller) AddMedia(ctx context.Context, name string, metadata map[string]any, file []byte) error {
	return ctrl.repo.AddMedia(ctx, name, metadata, file)
}
