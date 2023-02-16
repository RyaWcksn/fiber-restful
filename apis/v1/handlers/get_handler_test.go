package handlers

import (
	"testing"

	"bitbucket.org/ayopop/of-core/logger"
	"github.com/RyaWcksn/fiber-restful/apis/v1/services"
	"github.com/gofiber/fiber/v2"
)

func TestHandlerImpl_Get(t *testing.T) {
	type fields struct {
		L              logger.ILogger
		serviceUsecase services.IService
	}
	type args struct {
		c *fiber.Ctx
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &HandlerImpl{
				L:              tt.fields.L,
				serviceUsecase: tt.fields.serviceUsecase,
			}
			if err := h.Get(tt.args.c); (err != nil) != tt.wantErr {
				t.Errorf("HandlerImpl.Get() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
