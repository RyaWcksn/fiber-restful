package handlers

import (
	"bitbucket.org/ayopop/of-core/tracer"
	"github.com/RyaWcksn/fiber-restful/entities"
	"github.com/RyaWcksn/fiber-restful/forms"
	"github.com/gofiber/fiber/v2"
)

func (h *HandlerImpl) Get(c *fiber.Ctx) error {
	childCtx, span := tracer.StartSpan(c.Context(), "handler.get")
	defer span.End()

	payload := forms.GetCustomerRequest{}

	if err := c.BodyParser(&payload); err != nil {
		h.L.Errorf("[ERROR] Error := %v ", err)
		c.Next()
	}

	customerResp, err := h.serviceUsecase.Get(childCtx, &payload)
	if err != nil {
		h.L.Errorf("[ERROR] err := %v ", err)
		c.Next()
	}

	resp := struct {
		Code            int               `json:"code"`
		TransactionId   string            `json:"transactionId"`
		ReferenceNumber string            `json:"referenceNumber"`
		Message         string            `json:"message"`
		Customer        entities.Customer `json:"customer"`
	}{
		Code:            200,
		TransactionId:   "123",
		ReferenceNumber: "123",
		Message:         "ok",
		Customer:        *customerResp,
	}

	return c.JSON(resp)
}
