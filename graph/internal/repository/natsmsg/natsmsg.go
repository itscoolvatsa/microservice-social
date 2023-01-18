package natsmsg

import (
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

// ReceivesMessage receives the message from auth service to create a new user
func (s *NatServer) ReceivesMessage(data string, cb func(user *NatsUser)) (unsub func() error, err error) {
	res, err := s.server.Subscribe(data, func(user *NatsUser) {
		cb(user)
	})

	if err != nil {
		return nil, err
	}

	return res.Unsubscribe, nil
}
