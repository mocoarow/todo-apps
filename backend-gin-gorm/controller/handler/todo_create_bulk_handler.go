package handler

import (
	"log/slog"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/mocoarow/todo-apps/backend-gin-gorm/api"
	"github.com/mocoarow/todo-apps/backend-gin-gorm/domain"
)

// NewCreateBulkTodosResponse converts a slice of domain Todos to a CreateBulkTodosResponse API type.
func NewCreateBulkTodosResponse(todos []*domain.Todo) (*api.CreateBulkTodosResponse, error) {
	resp := &api.CreateBulkTodosResponse{
		Todos: make([]api.CreateTodoResponse, 0, len(todos)),
	}
	for _, todo := range todos {
		todoResp, err := NewCreateTodoResponse(todo)
		if err != nil {
			return nil, err
		}
		resp.Todos = append(resp.Todos, *todoResp)
	}
	return resp, nil
}

// CreateBulkTodos handles POST /todo/bulk and creates multiple todos at once for the authenticated user.
func (h *TodoHandler) CreateBulkTodos(c *gin.Context) {
	ctx := c.Request.Context()
	userID := c.GetInt(ContextFieldUserID)
	if userID <= 0 {
		h.logger.WarnContext(ctx, "unauthorized: missing or invalid user ID")
		c.JSON(http.StatusUnauthorized, NewErrorResponse("unauthorized", http.StatusText(http.StatusUnauthorized)))
		return
	}
	h.logger.InfoContext(ctx, "CreateBulkTodos called", slog.Int("userId", userID))

	var req api.CreateBulkTodosRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		h.logger.WarnContext(ctx, "invalid create bulk todos request", slog.Any("error", err))
		c.JSON(http.StatusBadRequest, NewErrorResponse("invalid_request", "request body is invalid"))
		return
	}

	todoInputs := make([]*domain.CreateTodoInput, len(req.Todos))
	for i, reqTodo := range req.Todos {
		input, err := domain.NewCreateTodoInput(userID, reqTodo.Text)
		if err != nil {
			h.logger.ErrorContext(ctx, "invalid create todo input", slog.Any("error", err), slog.Int("index", i))
			c.JSON(http.StatusBadRequest, NewErrorResponse("invalid_request", "request body is invalid"))
			return
		}
		todoInputs[i] = input
	}

	bulkInput, err := domain.NewCreateBulkTodosInput(userID, todoInputs)
	if err != nil {
		h.logger.ErrorContext(ctx, "invalid create bulk todos input", slog.Any("error", err))
		c.JSON(http.StatusBadRequest, NewErrorResponse("invalid_request", "request body is invalid"))
		return
	}

	output, err := h.usecase.CreateBulkTodos(ctx, bulkInput)
	if err != nil {
		h.logger.ErrorContext(ctx, "failed to create bulk todos", slog.Any("error", err))
		c.JSON(http.StatusInternalServerError, NewErrorResponse("internal_server_error", http.StatusText(http.StatusInternalServerError)))
		return
	}

	resp, err := NewCreateBulkTodosResponse(output.Todos)
	if err != nil {
		h.logger.ErrorContext(ctx, "failed to create response", slog.Any("error", err))
		c.JSON(http.StatusInternalServerError, NewErrorResponse("internal_server_error", http.StatusText(http.StatusInternalServerError)))
		return
	}

	c.JSON(http.StatusCreated, resp)
}
