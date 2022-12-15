package main

import (
	db "go-mitra-akosta/infrastructure/database"
	srv "go-mitra-akosta/infrastructure/server"
	"go-mitra-akosta/pkg/utils/config"
	"go-mitra-akosta/pkg/utils/logger"

	"github.com/sirupsen/logrus"
)

func main() {
	logging := logger.NewLogger(logrus.New(), "application.log")

	env := config.NewEnvironment(logging, "config", "../../", "yml")
	v := env.LoadConfig()

	cfg := env.ParseConfig(v)

	logging.InfoLog(cfg.Mysql.MysqlDriver)

	database := db.NewDatabase(cfg, logging)
	mysql := database.Mysql.ConnectDbMysql()

	s := srv.NewServer(mysql, cfg, logging)
	s.Run()
}
