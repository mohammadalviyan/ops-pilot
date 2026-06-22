package auth

import (
	"context"
	"errors"
	"fmt"
	"strings"
	"time"

	"golang.org/x/crypto/bcrypt"

	"opspilot/backend/internal/config"
	tokenpkg "opspilot/backend/internal/pkg/jwt"
)

const (
	tokenType        = "Bearer"
	tokenTTL         = 24 * time.Hour
	activeUserStatus = "ACTIVE"
)

var (
	ErrInvalidCredentials = errors.New("invalid credentials")
	ErrInactiveUser       = errors.New("inactive user")
	ErrMissingJWTSecret   = errors.New("JWT_SECRET is required")
)

type TokenManager interface {
	Generate(subject string, name string, email string, role string, ttl time.Duration) (string, int64, error)
	Verify(token string) (tokenpkg.Claims, error)
}

type Service struct {
	repository Repository
	tokens     TokenManager
}

func NewService(repository Repository, tokens TokenManager) *Service {
	return &Service{
		repository: repository,
		tokens:     tokens,
	}
}

func NewTokenManager(cfg config.Config) TokenManager {
	return tokenpkg.NewManager(cfg.JWTSecret)
}

func (s *Service) Login(ctx context.Context, request LoginRequest) (LoginResponse, error) {
	user, err := s.repository.FindByEmail(ctx, strings.ToLower(strings.TrimSpace(request.Email)))
	if errors.Is(err, ErrUserNotFound) {
		return LoginResponse{}, ErrInvalidCredentials
	}
	if err != nil {
		return LoginResponse{}, err
	}

	if user.Status != activeUserStatus {
		return LoginResponse{}, ErrInactiveUser
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(request.Password)); err != nil {
		return LoginResponse{}, ErrInvalidCredentials
	}

	authUser := user.ToAuthUser()
	accessToken, expiresIn, err := s.tokens.Generate(
		authUser.ID,
		authUser.Name,
		authUser.Email,
		authUser.Role,
		tokenTTL,
	)
	if err != nil {
		if errors.Is(err, tokenpkg.ErrMissingSecret) {
			return LoginResponse{}, ErrMissingJWTSecret
		}
		return LoginResponse{}, fmt.Errorf("generate access token: %w", err)
	}

	return LoginResponse{
		AccessToken: accessToken,
		TokenType:   tokenType,
		ExpiresIn:   expiresIn,
		User:        authUser,
	}, nil
}

func (s *Service) Me(ctx context.Context, userID string) (AuthUser, error) {
	user, err := s.repository.FindByID(ctx, userID)
	if errors.Is(err, ErrUserNotFound) {
		return AuthUser{}, ErrInvalidCredentials
	}
	if err != nil {
		return AuthUser{}, err
	}
	if user.Status != activeUserStatus {
		return AuthUser{}, ErrInactiveUser
	}

	return user.ToAuthUser(), nil
}
