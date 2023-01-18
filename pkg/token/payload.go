package token

import (
	"errors"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

var (
	ErrExpiredToken = errors.New("token has expired")
	ErrInvalidToken = errors.New("token is invalid")
)

// Payload contains the payload data of the token
type Payload struct {
	ID        uuid.UUID          `json:"id"`
	UserId    primitive.ObjectID `json:"user_id"`
	Email     string             `json:"email"`
	IssuedAt  time.Time          `json:"issued_at"`
	ExpiredAt time.Time          `json:"expired_at"`
}

// NewPayload creates a new token payload with a specific id, email and duration
func NewPayload(email string, userId string, duration time.Duration) (*Payload, error) {
	id, err := uuid.NewRandom()

	if err != nil {
		return nil, err
	}

	userIdConv, _ := primitive.ObjectIDFromHex(userId)

	return &Payload{
		ID:        id,
		UserId:    userIdConv,
		Email:     email,
		IssuedAt:  time.Now(),
		ExpiredAt: time.Now().Add(time.Duration(duration)),
	}, nil
}

// Valid checks if the token payload is valid or not
func (payload *Payload) Valid() error {
	if time.Now().After(payload.ExpiredAt) {
		return ErrExpiredToken
	}
	return nil
}
