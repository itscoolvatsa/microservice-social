package mongodb

import (
	"bytes"
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/gridfs"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"microservice/posts/pkg/model"
	"time"
)

// Repository defines the MongoDB-based movie user repository.
type Repository struct {
	db     *mongo.Collection
	bucket *gridfs.Bucket
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

	imageDatabase := client.Database("media")
	bucket, err := gridfs.NewBucket(imageDatabase)

	if err != nil {
		log.Fatal(err)
	}

	return &Repository{
		collection,
		bucket,
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

// AddMedia add media to the database collection.
func (r *Repository) AddMedia(name string, file []byte) (primitive.ObjectID, error) {

	fid, err := r.bucket.UploadFromStream(name, bytes.NewBuffer(file))
	if err != nil {
		return primitive.ObjectID{}, err
	}
	return fid, nil
}

// GetMedia send media stored from the database.
func (r *Repository) GetMedia(id primitive.ObjectID) (*bytes.Buffer, error) {
	_, cancel := context.WithTimeout(context.Background(), time.Second*60)
	defer cancel()

	//fid, err := r.bucket.UploadFromStream(name, bytes.NewBuffer(file))
	fileBuffer := bytes.NewBuffer(nil)
	_, err := r.bucket.DownloadToStream(id, fileBuffer)

	if err != nil {
		return fileBuffer, err
	}

	return fileBuffer, nil
}

// GetPostsForUser return posts for an user in descending manner
func (r *Repository) GetPostsForUser(id primitive.ObjectID) ([]model.Post, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*60)
	defer cancel()

	cur, err := r.db.Find(ctx, bson.M{"user_id": id})
	if err != nil {
		return nil, err
	}

	var data []model.Post

	err = cur.All(ctx, &data)

	if err != nil {
		return nil, err
	}
	defer cur.Close(ctx)
	return data, nil
}
