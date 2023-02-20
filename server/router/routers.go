package router

import (
	"github.com/RyaWcksn/fiber-restful/apis/v1/handlers"
	"github.com/RyaWcksn/fiber-restful/server/middleware"
	"github.com/gofiber/fiber/v2"
)

// InitiateRouter ...
func InitiateRouter(app *fiber.App, ctrl handlers.IHandler) {
	route := app.Group("/api/v1")

	route.Get("/customers", middleware.ErrHandler(ctrl.Get))
}
