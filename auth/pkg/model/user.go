package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

// User defines the user model.
type User struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	Name      string             `bson:"name" json:"name" validate:"required,min=4"`
	Email     string             `bson:"email,email" json:"email" validate:"required,email"`
	Password  string             `bson:"password" json:"password"  validate:"required,min=6"`
	CreatedAt time.Time          `bson:"created_at" json:"created_at,omitempty"`
	UpdatedAt time.Time          `bson:"updated_at" json:"updated_at,omitempty"`
}
