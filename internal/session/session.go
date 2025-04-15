package session

import (
	"github.com/MicahParks/keyfunc/v3"
	"github.com/golang-jwt/jwt/v5"
	"github.com/materials-resources/store-api/app"
	"log"
)

type Manager struct {
	keyFunc jwt.Keyfunc
}

func NewManager(a *app.App) *Manager {
	jwks, err := keyfunc.NewDefault([]string{a.Config.Session.JwksUrl})
	if err != nil {
		log.Fatal(err)
	}
	return &Manager{
		keyFunc: jwks.Keyfunc,
	}
}
