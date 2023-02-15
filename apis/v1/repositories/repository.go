package repositories

import (
	"context"

	"github.com/RyaWcksn/fiber-restful/entities"
)

//go:generate mockgen -source repository.go -destination repository_mock.go -package repositories
type CRepository interface {
	Get(ctx context.Context, id int) (resp *entities.Customer, err error)
}
