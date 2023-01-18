package natsmsg

import (
	"fmt"
	"github.com/nats-io/nats.go"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"log"
)

// NatServer creates a new type to handle nats server and messages
type NatServer struct {
	server *nats.EncodedConn
}

// New creates a new nats server instance
func New(NATS_URL string) *NatServer {
	nc, err := nats.Connect(NATS_URL)

	if err != nil {
		log.Panic(err)
	}

	c, err := nats.NewEncodedConn(nc, nats.JSON_ENCODER)
	if err != nil {
		log.Panic(err)
	}

	return &NatServer{
		c,
	}
}

// NatsUser user struct to send over the nats
type NatsUser struct {
	Name   string             `json:"name"`
	UserId primitive.ObjectID `json:"user_id"`
}

// SendMessage sends the message to the other services
func (s NatServer) SendMessage(user NatsUser) {
	// Publish the message
	if err := s.server.Publish("user", &user); err != nil {
		log.Fatal(err)
	}
}

func (s NatServer) ReceivesMessage() {
	// Publish the message
	s.server.Subscribe("user", func(m *NatsUser) {
		fmt.Printf("Received a message: %s\n", string(m.Name))
	})

}
