package services

import (
	"context"

	"bitbucket.org/ayopop/of-core/tracer"
	"github.com/RyaWcksn/fiber-restful/entities"
	"github.com/RyaWcksn/fiber-restful/forms"
)

// Get implements IService
func (s *ServiceImpl) Get(ctx context.Context, payload *forms.GetCustomerRequest) (resp *entities.Customer, err error) {
	childCtx, span := tracer.StartSpan(ctx, "service.get")
	defer span.End()

	sqlRes, err := s.customerPort.Get(childCtx, payload.ID)
	if err != nil {
		s.L.Errorf("[ERROR] err := %v ", err)
		return nil, err
	}

	return sqlRes, nil
}
