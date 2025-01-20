package main

import (
	"github.com/materials-resources/customer-api/internal/oas"
	"github.com/materials-resources/customer-api/internal/service"
	"github.com/materials-resources/customer-api/internal/session"
	"log"
	"net/http"
)

type application struct {
}

func main() {
	sm := session.NewManager("https://auth.materials-resources.com/oauth/v2/keys")

	h := NewHandler(service.NewService(), sm)

	srv, err := oas.NewServer(h, NewSecurityHandler(sm))
	if err != nil {
		log.Fatal(err)
	}

	if err := http.ListenAndServe(":8080", srv); err != nil {
		log.Fatal(err)
	}
}
