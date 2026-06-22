package auth

import (
	"context"
	"errors"
	"testing"
	"time"

	tokenpkg "opspilot/backend/internal/pkg/jwt"
)

const seedPasswordHash = "$2a$10$BfCNugPMPzHegtgQnIpjV.Nn3m0HbmxjbRPjccfP8JMY8OtimOiAq"

type fakeRepository struct {
	user User
	err  error
}

func (r fakeRepository) FindByEmail(context.Context, string) (User, error) {
	if r.err != nil {
		return User{}, r.err
	}
	return r.user, nil
}

func (r fakeRepository) FindByID(context.Context, string) (User, error) {
	if r.err != nil {
		return User{}, r.err
	}
	return r.user, nil
}

type fakeTokenManager struct{}

func (fakeTokenManager) Generate(string, string, string, string, time.Duration) (string, int64, error) {
	return "jwt-token", int64(tokenTTL.Seconds()), nil
}

func (fakeTokenManager) Verify(string) (tokenpkg.Claims, error) {
	return tokenpkg.Claims{}, nil
}

func TestLoginWithSeedPassword(t *testing.T) {
	service := NewService(fakeRepository{user: User{
		ID:           "user-id",
		Name:         "RPA Support User",
		Email:        "support@opspilot.local",
		PasswordHash: seedPasswordHash,
		Status:       activeUserStatus,
		RoleCode:     "SUPPORT",
	}}, fakeTokenManager{})

	result, err := service.Login(context.Background(), LoginRequest{
		Email:    "support@opspilot.local",
		Password: "password",
	})
	if err != nil {
		t.Fatalf("expected login success, got error: %v", err)
	}

	if result.AccessToken == "" {
		t.Fatal("expected access token")
	}
	if result.User.Email != "support@opspilot.local" {
		t.Fatalf("expected support user email, got %s", result.User.Email)
	}
}

func TestLoginRejectsInvalidPassword(t *testing.T) {
	service := NewService(fakeRepository{user: User{
		ID:           "user-id",
		Name:         "RPA Support User",
		Email:        "support@opspilot.local",
		PasswordHash: seedPasswordHash,
		Status:       activeUserStatus,
		RoleCode:     "SUPPORT",
	}}, fakeTokenManager{})

	_, err := service.Login(context.Background(), LoginRequest{
		Email:    "support@opspilot.local",
		Password: "wrong-password",
	})
	if !errors.Is(err, ErrInvalidCredentials) {
		t.Fatalf("expected invalid credentials, got %v", err)
	}
}
