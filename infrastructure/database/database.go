package database

import (
	"database/sql"
	"fmt"
	"go-mitra-akosta/pkg/utils/config"
	"go-mitra-akosta/pkg/utils/logger"

	_ "github.com/go-sql-driver/mysql"
)

type Database struct {
	Mysql *dbMysql
}
type dbMysql struct {
	Cfg *config.Config
	Log *logger.Logging
}

func NewDatabase(cfg *config.Config, log *logger.Logging) *Database {
	return &Database{
		Mysql: &dbMysql{
			Cfg: cfg,
			Log: log,
		},
	}
}

func (m *dbMysql) ConnectDbMysql() *sql.DB {
	dtSoure := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", m.Cfg.Mysql.MysqlUserName, m.Cfg.Mysql.MysqlPassword, m.Cfg.Mysql.MysqlHost, m.Cfg.Mysql.MysqlPort, m.Cfg.Mysql.MysqlName)
	db, error := sql.Open("mysql", dtSoure)
	m.Log.InfoLog(m.Cfg.Mysql.MysqlUserName, m.Cfg.Mysql.MysqlPassword, m.Cfg.Mysql.MysqlHost, m.Cfg.Mysql.MysqlPort, m.Cfg.Mysql.MysqlName)
	if error != nil {
		m.Log.ErrorLog("error open db mysql: ", error)
		return nil
	}

	error = db.Ping()
	if error != nil {
		m.Log.ErrorLog("failed connect database mysql")
		return nil
	}

	error = db.Close()
	if error != nil {
		m.Log.ErrorLog("close database mysql")
		return nil
	}

	m.Log.InfoLog("success connect database mysql")

	return db
}
