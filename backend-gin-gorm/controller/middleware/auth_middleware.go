package middleware

import (
	"log/slog"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"

	"github.com/mocoarow/todo-apps/backend-gin-gorm/controller"
	"github.com/mocoarow/todo-apps/backend-gin-gorm/controller/telemetry"
	"github.com/mocoarow/todo-apps/backend-gin-gorm/domain"
)

// AuthUsecase defines the use case for extracting user info from a JWT token and refreshing tokens.
type AuthUsecase interface {
	GetUserInfo(input *domain.GetUserInfoInput) (*domain.GetUserInfoOutput, error)
	RefreshToken(input *domain.RefreshTokenInput) (*domain.RefreshTokenOutput, error)
}

// NewAuthMiddleware returns a Gin middleware that validates the Bearer token
// from the Authorization header (with Cookie fallback) and sets the user ID in the Gin context.
// When the token is provided via cookie, sliding refresh is performed automatically.
func NewAuthMiddleware(authUsecase AuthUsecase, cookieConfig *controller.CookieConfig, tokenTTLMin int) gin.HandlerFunc {
	logger := slog.Default().With(slog.String(domain.LoggerNameKey, domain.AppName+"-AuthMiddleware"))

	return func(c *gin.Context) {
		ctx := c.Request.Context()

		ctx, span := tracer.Start(ctx, "AuthMiddleware")
		defer span.End()

		token, fromCookie := extractToken(c, cookieConfig)
		if token == "" {
			logger.InfoContext(ctx, "no token found in Authorization header or cookie")
			c.Status(http.StatusUnauthorized)
			c.Abort()
			return
		}

		input, err := domain.NewGetUserInfoInput(token)
		if err != nil {
			logger.WarnContext(ctx, "new get user info input", slog.Any("error", err))
			c.Status(http.StatusUnauthorized)
			c.Abort()
			return
		}
		output, err := authUsecase.GetUserInfo(input)
		if err != nil {
			logger.WarnContext(ctx, "get user info", slog.Any("error", err))
			c.Status(http.StatusUnauthorized)
			c.Abort()
			return
		}

		c.Set(controller.ContextFieldUserID{}, output.UserInfo.UserID)
		c.Set(controller.ContextFieldLoginID{}, output.UserInfo.LoginID)
		if newCtx, err := telemetry.AddBaggageMembers(ctx, map[string]string{
			"user_id": strconv.Itoa(output.UserInfo.UserID),
		}); err != nil {
			logger.WarnContext(ctx, "add baggage members", slog.Any("error", err))
		} else {
			ctx = newCtx
		}

		c.Request = c.Request.WithContext(ctx)

		if fromCookie && cookieConfig != nil {
			slidingRefresh(c, authUsecase, cookieConfig, tokenTTLMin, output.UserInfo, logger)
		}

		c.Next()
	}
}

func extractToken(c *gin.Context, cookieConfig *controller.CookieConfig) (string, bool) {
	authorization := c.GetHeader("Authorization")
	if strings.HasPrefix(authorization, "Bearer ") {
		return authorization[len("Bearer "):], false
	}

	if cookieConfig != nil {
		cookie, err := c.Cookie(cookieConfig.Name)
		if err == nil && cookie != "" {
			return cookie, true
		}
	}

	return "", false
}

func slidingRefresh(c *gin.Context, authUsecase AuthUsecase, cookieConfig *controller.CookieConfig, tokenTTLMin int, userInfo *domain.UserInfo, logger *slog.Logger) {
	ctx := c.Request.Context()
	refreshInput, err := domain.NewRefreshTokenInput(userInfo.LoginID, userInfo.UserID, userInfo.ExpiresAt)
	if err != nil {
		logger.WarnContext(ctx, "new refresh token input", slog.Any("error", err))
		return
	}

	refreshOutput, err := authUsecase.RefreshToken(refreshInput)
	if err != nil {
		logger.WarnContext(ctx, "refresh token", slog.Any("error", err))
		return
	}

	if refreshOutput.NewAccessToken == "" {
		return
	}

	cookieConfig.SetTokenCookie(c.Writer, refreshOutput.NewAccessToken, tokenTTLMin)
}
