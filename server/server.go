package server

import (
	"database/sql"
	"fmt"
	"os"
	"strconv"

	"bitbucket.org/ayopop/of-core/logger"
	"github.com/RyaWcksn/fiber-restful/apis/v1/handlers"
	"github.com/RyaWcksn/fiber-restful/apis/v1/services"
	"github.com/RyaWcksn/fiber-restful/configs"
	"github.com/RyaWcksn/fiber-restful/pkgs/database"
	"github.com/RyaWcksn/fiber-restful/ports/database/customer"
	"github.com/RyaWcksn/fiber-restful/server/router"
	"github.com/gofiber/fiber/v2"
)

type Server struct {
	cfg     *configs.Config
	logger  logger.ILogger
	service services.IService
	handler handlers.IHandler
}

var addr string
var SVR *Server
var db *sql.DB
var signalChan chan (os.Signal) = make(chan os.Signal, 1)

func init() {
	addr = ":9000"
	cfg := configs.Config{}
	if len(cfg.HTTPAddress) > 0 {
		if _, err := strconv.Atoi(cfg.HTTPAddress); err == nil {
			addr = fmt.Sprintf(":%v", cfg.HTTPAddress)
		} else {
			addr = cfg.HTTPAddress
		}
	}
}

// Register all instances.
func (s *Server) Register() {
	dbConn := database.NewDatabaseConnection(s.cfg.DBConf, s.logger)
	if dbConn == nil {
		s.logger.Fatal("Expecting db connection object but received nil")

	}

	db := dbConn.DBConnect()
	if db == nil {
		s.logger.Fatal("Expecting db connection object but received nil")

	}

	customerPort := customer.NewCustomer(s.logger, db)

	s.service = services.NewService(customerPort, s.logger)
	s.handler = handlers.NewHandler(s.logger, s.service)
}

func New(cfg *configs.Config, logger logger.ILogger) *Server {
	if SVR != nil {
		return SVR
	}
	SVR = &Server{
		cfg:    cfg,
		logger: logger,
	}

	SVR.Register()

	return SVR
}

func (s Server) Start() {
	// init tracer

	app := fiber.New()
	router.InitiateRouter(app)
}
