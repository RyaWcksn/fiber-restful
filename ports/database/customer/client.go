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

var _ repositories.CRepository = (*CustomerImpl)(nil)
