package middleware

import (
	"errors"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"

	tokenpkg "opspilot/backend/internal/pkg/jwt"
	"opspilot/backend/internal/pkg/response"
)

type AuthVerifier interface {
	Verify(token string) (tokenpkg.Claims, error)
}

func NewAuthMiddleware(verifier AuthVerifier) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		header := ctx.GetHeader("Authorization")
		if header == "" {
			unauthorized(ctx)
			return
		}

		const prefix = "Bearer "
		if !strings.HasPrefix(header, prefix) {
			unauthorized(ctx)
			return
		}

		rawToken := strings.TrimSpace(strings.TrimPrefix(header, prefix))
		claims, err := verifier.Verify(rawToken)
		if err != nil {
			if errors.Is(err, tokenpkg.ErrExpiredToken) {
				response.Error(ctx, http.StatusUnauthorized, "Token expired", nil)
				ctx.Abort()
				return
			}
			unauthorized(ctx)
			return
		}

		tokenpkg.SetClaims(ctx, claims)
		ctx.Next()
	}
}

func unauthorized(ctx *gin.Context) {
	response.Error(ctx, http.StatusUnauthorized, "Unauthorized", nil)
	ctx.Abort()
}
