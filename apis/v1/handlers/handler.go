package handlers

import (
	"bitbucket.org/ayopop/of-core/logger"
	"github.com/RyaWcksn/fiber-restful/apis/v1/services"
	"github.com/gofiber/fiber/v2"
)

type HandlerImpl struct {
	L              logger.ILogger
	serviceUsecase services.IService
}

// NewHandler initiate handler for getting the request.
func NewHandler(l logger.ILogger, s services.IService) *HandlerImpl {
	return &HandlerImpl{
		L:              l,
		serviceUsecase: s,
	}
}

type IHandler interface {
	Get(c *fiber.Ctx) error
}

var _ IHandler = (*HandlerImpl)(nil)
