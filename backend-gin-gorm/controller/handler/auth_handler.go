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

// GetMe handles GET /auth/me and returns the authenticated user's ID and login ID.
func (h *AuthHandler) GetMe(c *gin.Context) {
	ctx := c.Request.Context()
	userID := c.GetInt(controller.ContextFieldUserID{})
	if userID <= 0 {
		h.logger.WarnContext(ctx, "unauthorized: missing or invalid user ID")
		c.JSON(http.StatusUnauthorized, NewErrorResponse("unauthorized", http.StatusText(http.StatusUnauthorized)))
		return
	}

	loginID := c.GetString(controller.ContextFieldLoginID{})
	if loginID == "" {
		h.logger.WarnContext(ctx, "unauthorized: missing login ID")
		c.JSON(http.StatusUnauthorized, NewErrorResponse("unauthorized", http.StatusText(http.StatusUnauthorized)))
		return
	}

	userIDInt32, err := safeIntToInt32(userID)
	if err != nil {
		h.logger.ErrorContext(ctx, "convert user ID", slog.Any("error", err))
		c.JSON(http.StatusInternalServerError, NewErrorResponse("internal_server_error", http.StatusText(http.StatusInternalServerError)))
		return
	}

	c.JSON(http.StatusOK, api.GetMeResponse{
		UserID:  userIDInt32,
		LoginID: loginID,
	})
}

// Logout handles POST /auth/logout and clears the access-token cookie.
func (h *AuthHandler) Logout(c *gin.Context) {
	ctx := c.Request.Context()
	if h.cookieConfig == nil {
		h.logger.ErrorContext(ctx, "logout requested but cookie config is not available")
		c.JSON(http.StatusInternalServerError, NewErrorResponse("cookie_not_configured", "cookie delivery is not configured"))
		return
	}
	h.cookieConfig.ClearTokenCookie(c.Writer)
	c.Status(http.StatusNoContent)
}

// NewInitAuthRouterFunc returns an InitRouterGroupFunc that registers auth routes under an "auth" group.
// authMiddleware is applied only to routes that require authentication (e.g. /me).
func NewInitAuthRouterFunc(authUsecase AuthUsecase, cookieConfig *controller.CookieConfig, tokenTTLMin int, authMiddleware gin.HandlerFunc) InitRouterGroupFunc {
	return func(parentRouterGroup gin.IRouter, middleware ...gin.HandlerFunc) {
		auth := parentRouterGroup.Group("auth", middleware...)
		authHandler := NewAuthHandler(authUsecase, cookieConfig, tokenTTLMin)

		auth.POST("/authenticate", authHandler.Authenticate)
		auth.POST("/logout", authHandler.Logout)
		auth.GET("/me", authMiddleware, authHandler.GetMe)
	}
}
