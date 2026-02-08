package gateway

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"

	"github.com/mocoarow/todo-apps/backend-gin-gorm/domain"
)

type userClaims struct {
	LoginID string `json:"loginId"`
	UserID  int    `json:"userId"`
	jwt.RegisteredClaims
}

// AuthTokenManager implements JWT token creation and parsing using HMAC signing.
type AuthTokenManager struct {
	signingKey       []byte
	signingMethod    jwt.SigningMethod
	tokenTimeout     time.Duration
	refreshThreshold time.Duration
}

// NewAuthTokenManager returns a new AuthTokenManager with the given signing parameters.
func NewAuthTokenManager(signingKey []byte, signingMethod jwt.SigningMethod, tokenTimeout time.Duration, refreshThreshold time.Duration) *AuthTokenManager {
	return &AuthTokenManager{
		signingKey:       signingKey,
		signingMethod:    signingMethod,
		tokenTimeout:     tokenTimeout,
		refreshThreshold: refreshThreshold,
	}
}

// CreateToken generates a signed JWT for the given user.
func (m *AuthTokenManager) CreateToken(loginID string, userID int) (string, error) {
	accessToken, err := m.createJWT(loginID, userID, m.tokenTimeout)
	if err != nil {
		return "", fmt.Errorf("create token: %w", err)
	}

	return accessToken, nil
}

// ParseToken validates a JWT string and returns the embedded user info including token expiry.
func (m *AuthTokenManager) ParseToken(tokenString string) (*domain.UserInfo, error) {
	claims, err := m.parseToken(tokenString)
	if err != nil {
		return nil, fmt.Errorf("parse token: %w", err)
	}

	userInfo, err := domain.NewUserInfo(claims.UserID, claims.LoginID, claims.ExpiresAt.Time)
	if err != nil {
		return nil, fmt.Errorf("create user info: %w", err)
	}

	return userInfo, nil
}

// RefreshToken checks if the token's remaining lifetime is below the refresh threshold.
// If so, it issues a new token with a fresh expiry. Returns empty string if no refresh is needed.
func (m *AuthTokenManager) RefreshToken(loginID string, userID int, expiresAt time.Time) (string, error) {
	remaining := time.Until(expiresAt)
	if remaining > m.refreshThreshold {
		return "", nil
	}

	newToken, err := m.createJWT(loginID, userID, m.tokenTimeout)
	if err != nil {
		return "", fmt.Errorf("create refreshed token: %w", err)
	}

	return newToken, nil
}

func (m *AuthTokenManager) createJWT(loginID string, userID int, duration time.Duration) (string, error) {
	now := time.Now()
	claims := userClaims{
		LoginID: loginID,
		UserID:  userID,
		RegisteredClaims: jwt.RegisteredClaims{ //nolint:exhaustruct
			Issuer:    "backend-gin-gorm",
			Subject:   "AccessToken",
			Audience:  []string{"backend-gin-gorm"},
			NotBefore: jwt.NewNumericDate(now),
			IssuedAt:  jwt.NewNumericDate(now),
			ExpiresAt: jwt.NewNumericDate(now.Add(duration)),
		},
	}
	token := jwt.NewWithClaims(m.signingMethod, claims)
	signed, err := token.SignedString(m.signingKey)
	if err != nil {
		return "", fmt.Errorf("sign token: %w", err)
	}

	return signed, nil
}

func (m *AuthTokenManager) parseToken(tokenString string) (*userClaims, error) {
	keyFunc := func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return m.signingKey, nil
	}

	currentToken, err := jwt.ParseWithClaims(tokenString, &userClaims{}, keyFunc) //nolint:exhaustruct
	if err != nil {
		return nil, fmt.Errorf("parse token: %w", err)
	}
	if !currentToken.Valid {
		return nil, fmt.Errorf("invalid token")
	}

	currentClaims, ok := currentToken.Claims.(*userClaims)
	if !ok {
		return nil, fmt.Errorf("invalid claims")
	}

	return currentClaims, nil
}
