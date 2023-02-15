package customer

import (
	"context"
	"database/sql"
	"fmt"
	"reflect"
	"testing"

	"bitbucket.org/ayopop/of-core/logger"
	"github.com/RyaWcksn/fiber-restful/configs"
	"github.com/RyaWcksn/fiber-restful/entities"
	"github.com/RyaWcksn/fiber-restful/pkgs/containers"
	"github.com/RyaWcksn/fiber-restful/pkgs/database"
	_ "github.com/go-sql-driver/mysql"
)

func TestCustomerImpl_Get(t *testing.T) {

	ctx := context.Background()

	mysqlContainer, err := containers.SetupMysqlContainer(ctx)
	if err != nil {
		t.Fatal(err)
	}

	// Clean up the container after test is complete
	t.Cleanup(func() {
		if err := mysqlContainer.Terminate(ctx); err != nil {
			t.Fatalf("failed to terminate container: %s", err)
		}
	})

	// Assertion
	host, _ := mysqlContainer.Host(ctx)
	p, _ := mysqlContainer.MappedPort(ctx, "3306/tcp")
	port := p.Int()

	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?tls=skip-verify",
		"root", "password", host, port, "database")

	db, err := sql.Open("mysql", connectionString)
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()

	if err = db.Ping(); err != nil {
		t.Errorf("error pinging db: %+v\n", err)
	}

	_, err = db.Exec(`
CREATE TABLE IF NOT EXISTS customers (
  id INT PRIMARY KEY,
  name VARCHAR(255),
  status VARCHAR(255)
);`)
	if err != nil {
		t.Errorf("error creating table: %+v\n", err)
	}

	_, err = db.Exec(`
INSERT INTO customers
VALUES
(1, "Arya", "Active")
;`)
	if err != nil {
		t.Errorf("error add data to table: %+v\n", err)
	}

	l := logger.New("", "", "debug")
	address := fmt.Sprintf("%s:%d", host, port)

	dbConf := configs.DatabaseConfig{
		Username:           "root",
		Password:           "password",
		Protocol:           "tcp",
		Address:            address,
		Database:           "database",
		MaxIdleConnections: 1,
		MaxOpenConnections: 1,
	}

	mysqlInstance := database.NewDatabaseConnection(dbConf, l)
	connection := mysqlInstance.DBConnect()
	type args struct {
		ctx context.Context
		id  int
	}
	tests := []struct {
		name     string
		args     args
		wantDb   *sql.DB
		wantResp *entities.Customer
		wantErr  bool
	}{
		{
			name: "Success",
			args: args{
				ctx: ctx,
				id:  1,
			},
			wantDb: connection,
			wantResp: &entities.Customer{
				ID:     1,
				Name:   "Arya",
				Status: "Active",
			},
			wantErr: false,
		},
		{
			name: "No row",
			args: args{
				ctx: ctx,
				id:  2,
			},
			wantDb:   connection,
			wantResp: nil,
			wantErr:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &CustomerImpl{
				L:  l,
				db: tt.wantDb,
			}
			gotResp, err := c.Get(tt.args.ctx, tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("CustomerImpl.Get() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotResp, tt.wantResp) {
				t.Errorf("CustomerImpl.Get() = %v, want %v", gotResp, tt.wantResp)
			}
		})
	}
}
