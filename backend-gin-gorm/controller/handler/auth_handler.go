package handler

import (
	"errors"
	"log/slog"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/mocoarow/todo-apps/backend-gin-gorm/api"
	"github.com/mocoarow/todo-apps/backend-gin-gorm/controller"
	"github.com/mocoarow/todo-apps/backend-gin-gorm/domain"
)

// AuthUsecase defines the authentication use case required by the handler.
type AuthUsecase interface {
	Authenticate(input *domain.AuthenticateInput) (*domain.AuthenticateOutput, error)
}

// AuthHandler handles HTTP requests for user authentication.
type AuthHandler struct {
	usecase      AuthUsecase
	logger       *slog.Logger
	cookieConfig *controller.CookieConfig
	tokenTTLMin  int
}

// NewAuthHandler returns a new AuthHandler with the given use case.
func NewAuthHandler(usecase AuthUsecase, cookieConfig *controller.CookieConfig, tokenTTLMin int) *AuthHandler {
	return &AuthHandler{
		usecase:      usecase,
		logger:       slog.Default().With(slog.String(domain.LoggerNameKey, "AuthHandler")),
		cookieConfig: cookieConfig,
		tokenTTLMin:  tokenTTLMin,
	}
}

// Authenticate handles POST /auth/authenticate and returns a JWT access token.
func (h *AuthHandler) Authenticate(c *gin.Context) {
	ctx := c.Request.Context()
	var req api.AuthenticateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		h.logger.WarnContext(ctx, "invalid authenticate request", slog.Any("error", err))
		c.JSON(http.StatusBadRequest, NewErrorResponse("invalid_authenticate_request", "request body is invalid"))
		return
	}

	tokenDelivery := c.GetHeader("X-Token-Delivery")
	switch tokenDelivery {
	case "", "json", "cookie":
		// valid
	default:
		c.JSON(http.StatusBadRequest, NewErrorResponse("invalid_token_delivery", "X-Token-Delivery must be 'json' or 'cookie'"))
		return
	}

	input, err := domain.NewAuthenticateInput(req.LoginID, req.Password)
	if err != nil {
		h.logger.ErrorContext(ctx, "invalid authenticate input", slog.Any("error", err))
		c.JSON(http.StatusInternalServerError, NewErrorResponse("internal_server_error", http.StatusText(http.StatusInternalServerError)))
		return
	}

	output, err := h.usecase.Authenticate(input)
	if err != nil {
		if errors.Is(err, domain.ErrUnauthenticated) {
			h.logger.WarnContext(ctx, "unauthenticated", slog.Any("error", err))
			c.JSON(http.StatusUnauthorized, NewErrorResponse("unauthenticated", http.StatusText(http.StatusUnauthorized)))
			return
		}
		h.logger.ErrorContext(ctx, "authenticate", slog.Any("error", err))
		c.JSON(http.StatusInternalServerError, NewErrorResponse("internal_server_error", http.StatusText(http.StatusInternalServerError)))
		return
	}

	switch tokenDelivery {
	case "cookie":
		if h.cookieConfig == nil {
			h.logger.ErrorContext(ctx, "cookie delivery requested but cookie config is not available")
			c.JSON(http.StatusInternalServerError, NewErrorResponse("cookie_not_configured", "cookie delivery is not configured"))
			return
		}
		h.cookieConfig.SetTokenCookie(c.Writer, output.AccessToken, h.tokenTTLMin)
		c.JSON(http.StatusOK, api.AuthenticateResponse{
			AccessToken: nil,
		})
	default:
		resp := api.AuthenticateResponse{
			AccessToken: &output.AccessToken,
		}
		c.JSON(http.StatusOK, resp)
	}
}

// NewInitAuthRouterFunc returns an InitRouterGroupFunc that registers auth routes under an "auth" group.
func NewInitAuthRouterFunc(authUsecase AuthUsecase, cookieConfig *controller.CookieConfig, tokenTTLMin int) InitRouterGroupFunc {
	return func(parentRouterGroup gin.IRouter, middleware ...gin.HandlerFunc) {
		auth := parentRouterGroup.Group("auth", middleware...)
		authHandler := NewAuthHandler(authUsecase, cookieConfig, tokenTTLMin)

		auth.POST("/authenticate", authHandler.Authenticate)
	}
}
