package mongodb

import (
	"bytes"
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/gridfs"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"time"
)

// Repository defines the MongoDB-based movie user repository.
type Repository struct {
	bucket *gridfs.Bucket
}

//database
//client         *Client
//name           string
//readConcern    *readconcern.ReadConcern
//writeConcern   *writeconcern.WriteConcern
//readPreference *readpref.ReadPref
//readSelector   description.ServerSelector
//writeSelector  description.ServerSelector
//registry       *bsoncodec.Registry

//collection
//client         *Client
//db             *Database
//name           string
//readConcern    *readconcern.ReadConcern
//writeConcern   *writeconcern.WriteConcern
//readPreference *readpref.ReadPref
//readSelector   description.ServerSelector
//writeSelector  description.ServerSelector
//registry       *bsoncodec.Registry

// New creates a new MongoDB-based collection.
func New(MONGOURI string) *Repository {
	// Set client options
	clientOptions := options.Client().ApplyURI(MONGOURI)

	// Connect to MongoDB
	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		log.Fatal(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()
	//Check the connection
	err = client.Ping(ctx, nil)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to MongoDB!")

	//collection := client.Database("test").Collection("posts")
	imageDatabase := client.Database("media")
	bucket, err := gridfs.NewBucket(imageDatabase)

	if err != nil {
		log.Fatal(err)
	}

	return &Repository{
		bucket,
	}
}

// AddMedia add user to the database collection.
func (r *Repository) AddMedia(ctx context.Context, name string, metadata map[string]any, file []byte) error {
	//_, err := r.db.InsertOne(ctx, media)

	uploadOpts := options.GridFSUpload().
		SetMetadata(metadata)
	_, err := r.bucket.UploadFromStream(
		name,
		bytes.NewBuffer(file),
		uploadOpts)
	if err != nil {
		log.Fatal(err)
	}

	if err != nil {
		return err
	}
	return nil
}
