package auth

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"

	tokenpkg "opspilot/backend/internal/pkg/jwt"
	"opspilot/backend/internal/pkg/response"
)

type Handler struct {
	service *Service
}

func NewHandler(service *Service) *Handler {
	return &Handler{service: service}
}

func (h *Handler) Login(ctx *gin.Context) {
	var request LoginRequest
	if err := ctx.ShouldBindJSON(&request); err != nil {
		response.Error(ctx, http.StatusBadRequest, "Validation error", []gin.H{
			{
				"field":   "email/password",
				"message": "valid email and password are required",
			},
		})
		return
	}

	result, err := h.service.Login(ctx.Request.Context(), request)
	if err != nil {
		h.handleAuthError(ctx, err)
		return
	}

	response.Success(ctx, http.StatusOK, "Login success", result)
}

func (h *Handler) Me(ctx *gin.Context) {
	claims, ok := tokenpkg.ClaimsFromContext(ctx)
	if !ok {
		response.Error(ctx, http.StatusUnauthorized, "Unauthorized", nil)
		return
	}

	user, err := h.service.Me(ctx.Request.Context(), claims.Subject)
	if err != nil {
		h.handleAuthError(ctx, err)
		return
	}

	response.Success(ctx, http.StatusOK, "Success", user)
}

func (h *Handler) Logout(ctx *gin.Context) {
	response.Success(ctx, http.StatusOK, "Logout success", nil)
}

func (h *Handler) handleAuthError(ctx *gin.Context, err error) {
	switch {
	case errors.Is(err, ErrInvalidCredentials):
		response.Error(ctx, http.StatusUnauthorized, "Invalid email or password", nil)
	case errors.Is(err, ErrInactiveUser):
		response.Error(ctx, http.StatusForbidden, "User is inactive", nil)
	case errors.Is(err, ErrMissingJWTSecret):
		response.Error(ctx, http.StatusInternalServerError, "Authentication is not configured", nil)
	default:
		response.Error(ctx, http.StatusInternalServerError, "Internal server error", nil)
	}
}
