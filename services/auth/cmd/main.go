package main

import (
    "log"

    micro "go-micro.dev/v5"
    "go-micro.dev/v5/server"
    "github.com/trapilot/gomicro-boilerplate-practice/internal/handler"
    "github.com/trapilot/gomicro-boilerplate-practice/internal/repository"
)

func main() {
    // Create a new go-micro service
    service := micro.NewService(
        micro.Name("auth.service"),
    )
    service.Init()

    // Set up dependencies
    repo := repository.NewInMemoryUserRepo()
    authHandler := handler.NewAuthHandler(repo)

    // Register the handler
    if err := micro.RegisterHandler(service.Server(), authHandler); err != nil {
        log.Fatalf("Failed to register handler: %v", err)
    }

    // Run the service (blocks)
    if err := service.Run(); err != nil {
        log.Fatalf("Service run error: %v", err)
    }
}
