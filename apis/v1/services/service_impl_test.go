package services

import (
	"context"
	"fmt"
	"reflect"
	"testing"

	"bitbucket.org/ayopop/of-core/logger"
	"github.com/RyaWcksn/fiber-restful/apis/v1/repositories"
	"github.com/RyaWcksn/fiber-restful/entities"
	"github.com/RyaWcksn/fiber-restful/forms"
	"github.com/golang/mock/gomock"
)

func TestServiceImpl_Get(t *testing.T) {
	ctrl := gomock.NewController(t)

	customerMock := repositories.NewMockCRepository(ctrl)
	ctx := context.Background()

	type args struct {
		ctx     context.Context
		payload *forms.GetCustomerRequest
	}
	tests := []struct {
		name     string
		args     args
		wantMock func()
		wantResp *entities.Customer
		wantErr  bool
	}{
		{
			name: "Succes",
			args: args{
				ctx: ctx,
				payload: &forms.GetCustomerRequest{
					ID:     1,
					Name:   "Arya",
					Status: "Active",
				},
			},
			wantMock: func() {
				customerMock.EXPECT().Get(gomock.Any(), gomock.Any()).Return(&entities.Customer{
					ID:     1,
					Name:   "Arya",
					Status: "Active",
				}, nil).Times(1)
			},
			wantResp: &entities.Customer{
				ID:     1,
				Name:   "Arya",
				Status: "Active",
			},
			wantErr: false,
		},
		{
			name: "Failed - no data",
			args: args{
				ctx: ctx,
				payload: &forms.GetCustomerRequest{
					ID:     1,
					Name:   "Arya",
					Status: "Active",
				},
			},
			wantMock: func() {
				customerMock.EXPECT().Get(gomock.Any(), gomock.Any()).Return(nil, fmt.Errorf("[ERROR] err := %v", "err")).Times(1)
			},
			wantResp: nil,
			wantErr:  true,
		},
	}
	l := logger.New("", "", "debug")
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.wantMock()
			s := NewService(customerMock, l)
			gotResp, err := s.Get(tt.args.ctx, tt.args.payload)
			if (err != nil) != tt.wantErr {
				t.Errorf("ServiceImpl.Get() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotResp, tt.wantResp) {
				t.Errorf("ServiceImpl.Get() = %v, want %v", gotResp, tt.wantResp)
			}
		})
	}
}
