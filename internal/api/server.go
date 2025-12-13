package api

import (
	"fmt"
	"log"

	"github.com/SunilKividor/PillNet-Backend/internal/config"
	"github.com/gin-gonic/gin"
)

type Server struct {
	Port   string
	Engine *gin.Engine
}

func NewServer(cfg *config.Config) *Server {

	engine := gin.New()
	engine.Use(gin.Logger())

	s := &Server{
		Port:   cfg.ServerConfig.Port,
		Engine: engine,
	}

	return s
}

func (s *Server) Serve() error {

	if s.Port == "" {
		s.Port = "3030"
		log.Printf("[INFO] No port provided. Using default %s\n", s.Port)
	}

	if err := s.Engine.Run(":" + s.Port); err != nil {
		return fmt.Errorf("failed to start server: %w", err)
	}

	return nil
}
