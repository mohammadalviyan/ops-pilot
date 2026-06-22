package server

import (
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"

	"opspilot/backend/internal/config"
	"opspilot/backend/internal/middleware"
	"opspilot/backend/internal/module/auth"
	"opspilot/backend/internal/module/health"
)

func NewRouter(cfg config.Config, dbPool *pgxpool.Pool) *gin.Engine {
	if cfg.AppEnv == "production" {
		gin.SetMode(gin.ReleaseMode)
	}

	router := gin.New()
	router.Use(gin.Logger(), gin.Recovery())

	api := router.Group("/api/v1")

	healthHandler := health.NewHandler(cfg)
	health.RegisterRoutes(api, healthHandler)

	authRepository := auth.NewRepository(dbPool)
	tokenManager := auth.NewTokenManager(cfg)
	authService := auth.NewService(authRepository, tokenManager)
	authHandler := auth.NewHandler(authService)
	authMiddleware := middleware.NewAuthMiddleware(tokenManager)
	auth.RegisterRoutes(api, authHandler, authMiddleware)

	return router
}
