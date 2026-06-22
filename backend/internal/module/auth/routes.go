package auth

import (
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.RouterGroup, handler *Handler, authMiddleware gin.HandlerFunc) {
	authGroup := router.Group("/auth")
	authGroup.POST("/login", handler.Login)
	authGroup.POST("/logout", handler.Logout)
	authGroup.GET("/me", authMiddleware, handler.Me)
}
