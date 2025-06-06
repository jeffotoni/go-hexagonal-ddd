package main

import (
	"log"

	"github.com/jeffotoni/quick"

	"api/internal/adapters/inbound/http"
	"api/internal/adapters/outbound/postgres"
	"api/internal/application"
)

func main() {
	// Repository instance (mock in memory)
	repo := postgres.NewInMemoryItemRepository()

	// Use cases
	createUC := application.NewCreateItemUseCase(repo)
	getUC := application.NewGetItemUseCase(repo)

	// HTTP handler
	handler := http.NewItemHandler(createUC, getUC)

	// Initialize the Quick server
	q := quick.New()

	// Register routes
	handler.RegisterRoutes(q)

	// Start the server
	if err := q.Listen(":8080"); err != nil {
		log.Fatal(err)
	}
}
