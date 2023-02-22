package main

import (
	"fmt"

	"bitbucket.org/ayopop/of-core/logger"
	"github.com/RyaWcksn/fiber-restful/configs"
	"github.com/RyaWcksn/fiber-restful/server"
)

func main() {
	cfg := configs.Cfg

	logger := logger.New(cfg.ENV, cfg.ENV, cfg.LogLevel)
	fmt.Println(cfg.MySql.Password)

	fmt.Println(cfg.ENV)
	sv := server.New(cfg, logger)
	sv.Start()
}
