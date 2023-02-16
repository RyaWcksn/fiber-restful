package customer

import (
	"database/sql"

	"bitbucket.org/ayopop/of-core/logger"
	"github.com/RyaWcksn/fiber-restful/apis/v1/repositories"
)

type CustomerImpl struct {
	L  logger.ILogger
	db *sql.DB
}

// NewCustomer initiate customer database instance.
func NewCustomer(l logger.ILogger, db *sql.DB) *CustomerImpl {
	return &CustomerImpl{
		L:  l,
		db: db,
	}
}

var _ repositories.CRepository = (*CustomerImpl)(nil)
