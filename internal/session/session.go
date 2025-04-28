package session

import (
	"github.com/MicahParks/keyfunc/v3"
	"github.com/golang-jwt/jwt/v5"
	"github.com/materials-resources/store-api/app"
	"github.com/materials-resources/store-api/internal/zitadel"
	"log"
)

type Manager struct {
	keyFunc       jwt.Keyfunc
	zitadelClient *zitadel.Client
}

func NewManager(a *app.App, zitadelClient *zitadel.Client) *Manager {
	jwks, err := keyfunc.NewDefault([]string{a.Config.Session.JwksUrl})
	if err != nil {
		log.Fatal(err)
	}
	return &Manager{
		keyFunc:       jwks.Keyfunc,
		zitadelClient: zitadelClient,
	}
}
