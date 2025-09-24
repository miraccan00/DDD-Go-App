package httpserver

import (
	"github.com/gofiber/fiber/v3"
)

// RegisterRoutes only registers the K8s endpoints using the handler
func RegisterRoutes(app *fiber.App, handler *K8sHandler) {
	app.Get("/unused-configmaps", handler.GetUnusedConfigMaps)
	app.Get("/unused-secrets", handler.GetUnusedSecrets)
	app.Get("/unused-pvc", handler.GetUnusedPVCs)
}
