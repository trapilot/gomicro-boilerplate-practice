package main

import (
    "log"

    "go-micro.dev/v5"
    "go-micro.dev/v5/logger"
    "github.com/trapilot/gomicro-boilerplate-practice/internal/handler"
)

func main() {
    // Create a new go-micro service
    service := micro.NewService(
        micro.Name("api.gateway"),
        micro.Version("latest"),
    )

    // Initialise service (parse flags etc.)
    service.Init()

    // Register health handler directly on the server
    if err := micro.RegisterHandler(service.Server(), handler.NewHealthHandler()); err != nil {
        logger.Fatal(err)
    }

    // Run the service
    if err := service.Run(); err != nil {
        log.Fatalf("service failed: %v", err)
    }
}
