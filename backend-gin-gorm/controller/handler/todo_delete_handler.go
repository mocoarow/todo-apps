package handler

import (
	"errors"
	"log/slog"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/mocoarow/todo-apps/backend-gin-gorm/domain"
)

// DeleteTodo handles DELETE /todo/:id and removes a todo for the authenticated user.
func (h *TodoHandler) DeleteTodo(c *gin.Context) {
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
	h.logger.InfoContext(ctx, "DeleteTodo called", slog.Int("userId", userID), slog.Int("todoId", todoID))

	input, err := domain.NewDeleteTodoInput(todoID, userID)
	if err != nil {
		h.logger.WarnContext(ctx, "invalid delete todo input", slog.Any("error", err))
		c.JSON(http.StatusBadRequest, NewErrorResponse("invalid_request", "request is invalid"))
		return
	}

	err = h.usecase.DeleteTodo(ctx, input)
	if errors.Is(err, domain.ErrTodoNotFound) {
		h.logger.WarnContext(ctx, "todo not found", slog.Int("todoId", todoID))
		c.JSON(http.StatusNotFound, NewErrorResponse("todo_not_found", http.StatusText(http.StatusNotFound)))
		return
	}
	if err != nil {
		h.logger.ErrorContext(ctx, "failed to delete todo", slog.Any("error", err))
		c.JSON(http.StatusInternalServerError, NewErrorResponse("internal_server_error", http.StatusText(http.StatusInternalServerError)))
		return
	}

	c.Status(http.StatusNoContent)
}
