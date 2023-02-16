package handlers

import (
	"bitbucket.org/ayopop/of-core/tracer"
	"github.com/gofiber/fiber/v2"
)

func (h *HandlerImpl) Get(c *fiber.Ctx) error {
	_, span := tracer.StartSpan(c.Context(), "handler.get")
	defer span.End()
	return nil
}
