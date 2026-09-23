package app

import (
	"github.com/labstack/echo/v5"
	"github.com/mamun-jsx/Go-lang-Blog-Apps/config"
	"gorm.io/gorm"
)

type Server struct {
	E   *echo.Echo
	DB  *gorm.DB
	Cfg *config.Config
}

func NewServer(cfg *config.Config, db *gorm.DB) *Server {
	e := echo.New()
	// RegisterRoutes(e, db, cfg)
	return &Server{E: e, DB: db, Cfg: cfg}
}

func (s *Server) Start(addr string) error {
	return s.E.Start(addr)
}
