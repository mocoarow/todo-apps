package handler

import (
	"errors"
	"fmt"
	"log/slog"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/mocoarow/todo-apps/backend-gin-gorm/api"
	"github.com/mocoarow/todo-apps/backend-gin-gorm/controller"
	"github.com/mocoarow/todo-apps/backend-gin-gorm/domain"
)

// NewFindTodoResponseTodo converts a domain Todo to a FindTodoResponseTodo API type.
func NewFindTodoResponseTodo(todo *domain.Todo) (*api.FindTodoResponseTodo, error) {
	if todo == nil {
		return nil, errors.New("todo is nil")
	}
	id, err := safeIntToInt32(todo.ID)
	if err != nil {
		return nil, fmt.Errorf("convert todo ID: %w", err)
	}
	return &api.FindTodoResponseTodo{
		ID:         id,
		Text:       todo.Text,
		IsComplete: todo.IsComplete,
		CreatedAt:  todo.CreatedAt,
		UpdatedAt:  todo.UpdatedAt,
	}, nil
}

// NewFindTodoResponse converts a slice of domain Todos to a FindTodoResponse API type.
func NewFindTodoResponse(todos []*domain.Todo) (*api.FindTodoResponse, error) {
	resp := &api.FindTodoResponse{
		Todos: make([]api.FindTodoResponseTodo, 0, len(todos)),
	}
	for _, todo := range todos {
		todoResp, err := NewFindTodoResponseTodo(todo)
		if err != nil {
			return nil, fmt.Errorf("convert todo: %w", err)
		}
		resp.Todos = append(resp.Todos, *todoResp)
	}
	return resp, nil
}

// FindTodos handles GET /todo and returns all todos for the authenticated user.
func (h *TodoHandler) FindTodos(c *gin.Context) {
	ctx := c.Request.Context()
	userID := c.GetInt(controller.ContextFieldUserID{})
	if userID <= 0 {
		h.logger.WarnContext(ctx, "unauthorized: missing or invalid user ID")
		c.JSON(http.StatusUnauthorized, NewErrorResponse("unauthorized", http.StatusText(http.StatusUnauthorized)))
		return
	}

	h.logger.InfoContext(ctx, "FindTodos called", slog.Int("userId", userID))

	todos, err := h.usecase.FindTodos(ctx, userID)
	if err != nil {
		h.logger.ErrorContext(ctx, "failed to find todos", slog.Any("error", err))
		c.JSON(http.StatusInternalServerError, NewErrorResponse("internal_server_error", http.StatusText(http.StatusInternalServerError)))
		return
	}

	resp, err := NewFindTodoResponse(todos)
	if err != nil {
		h.logger.ErrorContext(ctx, "failed to create find todo response", slog.Any("error", err))
		c.JSON(http.StatusInternalServerError, NewErrorResponse("internal_server_error", http.StatusText(http.StatusInternalServerError)))
		return
	}

	c.JSON(http.StatusOK, resp)
}
