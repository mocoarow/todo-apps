package handler

import (
	"errors"
	"log/slog"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/mocoarow/todo-apps/backend-gin-gorm/api"
	"github.com/mocoarow/todo-apps/backend-gin-gorm/domain"
)

type AuthUsecase interface {
	Authenticate(input *domain.AuthenticateInput) (*domain.AuthenticateOutput, error)
}

type AuthHandler struct {
	usecase AuthUsecase
	logger  *slog.Logger
}

func NewAuthHandler(usecase AuthUsecase) *AuthHandler {
	return &AuthHandler{
		usecase: usecase,
		logger:  slog.Default().With(slog.String(domain.LoggerNameKey, "AuthHandler")),
	}
}

func (h *AuthHandler) Authenticate(c *gin.Context) {
	ctx := c.Request.Context()
	var req api.AuthenticateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		h.logger.WarnContext(ctx, "invalid authenticate request", slog.Any("error", err))
		c.JSON(http.StatusBadRequest, NewErrorResponse("invalid_authenticate_request", "request body is invalid"))
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

	resp := api.AuthenticateResponse{
		AccessToken: output.AccessToken,
	}
	c.JSON(http.StatusOK, resp)
}

func NewInitAuthRouterFunc(authUsecase AuthUsecase) InitRouterGroupFunc {
	return func(parentRouterGroup gin.IRouter, middleware ...gin.HandlerFunc) {
		auth := parentRouterGroup.Group("auth", middleware...)
		authHandler := NewAuthHandler(authUsecase)

		auth.POST("/authenticate", authHandler.Authenticate)
	}
}
