package router

import (
	"github.com/RyaWcksn/fiber-restful/apis/v1/handlers"
	"github.com/gofiber/fiber/v2"
)

// InitiateRouter ...
func InitiateRouter(app *fiber.App, ctrl handlers.IHandler) {
	route := app.Group("/api/v1")

	route.Get("/customers", ctrl.Get)
}
