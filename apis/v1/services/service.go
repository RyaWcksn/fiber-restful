package services

import (
	"context"

	"bitbucket.org/ayopop/of-core/logger"
	"github.com/RyaWcksn/fiber-restful/apis/v1/repositories"
	"github.com/RyaWcksn/fiber-restful/entities"
	"github.com/RyaWcksn/fiber-restful/forms"
)

type ServiceImpl struct {
	L            logger.ILogger
	customerPort repositories.CRepository
}

//go:generate mockgen -source service.go -destination service_mock.go -package services
type IService interface {
	Get(ctx context.Context, payload *forms.GetCustomerRequest) (resp *entities.Customer, err error)
}

// NewService initiate service.
func NewService(c repositories.CRepository, l logger.ILogger) *ServiceImpl {
	return &ServiceImpl{
		customerPort: c,
		L:            l,
	}
}

var _ IService = (*ServiceImpl)(nil)
