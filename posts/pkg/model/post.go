package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

// Post defines the user model.
type Post struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	UserId    primitive.ObjectID `bson:"user_id" json:"user_id"`
	Caption   string             `bson:"caption" json:"caption"`
	ImageId   primitive.ObjectID `bson:"image_id" json:"image_id"`
	CreatedAt time.Time          `bson:"created_at" json:"created_at,omitempty"`
	UpdatedAt time.Time          `bson:"updated_at" json:"updated_at,omitempty"`
}
