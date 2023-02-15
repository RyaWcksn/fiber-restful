package customer

import (
	"context"
	"database/sql"

	"bitbucket.org/ayopop/of-core/tracer"
	"github.com/RyaWcksn/fiber-restful/entities"
)

func (c *CustomerImpl) Get(ctx context.Context, id int) (resp *entities.Customer, err error) {
	cldCtx, span := tracer.StartSpan(ctx, "port.database.GetOne")
	defer span.End()

	res := entities.Customer{}

	err = c.db.QueryRowContext(cldCtx, `SELECT * FROM customers WHERE id = ?`, id).Scan(&res.ID, &res.Name, &res.Status)
	if err != nil {
		if err == sql.ErrNoRows {
			c.L.Errorf("[ERROR] Err := %v", err)
			return nil, err
		}
		c.L.Errorf("[ERROR] Err := %v", err)
		return nil, err
	}

	return &res, nil
}
