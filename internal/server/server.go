package server

import (
	"net/http"

	"github.com/carlosEA28/go_chat/internal/config"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
)

type Server struct {
	config *config.Config
	log    *zerolog.Logger
}

func New(cfg *config.Config, log *zerolog.Logger) *Server {
	return &Server{
		config: cfg,
		log:    log,
	}
}

func (s *Server) SetupRoutes() *gin.Engine {
	router := gin.New()

	router.GET("/health", s.healthCheck)

	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	router.Use(s.corsMiddleware())

	return router
}

func (s *Server) healthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": "ok"})
}

func (s *Server) corsMiddleware() gin.HandlerFunc {

	return func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		c.Header("Access-Control-Allow-Headers", "Content-Type, Authorization")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}

}
