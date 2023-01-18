package token

import (
	"fmt"
	"github.com/vk-rv/pvx"
	"golang.org/x/crypto/chacha20poly1305"
	"time"
)

type Maker interface {
	// CreateToken creates a new token for a specific username and duration
	CreateToken(email string, id string, duration time.Duration) (string, *Payload, error)

	// VerifyToken checks if the token is valid or not
	VerifyToken(token string) (*Payload, error)
}

type PasetoMaker struct {
	paseto       *pvx.ProtoV4Local
	symmetricKey []byte
}

// New creates a new PasetoMaker
func New(symmetricKey string) (Maker, error) {
	if len(symmetricKey) != chacha20poly1305.KeySize {
		return nil, fmt.Errorf("invalid key size: must be exactly %d characters", chacha20poly1305.KeySize)
	}

	maker := &PasetoMaker{
		paseto:       pvx.NewPV4Local(),
		symmetricKey: []byte(symmetricKey),
	}

	return maker, nil
}

// CreateToken generates a new paseto toke
func (p *PasetoMaker) CreateToken(email string, id string, duration time.Duration) (string, *Payload, error) {
	payload, err := NewPayload(email, id, duration)

	if err != nil {
		return "", payload, err
	}

	key := pvx.NewSymmetricKey(p.symmetricKey, pvx.Version4)
	token, err := p.paseto.Encrypt(key, payload, pvx.WithAssert([]byte("test")))
	//token, err := maker.paseto.Encrypt(maker.symmetricKey, payload, nil)
	return token, payload, err
}

// VerifyToken verifies the token provided and returns the payload data
func (p *PasetoMaker) VerifyToken(token string) (*Payload, error) {
	payload := &Payload{}

	key := pvx.NewSymmetricKey(p.symmetricKey, pvx.Version4)
	err := p.paseto.Decrypt(token, key, pvx.WithAssert([]byte("test"))).ScanClaims(payload)

	if err != nil {
		return nil, err
	}

	err = payload.Valid()
	if err != nil {
		return nil, err
	}

	return payload, nil
}
