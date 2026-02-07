package handler_test

import (
	"context"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/mocoarow/todo-apps/backend-gin-gorm/controller"
	"github.com/mocoarow/todo-apps/backend-gin-gorm/controller/handler"
	"github.com/mocoarow/todo-apps/backend-gin-gorm/domain"
)

func mockAuthMiddleware(userID int) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set(controller.ContextFieldUserID{}, userID)
		c.Next()
	}
}

func initTodoRouter(t *testing.T, ctx context.Context, todoUsecase handler.TodoUsecase, userID int) *gin.Engine {
	t.Helper()

	router := handler.InitRootRouterGroup(ctx, config, domain.AppName)
	api := router.Group("api")
	v1 := api.Group("v1")

	v1.Use(mockAuthMiddleware(userID))

	initTodoRouterFunc := handler.NewInitTodoRouterFunc(todoUsecase)
	initTodoRouterFunc(v1)

	return router
}
