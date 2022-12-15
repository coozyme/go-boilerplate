package server

import (
	"database/sql"
	"go-mitra-akosta/internal/routes"
	"go-mitra-akosta/pkg/utils/config"

	"github.com/labstack/echo/v4"
)

func (s *Server) InitRoute(cfg *config.Config, dbMysql *sql.DB, route *echo.Echo) {
	routes.InitRouteAuth(cfg, dbMysql, route)
}
