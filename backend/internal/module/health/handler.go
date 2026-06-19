package health

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"opspilot/backend/internal/config"
	"opspilot/backend/internal/pkg/response"
)

type Handler struct {
	config config.Config
}

func NewHandler(config config.Config) *Handler {
	return &Handler{config: config}
}

func (h *Handler) Check(ctx *gin.Context) {
	response.Success(ctx, http.StatusOK, "Success", gin.H{
		"status": "ok",
		"env":    h.config.AppEnv,
	})
}
