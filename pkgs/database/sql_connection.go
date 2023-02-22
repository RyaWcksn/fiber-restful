package database

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"

	"bitbucket.org/ayopop/of-core/logger"
	"github.com/RyaWcksn/fiber-restful/configs"
)

type Connection struct {
	MYSQL configs.DatabaseConfig
	L     logger.ILogger
}
type DBConnection interface {
	DBConnect()
}

func NewDatabaseConnection(M configs.DatabaseConfig, l logger.ILogger) *Connection {
	return &Connection{MYSQL: M, L: l}
}

func (db *Connection) DBConnect() *sql.DB {

	fmt.Println(db.MYSQL)
	dbConn, errConn := sql.Open("mysql", db.MYSQL.Username+":"+db.MYSQL.Password+"@tcp("+db.MYSQL.Address+")/"+db.MYSQL.Database)

	fmt.Println(errConn)
	if errConn != nil {
		return nil
	}

	fmt.Println("Masuk sini")
	errPing := dbConn.Ping()
	fmt.Println(errPing)
	if errPing != nil {
		return nil
	}
	fmt.Println("err ping :: ", errPing)
	fmt.Println("err conn :: ", errConn)
	dbConn.SetMaxIdleConns(db.MYSQL.MaxIdleConnections)
	dbConn.SetMaxOpenConns(db.MYSQL.MaxOpenConnections)
	return dbConn
}
