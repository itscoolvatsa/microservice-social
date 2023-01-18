package mongodb

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"microservice/posts/pkg/model"
	"time"
)

// Repository defines the MongoDB-based movie user repository.
type Repository struct {
	db *mongo.Collection
}

// New creates a new MongoDB-based collection.
func New(MONGOURI string) *Repository {
	// Set client options
	clientOptions := options.Client().ApplyURI(MONGOURI)

	// Connect to MongoDB
	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		log.Fatal(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	//Check the connection
	err = client.Ping(ctx, nil)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to MongoDB!")

	collection := client.Database("test").Collection("posts")

	return &Repository{
		collection,
	}
}

// AddPost add user to the database collection.
func (r *Repository) AddPost(ctx context.Context, post *model.Post) error {
	_, err := r.db.InsertOne(ctx, post)

	if err != nil {
		return err
	}
	return nil
}
