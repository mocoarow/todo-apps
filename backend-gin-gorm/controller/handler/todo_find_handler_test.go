package handler_test

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/mocoarow/todo-apps/backend-gin-gorm/domain"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func Test_TodoHandler_FindTodos_shouldReturn200_whenUsecaseReturnsTodos(t *testing.T) {
	t.Parallel()
	ctx := context.Background()

	// given
	userID := randomUserID()
	todoUsecase := NewMockTodoUsecase(t)
	todoUsecase.EXPECT().FindTodos(mock.Anything, userID).Return([]domain.Todo{
		{
			ID:         userID,
			Text:       "task A",
			IsComplete: false,
		},
	}, nil).Once()
	r := initTodoRouter(t, ctx, todoUsecase, userID)
	w := httptest.NewRecorder()

	// when
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, "/api/v1/todo", nil)
	require.NoError(t, err)
	r.ServeHTTP(w, req)
	respBytes := readBytes(t, w.Body)

	// then
	// - status code
	assert.Equal(t, http.StatusOK, w.Code, "status code should be 200")

	jsonObj := parseJSON(t, respBytes)

	// - todos
	todosExpr := parseExpr(t, "$.todos")
	todos := todosExpr.Get(jsonObj)
	require.Len(t, todos, 1, "response should have one todo")

	// - id
	idExpr := parseExpr(t, "$.todos[0].id")
	id := idExpr.Get(jsonObj)
	require.Len(t, id, 1, "response should have one id")
	assert.Equal(t, int64(userID), id[0])

	// - text
	textExpr := parseExpr(t, "$.todos[0].text")
	text := textExpr.Get(jsonObj)
	require.Len(t, text, 1, "response should have one text")
	assert.Equal(t, "task A", text[0], "text should be 'task A'")

	// - isComplete
	isCompleteExpr := parseExpr(t, "$.todos[0].isComplete")
	isComplete := isCompleteExpr.Get(jsonObj)
	require.Len(t, isComplete, 1, "response should have one isComplete")
	assert.Equal(t, false, isComplete[0], "isComplete should be false")
}

func Test_TodoHandler_FindTodos_shouldReturn500_whenUsecaseReturnsError(t *testing.T) {
	t.Parallel()
	ctx := context.Background()

	// given
	userID := randomUserID()
	todoUsecase := NewMockTodoUsecase(t)
	todoUsecase.EXPECT().FindTodos(mock.Anything, userID).Return(nil, fmt.Errorf("database error")).Once()
	r := initTodoRouter(t, ctx, todoUsecase, userID)
	w := httptest.NewRecorder()

	// when
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, "/api/v1/todo", nil)
	require.NoError(t, err)
	r.ServeHTTP(w, req)
	respBytes := readBytes(t, w.Body)

	// then
	assert.Equal(t, http.StatusInternalServerError, w.Code, "status code should be 500")
	validateErrorResponse(t, respBytes, "internal_server_error", "Internal Server Error")
}

func Test_TodoHandler_FindTodos_shouldReturn500_whenUsecaseReturnsInvalidTodos(t *testing.T) {
	t.Parallel()
	ctx := context.Background()

	// given
	userID := randomUserID()
	todoUsecase := NewMockTodoUsecase(t)
	todoUsecase.EXPECT().FindTodos(mock.Anything, userID).Return([]domain.Todo{{}}, nil).Once()
	r := initTodoRouter(t, ctx, todoUsecase, userID)
	w := httptest.NewRecorder()

	// when
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, "/api/v1/todo", nil)
	require.NoError(t, err)
	r.ServeHTTP(w, req)
	respBytes := readBytes(t, w.Body)
	// then
	assert.Equal(t, http.StatusInternalServerError, w.Code, "status code should be 500")
	validateErrorResponse(t, respBytes, "internal_server_error", "Internal Server Error")
}
