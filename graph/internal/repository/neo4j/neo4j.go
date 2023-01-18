package neo4j

import (
	"context"
	"fmt"
	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
	"microservice/graph/pkg/model"
)

type Repository struct {
	driver neo4j.DriverWithContext
}

// New creating a new repository to execute queries
func New(URI string) *Repository {
	driver, err := neo4j.NewDriverWithContext(URI, neo4j.NoAuth())

	if err != nil {
		panic(err)
	}

	return &Repository{
		driver,
	}
}

// AddRelationShip adding a new relationship for following another user
func (r *Repository) AddRelationShip(ctx context.Context, followerId string, followingId string) error {
	//	TODO
	session := r.driver.NewSession(ctx, neo4j.SessionConfig{AccessMode: neo4j.AccessModeWrite})
	defer session.Close(ctx)

	//(lebron:PLAYER{name:"Russell Westbrook", age: 33, number: 0, height: 1.91, weight: 91})

	_, err := session.ExecuteWrite(ctx, func(transaction neo4j.ManagedTransaction) (any, error) {
		result, err := transaction.Run(ctx,
			"MATCH (a:User), (b:User) WHERE a.user_id = $follower_id AND b.user_id = $following_id CREATE (a)-[:FOLLOWS]->(b)", map[string]any{
				"follower_id":  followerId,
				"following_id": followingId,
			})
		if err != nil {
			return nil, err
		}

		if result.Next(ctx) {
			return result.Record().Values[0], nil
		}

		return nil, result.Err()
	})
	if err != nil {
		return err
	}
	return nil
}

// AddUser adding the new user that can be used to establish link t other user
func (r *Repository) AddUser(ctx context.Context, user *model.User) (bool, error) {
	session := r.driver.NewSession(ctx, neo4j.SessionConfig{AccessMode: neo4j.AccessModeWrite})
	defer session.Close(ctx)

	newUser, err := session.ExecuteWrite(ctx, func(transaction neo4j.ManagedTransaction) (any, error) {
		result, err := transaction.Run(ctx,
			"CREATE (u:User {user_id: $user_id , name: $name })", map[string]any{
				"user_id": user.UserId,
				"name":    user.Name,
			})
		if err != nil {
			return nil, err
		}

		if result.Next(ctx) {
			return result.Record().Values[0], nil
		}

		return nil, result.Err()
	})
	if err != nil {
		return false, err
	}

	fmt.Println(newUser)

	return true, nil
}
