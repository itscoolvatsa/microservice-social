package controller

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"microservice/auth/pkg/model"
)

// userRepository it gives methods that can be implemented by more than one handler.
type userRepository interface {
	AddUser(ctx context.Context, user *model.User) (*mongo.InsertOneResult, error)
	FindUser(ctx context.Context, id string) (*model.User, error)
	CountUser(ctx context.Context, email string) (int64, error)
}

// Controller defines user service controller.
type Controller struct {
	repo userRepository
}

// New creates a new user service controller.
func New(repo userRepository) *Controller {
	return &Controller{
		repo,
	}
}

// AddUser adds a user to the database.
func (ctrl *Controller) AddUser(ctx context.Context, user *model.User) (*mongo.InsertOneResult, error) {
	return ctrl.repo.AddUser(ctx, user)
}

// FindUser return the user data // if exists.
func (ctrl *Controller) FindUser(ctx context.Context, email string) (*model.User, error) {
	return ctrl.repo.FindUser(ctx, email)
}

// CountUser returns the count of the user.
func (ctrl *Controller) CountUser(ctx context.Context, email string) (int64, error) {
	return ctrl.repo.CountUser(ctx, email)
}
