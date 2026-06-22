package jwt

import (
	"testing"
	"time"
)

func TestGenerateAndVerify(t *testing.T) {
	manager := NewManager("test-secret")

	token, expiresIn, err := manager.Generate(
		"user-id",
		"RPA Support User",
		"support@opspilot.local",
		"SUPPORT",
		time.Hour,
	)
	if err != nil {
		t.Fatalf("expected token generation success, got error: %v", err)
	}
	if token == "" {
		t.Fatal("expected token")
	}
	if expiresIn != 3600 {
		t.Fatalf("expected expires_in 3600, got %d", expiresIn)
	}

	claims, err := manager.Verify(token)
	if err != nil {
		t.Fatalf("expected token verification success, got error: %v", err)
	}
	if claims.Subject != "user-id" {
		t.Fatalf("expected subject user-id, got %s", claims.Subject)
	}
	if claims.Role != "SUPPORT" {
		t.Fatalf("expected role SUPPORT, got %s", claims.Role)
	}
}

func TestVerifyRejectsTamperedToken(t *testing.T) {
	manager := NewManager("test-secret")

	token, _, err := manager.Generate(
		"user-id",
		"RPA Support User",
		"support@opspilot.local",
		"SUPPORT",
		time.Hour,
	)
	if err != nil {
		t.Fatalf("expected token generation success, got error: %v", err)
	}

	_, err = manager.Verify(token + "tampered")
	if err != ErrInvalidToken {
		t.Fatalf("expected invalid token, got %v", err)
	}
}
