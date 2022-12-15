package server

import (
	"database/sql"
	"fmt"
	"go-mitra-akosta/pkg/utils/config"
	"go-mitra-akosta/pkg/utils/logger"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"golang.org/x/net/http2"
)

type Server struct {
	Mysql  *sql.DB
	Echo   *echo.Echo
	Cfg    *config.Config
	Logger *logger.Logging
}

func NewServer(mysql *sql.DB, cfg *config.Config, log *logger.Logging) *Server {
	return &Server{
		Echo:   echo.New(),
		Mysql:  mysql,
		Cfg:    cfg,
		Logger: log,
	}
}

func (s *Server) Run() error {
	s.InitRoute(s.Cfg, s.Mysql, s.Echo)

	http2 := &http2.Server{
		MaxConcurrentStreams: 250,
		MaxReadFrameSize:     1048576,
		IdleTimeout:          10 * time.Second,
	}

	if error := s.Echo.StartH2CServer(fmt.Sprintf(":%s", s.Cfg.Server.Port), http2); error != http.ErrServerClosed {
		s.Logger.InfoLog(error)
		return error
	}

	return nil
}
