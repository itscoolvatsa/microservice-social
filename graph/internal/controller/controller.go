package controller

import (
	"context"
	"microservice/graph/internal/repository/natsmsg"
	"microservice/graph/internal/repository/neo4j"
	"microservice/graph/pkg/model"
)

// userRepository it gives methods that can be implemented by more than one handler.
type userRepository interface {
	AddUser(ctx context.Context, user *model.User) (bool, error)
	AddRelationShip(ctx context.Context, followerId string, followingId string) error
}

// natsMsg it gives methods that can be implemented by more than one handler.
type natsMsg interface {
	ReceivesMessage(data string, cb func(user *natsmsg.NatsUser)) (unsub func() error, err error)
}

// Controller defines user service controller.
type Controller struct {
	repo userRepository
	sub  natsMsg
}

// New creates a new user service controller.
func New(repo *neo4j.Repository, natsmsg natsMsg) *Controller {
	return &Controller{
		repo: repo,
		sub:  natsmsg,
	}
}

// AddUser adds a user to the graph.
func (ctrl *Controller) AddUser(ctx context.Context, user *model.User) (bool, error) {
	return ctrl.repo.AddUser(ctx, user)
}

func (ctrl *Controller) AddRelationShip(ctx context.Context, followerId string, followingId string) error {
	return ctrl.repo.AddRelationShip(ctx, followerId, followingId)
}

// ReceivesMessage will receive the message
func (ctrl *Controller) ReceivesMessage(data string, cb func(user *natsmsg.NatsUser)) (unsub func() error, err error) {
	return ctrl.sub.ReceivesMessage(data, cb)
}
