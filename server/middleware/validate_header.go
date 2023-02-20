package middleware

import (
	"github.com/RyaWcksn/fiber-restful/constants"
	"github.com/gofiber/fiber/v2"
)

func ValidateHeader() fiber.Handler {
	return func(c *fiber.Ctx) error {
		ctx := c.Context()

		transactionID := c.Query("transactionId")
		referenceNumber := c.Query("referenceNumber")

		ctx.SetUserValue(constants.CtxTransactionId, transactionID)
		ctx.SetUserValue(constants.CtxReferenceNumber, referenceNumber)

		RequestBody := struct {
			TransactionId   string `json:"transactionId"`
			ReferenceNumber string `json:"referenceNumber"`
		}{}

		if c.Method() != "GET" {
			if err := c.BodyParser(&RequestBody); err != nil {
				return err
			}
			ctx.SetUserValue(constants.CtxTransactionId, RequestBody.TransactionId)
			ctx.SetUserValue(constants.CtxReferenceNumber, RequestBody.ReferenceNumber)
		}

		return c.Next()
	}
}
