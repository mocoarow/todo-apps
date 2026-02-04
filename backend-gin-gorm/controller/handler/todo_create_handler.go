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

// NewCreateTodoResponse converts a domain Todo to a CreateTodoResponse API type.
func NewCreateTodoResponse(todo *domain.Todo) (*api.CreateTodoResponse, error) {
	if todo == nil {
		return nil, errors.New("todo is nil")
	}
	id, err := safeIntToInt32(todo.ID)
	if err != nil {
		return nil, fmt.Errorf("convert todo ID: %w", err)
	}
	return &api.CreateTodoResponse{
		ID:         id,
		Text:       todo.Text,
		IsComplete: todo.IsComplete,
		CreatedAt:  todo.CreatedAt,
		UpdatedAt:  todo.UpdatedAt,
	}, nil
}

// CreateTodo handles POST /todo and creates a single todo for the authenticated user.
func (h *TodoHandler) CreateTodo(c *gin.Context) {
	ctx := c.Request.Context()
	userID := c.GetInt(ContextFieldUserID)
	if userID <= 0 {
		h.logger.WarnContext(ctx, "unauthorized: missing or invalid user ID")
		c.JSON(http.StatusUnauthorized, NewErrorResponse("unauthorized", http.StatusText(http.StatusUnauthorized)))
		return
	}
	h.logger.InfoContext(ctx, "CreateTodo called", slog.Int("userId", userID))

	var req api.CreateTodoRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		h.logger.WarnContext(ctx, "invalid create todo request", slog.Any("error", err))
		c.JSON(http.StatusBadRequest, NewErrorResponse("invalid_request", "request body is invalid"))
		return
	}

	input, err := domain.NewCreateTodoInput(userID, req.Text)
	if err != nil {
		h.logger.WarnContext(ctx, "invalid create todo input", slog.Any("error", err))
		c.JSON(http.StatusBadRequest, NewErrorResponse("invalid_request", "request body is invalid"))
		return
	}

	output, err := h.usecase.CreateTodo(ctx, input)
	if err != nil {
		h.logger.ErrorContext(ctx, "failed to create todo", slog.Any("error", err))
		c.JSON(http.StatusInternalServerError, NewErrorResponse("internal_server_error", http.StatusText(http.StatusInternalServerError)))
		return
	}

	resp, err := NewCreateTodoResponse(output.Todo)
	if err != nil {
		h.logger.ErrorContext(ctx, "failed to create response", slog.Any("error", err))
		c.JSON(http.StatusInternalServerError, NewErrorResponse("internal_server_error", http.StatusText(http.StatusInternalServerError)))
		return
	}
	c.JSON(http.StatusCreated, resp)
}
