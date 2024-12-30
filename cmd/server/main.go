package main

import (
	"github.com/materials-resources/customer-api/internal/handler"
	"github.com/materials-resources/customer-api/internal/oas"
	"github.com/materials-resources/customer-api/internal/service"
	"log"
	"net/http"
)

func main() {
	h := handler.NewHandler(service.NewService())

	srv, err := oas.NewServer(h, handler.NewSecurityHandler())
	if err != nil {
		log.Fatal(err)
	}

	if err := http.ListenAndServe(":8080", srv); err != nil {
		log.Fatal(err)
	}
}
