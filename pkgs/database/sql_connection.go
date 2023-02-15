package database

import (
	"database/sql"

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

	dbConn, errConn := sql.Open("mysql", db.MYSQL.Username+":"+db.MYSQL.Password+"@tcp("+db.MYSQL.Address+")/"+db.MYSQL.Database)

	if errConn != nil {
		return nil
	}
	errPing := dbConn.Ping()
	if errPing != nil {
		return nil
	}
	dbConn.SetMaxIdleConns(db.MYSQL.MaxIdleConnections)
	dbConn.SetMaxOpenConns(db.MYSQL.MaxOpenConnections)
	return dbConn
}
