package handlers

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http/httptest"
	"testing"

	"bitbucket.org/ayopop/of-core/logger"
	"github.com/RyaWcksn/fiber-restful/apis/v1/services"
	"github.com/RyaWcksn/fiber-restful/entities"
	"github.com/gofiber/fiber/v2"
	"github.com/golang/mock/gomock"
	"github.com/google/go-cmp/cmp"
)

func TestHandlerImpl_Get(t *testing.T) {
	ctrl := gomock.NewController(t)
	serviceMock := services.NewMockIService(ctrl)

	l := logger.New("", "", "debug")
	actualPath := "/api/v1"

	ch := HandlerImpl{
		L:              l,
		serviceUsecase: serviceMock,
	}

	app := fiber.New()
	app.Get(actualPath, ch.Get)

	type args struct {
		c *fiber.Ctx
	}
	tests := []struct {
		name                  string
		method                string
		path                  string
		statusCode            int
		body                  string
		requestBody           map[string]interface{}
		handlerMethodName     string
		handlerToBeCalledWith []interface{}
		requestHeaders        map[string]string
		setMock               func()
	}{
		{
			name:              "Success",
			method:            "GET",
			path:              "/api/v1",
			statusCode:        200,
			body:              `{"code":200,"transactionId":"123","referenceNumber":"123","message":"ok","customer":{"id":1,"name":"arya","status":"ok"}}`,
			requestBody:       nil,
			handlerMethodName: "GetCustomerHandler",
			setMock: func() {
				serviceMock.EXPECT().Get(gomock.Any(), gomock.Any()).Return(&entities.Customer{
					ID:     1,
					Name:   "arya",
					Status: "ok",
				}, nil).Times(1)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.setMock()
			rbody, _ := json.Marshal(tt.requestBody)
			request := httptest.NewRequest(tt.method, tt.path, bytes.NewReader(rbody))
			request.Header.Add(`Content-Type`, `application/json`)

			response, _ := app.Test(request)

			// Validating
			// Status Code
			statusCode := response.StatusCode
			if diff := cmp.Diff(statusCode, tt.statusCode); diff != "" {
				t.Fatalf("\t%s\tStatusCode was incorrect, got: %d, want: %d.", "error", tt.statusCode,
					statusCode)
			}
			t.Logf("\t%s\tShould get statusCode is %v", "ok", statusCode)

			// Response Body
			body, _ := io.ReadAll(response.Body)
			actual := string(body)
			if diff := cmp.Diff(actual, tt.body); diff != "" {
				t.Fatalf("\t%s\tBody was incorrect, got: %v, want: %v", "error", tt.body, actual)
			}
			t.Logf("\t%s\tShould get body is %v", "ok", actual)
		})
	}
}
