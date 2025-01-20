package session

import (
	"github.com/MicahParks/keyfunc/v3"
	"github.com/golang-jwt/jwt/v5"
	"log"
)

type Manager struct {
	keyFunc jwt.Keyfunc
}

func NewManager(jwksUrl string) *Manager {
	jwks, err := keyfunc.NewDefault([]string{jwksUrl})
	if err != nil {
		log.Fatal(err)
	}
	return &Manager{
		keyFunc: jwks.Keyfunc,
	}
}
