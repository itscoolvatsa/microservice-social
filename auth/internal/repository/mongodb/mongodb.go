package mongodb

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"microservice/auth/pkg/model"
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

	collection := client.Database("test").Collection("user")

	return &Repository{
		collection,
	}
}

// AddUser add user to the database collection.
func (r *Repository) AddUser(ctx context.Context, user *model.User) (*mongo.InsertOneResult, error) {
	res, err := r.db.InsertOne(ctx, user)

	if err != nil {
		return nil, err
	}
	return res, nil
}

// FindUser find the existing user inside the collection and returns it.
func (r *Repository) FindUser(ctx context.Context, email string) (*model.User, error) {
	var user *model.User

	err := r.db.FindOne(ctx, bson.M{"email": email}).Decode(&user)

	if err != nil {
		return nil, err
	}

	return user, nil
}

// CountUser returns the count of the user inside the database.
func (r *Repository) CountUser(ctx context.Context, email string) (int64, error) {
	count, err := r.db.CountDocuments(ctx, bson.M{"email": email})
	if err != nil {
		return 0, err
	}

	return count, nil
}
