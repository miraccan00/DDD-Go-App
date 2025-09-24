package main

import (
	"log"

	"github.com/gofiber/fiber/v3"
	"github.com/miraccan00/k8s-cleanup-api/infrastructure/httpserver"
)

func main() {
	app := fiber.New()
	k8sHandler := httpserver.NewK8sHandler()
	httpserver.RegisterRoutes(app, k8sHandler)
	log.Fatal(app.Listen(":8081"))
}
