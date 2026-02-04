package handler

import (
	"errors"
	"fmt"
	"log/slog"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/mocoarow/todo-apps/backend-gin-gorm/api"
	"github.com/mocoarow/todo-apps/backend-gin-gorm/domain"
)

// NewUpdateTodoResponse converts a domain Todo to an UpdateTodoResponse API type.
func NewUpdateTodoResponse(todo *domain.Todo) (*api.UpdateTodoResponse, error) {
	if todo == nil {
		return nil, errors.New("todo is nil")
	}
	id, err := safeIntToInt32(todo.ID)
	if err != nil {
		return nil, fmt.Errorf("convert todo ID: %w", err)
	}
	return &api.UpdateTodoResponse{
		ID:         id,
		Text:       todo.Text,
		IsComplete: todo.IsComplete,
		CreatedAt:  todo.CreatedAt,
		UpdatedAt:  todo.UpdatedAt,
	}, nil
}

// UpdateTodo handles PUT /todo/:id and updates an existing todo for the authenticated user.
func (h *TodoHandler) UpdateTodo(c *gin.Context) {
	ctx := c.Request.Context()
	todoID, err := GetIntFromPath(c, "id")
	if err != nil {
		h.logger.WarnContext(ctx, "invalid todo id in path", slog.Any("error", err))
		c.JSON(http.StatusBadRequest, NewErrorResponse("invalid_todo_id", "todo id must be a positive integer"))
		return
	}
	if todoID <= 0 {
		h.logger.WarnContext(ctx, "invalid todo id", slog.Int("todoId", todoID))
		c.JSON(http.StatusBadRequest, NewErrorResponse("invalid_todo_id", "todo id must be a positive integer"))
		return
	}

	userID := c.GetInt(ContextFieldUserID)
	if userID <= 0 {
		h.logger.WarnContext(ctx, "unauthorized: missing or invalid user ID")
		c.JSON(http.StatusUnauthorized, NewErrorResponse("unauthorized", http.StatusText(http.StatusUnauthorized)))
		return
	}
	h.logger.InfoContext(ctx, "UpdateTodo called", slog.Int("userId", userID))

	var req api.UpdateTodoRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		h.logger.WarnContext(ctx, "invalid update todo request", slog.Any("error", err))
		c.JSON(http.StatusBadRequest, NewErrorResponse("invalid_request", "request body is invalid"))
		return
	}

	input, err := domain.NewUpdateTodoInput(todoID, userID, req.Text, req.IsComplete)
	if err != nil {
		h.logger.WarnContext(ctx, "invalid update todo input", slog.Any("error", err))
		c.JSON(http.StatusBadRequest, NewErrorResponse("invalid_request", "request body is invalid"))
		return
	}

	output, err := h.usecase.UpdateTodo(ctx, input)
	if errors.Is(err, domain.ErrTodoNotFound) {
		h.logger.WarnContext(ctx, "todo not found", slog.Int("todoId", todoID))
		c.JSON(http.StatusNotFound, NewErrorResponse("todo_not_found", http.StatusText(http.StatusNotFound)))
		return
	}
	if err != nil {
		h.logger.ErrorContext(ctx, "failed to update todo", slog.Any("error", err))
		c.JSON(http.StatusInternalServerError, NewErrorResponse("internal_server_error", http.StatusText(http.StatusInternalServerError)))
		return
	}

	resp, err := NewUpdateTodoResponse(output.Todo)
	if err != nil {
		h.logger.ErrorContext(ctx, "failed to create response", slog.Any("error", err))
		c.JSON(http.StatusInternalServerError, NewErrorResponse("internal_server_error", http.StatusText(http.StatusInternalServerError)))
		return
	}
	c.JSON(http.StatusOK, resp)
}
