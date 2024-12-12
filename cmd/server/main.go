package main

import (
	"customer-api/internal/handler"
	"customer-api/internal/oas"
	"customer-api/internal/service"
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
