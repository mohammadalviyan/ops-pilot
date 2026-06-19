package server

import (
	"github.com/gin-gonic/gin"

	"opspilot/backend/internal/config"
	"opspilot/backend/internal/module/health"
)

func NewRouter(cfg config.Config) *gin.Engine {
	if cfg.AppEnv == "production" {
		gin.SetMode(gin.ReleaseMode)
	}

	router := gin.New()
	router.Use(gin.Logger(), gin.Recovery())

	api := router.Group("/api/v1")

	healthHandler := health.NewHandler(cfg)
	health.RegisterRoutes(api, healthHandler)

	return router
}
