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

type AuthUsecase interface {
	GetUserInfo(input *domain.GetUserInfoInput) (*domain.GetUserInfoOutput, error)
}

func NewAuthMiddleware(authUsecase AuthUsecase) gin.HandlerFunc {
	logger := slog.Default().With(slog.String(domain.LoggerNameKey, domain.AppName+"-AuthMiddleware"))

	return func(c *gin.Context) {
		ctx := c.Request.Context()

		ctx, span := tracer.Start(ctx, "AuthMiddleware")
		defer span.End()

		authorization := c.GetHeader("Authorization")
		if !strings.HasPrefix(authorization, "Bearer ") {
			logger.InfoContext(ctx, "invalid header. Bearer not found")
			c.Status(http.StatusUnauthorized)
			c.Abort()
			return
		}

		bearerToken := authorization[len("Bearer "):]
		input, err := domain.NewGetUserInfoInput(bearerToken)
		if err != nil {
			logger.WarnContext(ctx, "new get user info input", slog.String("error", err.Error()))
			c.Status(http.StatusUnauthorized)
			c.Abort()
			return
		}
		output, err := authUsecase.GetUserInfo(input)
		if err != nil {
			logger.WarnContext(ctx, "get user info", slog.String("error", err.Error()))
			c.Status(http.StatusUnauthorized)
			c.Abort()
			return
		}

		c.Set(controller.ContextFieldUserID{}, output.UserInfo.UserID)
		if newCtx, err := telemetry.AddBaggageMembers(ctx, map[string]string{
			"user_id": strconv.Itoa(output.UserInfo.UserID),
		}); err != nil {
			logger.WarnContext(ctx, "add baggage members", slog.String("error", err.Error()))
		} else {
			ctx = newCtx
		}

		c.Request = c.Request.WithContext(ctx)

		c.Next()
	}
}
