package middleware

import (
	"runtime"

	"bitbucket.org/ayopop/of-core/logger"
	"github.com/RyaWcksn/fiber-restful/constants"
	"github.com/RyaWcksn/fiber-restful/forms"
	"github.com/gofiber/fiber/v2"
)

type ErrHandler func(*fiber.Ctx) error

func (fn ErrHandler) ServeHTTP(c *fiber.Ctx) {
	defer func() {
		if err := recover(); err != nil {
			var stackSize int = 4 << 10 // 4 KB
			stack := make([]byte, stackSize)
			length := runtime.Stack(stack, true)
			logger.Log.Error(string(stack[:length]))
			xerr := forms.ErrorForm{
				HttpCode: 503,
				Message:  "Error!",
				Error:    "Internal Server Error",
			}
			c.SendStatus(xerr.HttpCode)
			c.JSON(xerr)
		}
	}()
	if err := fn(c); err != nil {
		c.Set(constants.HeaderContentType, constants.MIMEApplicationJson)

		xerr := forms.ErrorForm{
			HttpCode: 503,
			Message:  "Error",
			Error:    err.Error(),
		}
		c.SendStatus(xerr.HttpCode)
		c.JSON(xerr)
		return
	}
}
