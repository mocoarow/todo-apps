package controller

import (
	"net/http"
)

// CookieConfig holds settings for HTTP cookie-based token delivery.
type CookieConfig struct {
	Name                string `yaml:"name" validate:"required"`
	Path                string `yaml:"path" validate:"required"`
	Secure              bool   `yaml:"secure"`
	SameSite            string `yaml:"sameSite" validate:"required,oneof=Lax Strict"`
	RefreshThresholdMin int    `yaml:"refreshThresholdMin" validate:"gte=1"`
}

// SetTokenCookie writes an access-token cookie to the response with the configured attributes.
func (c *CookieConfig) SetTokenCookie(w http.ResponseWriter, token string, tokenTTLMin int) {
	maxAge := tokenTTLMin * 60
	http.SetCookie(w, &http.Cookie{ //nolint:exhaustruct
		Name:     c.Name,
		Value:    token,
		Path:     c.Path,
		MaxAge:   maxAge,
		HttpOnly: true,
		Secure:   c.Secure,
		SameSite: parseSameSite(c.SameSite),
	})
}

func parseSameSite(s string) http.SameSite {
	switch s {
	case "Strict":
		return http.SameSiteStrictMode
	default:
		return http.SameSiteLaxMode
	}
}
