package handler

import (
	"context"
	"log/slog"

	"github.com/gin-gonic/gin"

	"github.com/mocoarow/todo-apps/backend-gin-gorm/domain"
)

// TodoUsecase defines the use case operations for managing todos.
type TodoUsecase interface {
	FindTodos(ctx context.Context, userID int) ([]domain.Todo, error)
	CreateTodo(ctx context.Context, input *domain.CreateTodoInput) (*domain.CreateTodoOutput, error)
	CreateBulkTodos(ctx context.Context, input *domain.CreateBulkTodosInput) (*domain.CreateBulkTodosOutput, error)
	UpdateTodo(ctx context.Context, input *domain.UpdateTodoInput) (*domain.UpdateTodoOutput, error)
	DeleteTodo(ctx context.Context, input *domain.DeleteTodoInput) error
}

// TodoHandler handles HTTP requests for todo CRUD operations.
type TodoHandler struct {
	usecase TodoUsecase
	logger  *slog.Logger
}

// NewTodoHandler creates a new TodoHandler with the given use case.
func NewTodoHandler(usecase TodoUsecase) *TodoHandler {
	return &TodoHandler{
		usecase: usecase,
		logger:  slog.Default().With(slog.String(domain.LoggerNameKey, "TodoHandler")),
	}
}

// NewInitTodoRouterFunc returns an InitRouterGroupFunc that registers todo routes under a "todo" group.
func NewInitTodoRouterFunc(todoUsecase TodoUsecase) InitRouterGroupFunc {
	return func(parentRouterGroup gin.IRouter, middleware ...gin.HandlerFunc) {
		todo := parentRouterGroup.Group("todo", middleware...)
		todoHandler := NewTodoHandler(todoUsecase)

		todo.POST("", todoHandler.CreateTodo)
		todo.POST("/bulk", todoHandler.CreateBulkTodos)
		todo.GET("", todoHandler.FindTodos)
		todo.PUT("/:id", todoHandler.UpdateTodo)
		todo.DELETE("/:id", todoHandler.DeleteTodo)
	}
}
