package jwt

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

const claimsContextKey = "auth_claims"

var (
	ErrMissingSecret = errors.New("jwt secret is required")
	ErrInvalidToken  = errors.New("invalid token")
	ErrExpiredToken  = errors.New("expired token")
)

type Manager struct {
	secret string
	now    func() time.Time
}

type Claims struct {
	Subject   string `json:"sub"`
	Name      string `json:"name"`
	Email     string `json:"email"`
	Role      string `json:"role"`
	IssuedAt  int64  `json:"iat"`
	ExpiresAt int64  `json:"exp"`
}

func NewManager(secret string) *Manager {
	return &Manager{
		secret: secret,
		now:    time.Now,
	}
}

func (m *Manager) Generate(subject string, name string, email string, role string, ttl time.Duration) (string, int64, error) {
	if m.secret == "" {
		return "", 0, ErrMissingSecret
	}

	now := m.now()
	expiresAt := now.Add(ttl)
	claims := Claims{
		Subject:   subject,
		Name:      name,
		Email:     email,
		Role:      role,
		IssuedAt:  now.Unix(),
		ExpiresAt: expiresAt.Unix(),
	}

	header := map[string]string{
		"alg": "HS256",
		"typ": "JWT",
	}

	headerPart, err := encodeJSON(header)
	if err != nil {
		return "", 0, fmt.Errorf("encode jwt header: %w", err)
	}

	claimsPart, err := encodeJSON(claims)
	if err != nil {
		return "", 0, fmt.Errorf("encode jwt claims: %w", err)
	}

	unsignedToken := headerPart + "." + claimsPart
	signature := sign(unsignedToken, m.secret)

	return unsignedToken + "." + signature, int64(ttl.Seconds()), nil
}

func (m *Manager) Verify(token string) (Claims, error) {
	if m.secret == "" {
		return Claims{}, ErrMissingSecret
	}

	parts := strings.Split(token, ".")
	if len(parts) != 3 {
		return Claims{}, ErrInvalidToken
	}

	unsignedToken := parts[0] + "." + parts[1]
	expectedSignature := sign(unsignedToken, m.secret)
	if !hmac.Equal([]byte(expectedSignature), []byte(parts[2])) {
		return Claims{}, ErrInvalidToken
	}

	var claims Claims
	payload, err := base64.RawURLEncoding.DecodeString(parts[1])
	if err != nil {
		return Claims{}, ErrInvalidToken
	}
	if err := json.Unmarshal(payload, &claims); err != nil {
		return Claims{}, ErrInvalidToken
	}

	if claims.Subject == "" || claims.ExpiresAt == 0 {
		return Claims{}, ErrInvalidToken
	}
	if m.now().Unix() > claims.ExpiresAt {
		return Claims{}, ErrExpiredToken
	}

	return claims, nil
}

func SetClaims(ctx *gin.Context, claims Claims) {
	ctx.Set(claimsContextKey, claims)
}

func ClaimsFromContext(ctx *gin.Context) (Claims, bool) {
	value, exists := ctx.Get(claimsContextKey)
	if !exists {
		return Claims{}, false
	}

	claims, ok := value.(Claims)
	return claims, ok
}

func encodeJSON(value interface{}) (string, error) {
	encoded, err := json.Marshal(value)
	if err != nil {
		return "", err
	}

	return base64.RawURLEncoding.EncodeToString(encoded), nil
}

func sign(unsignedToken string, secret string) string {
	mac := hmac.New(sha256.New, []byte(secret))
	mac.Write([]byte(unsignedToken))
	return base64.RawURLEncoding.EncodeToString(mac.Sum(nil))
}
