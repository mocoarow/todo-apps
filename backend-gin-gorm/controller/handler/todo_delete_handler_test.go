package handler_test

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/mocoarow/todo-apps/backend-gin-gorm/domain"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func Test_TodoHandler_DeleteTodo_shouldReturn400_whenInvalidPath(t *testing.T) {
	t.Parallel()
	ctx := context.Background()

	// given
	userID := randomUserID()

	tests := []struct {
		name            string
		url             string
		expectedMessage string
	}{
		{
			name:            "non-integer id",
			url:             "/api/v1/todo/invalid",
			expectedMessage: "todo id must be a positive integer",
		},
		{
			name:            "missing id",
			url:             "/api/v1/todo/0",
			expectedMessage: "todo id must be a positive integer",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			todoUsecase := NewMockTodoUsecase(t)
			r := initTodoRouter(t, ctx, todoUsecase, userID)
			w := httptest.NewRecorder()

			// when
			req, err := http.NewRequestWithContext(ctx, http.MethodDelete, tt.url, nil)
			require.NoError(t, err)
			r.ServeHTTP(w, req)
			respBytes := readBytes(t, w.Body)

			// then
			assert.Equal(t, http.StatusBadRequest, w.Code, "status code should be 400")
			validateErrorResponse(t, respBytes, "invalid_todo_id", tt.expectedMessage)
		})
	}
}

func Test_TodoHandler_DeleteTodo_shouldReturn404_whenTodoNotFound(t *testing.T) {
	t.Parallel()
	ctx := context.Background()

	// given
	userID := randomUserID()
	todoUsecase := NewMockTodoUsecase(t)
	todoUsecase.EXPECT().DeleteTodo(mock.Anything, &domain.DeleteTodoInput{
		UserID: userID,
		ID:     999999,
	}).Return(domain.ErrTodoNotFound).Once()
	r := initTodoRouter(t, ctx, todoUsecase, userID)
	w := httptest.NewRecorder()

	// when
	req, err := http.NewRequestWithContext(ctx, http.MethodDelete, "/api/v1/todo/999999", nil)
	require.NoError(t, err)
	r.ServeHTTP(w, req)
	respBytes := readBytes(t, w.Body)

	// then
	assert.Equal(t, http.StatusNotFound, w.Code, "status code should be 404")
	validateErrorResponse(t, respBytes, "todo_not_found", "Not Found")
}

func Test_TodoHandler_DeleteTodo_shouldReturn204_whenValidRequest(t *testing.T) {
	t.Parallel()
	ctx := context.Background()

	// given
	userID := randomUserID()
	todoUsecase := NewMockTodoUsecase(t)
	todoUsecase.EXPECT().DeleteTodo(mock.Anything, &domain.DeleteTodoInput{
		UserID: userID,
		ID:     1,
	}).Return(nil).Once()
	r := initTodoRouter(t, ctx, todoUsecase, userID)
	w := httptest.NewRecorder()

	// when
	req, err := http.NewRequestWithContext(ctx, http.MethodDelete, "/api/v1/todo/1", nil)
	require.NoError(t, err)
	r.ServeHTTP(w, req)
	readBytes(t, w.Body)

	// then
	assert.Equal(t, http.StatusNoContent, w.Code, "status code should be 204")
	assert.Empty(t, w.Body.String(), "response body should be empty")
}

func Test_TodoHandler_DeleteTodo_shouldReturn500_whenUsecaseReturnsError(t *testing.T) {
	t.Parallel()
	ctx := context.Background()

	// given
	userID := randomUserID()
	todoUsecase := NewMockTodoUsecase(t)
	todoUsecase.EXPECT().DeleteTodo(mock.Anything, &domain.DeleteTodoInput{
		UserID: userID,
		ID:     1,
	}).Return(assert.AnError).Once()
	r := initTodoRouter(t, ctx, todoUsecase, userID)
	w := httptest.NewRecorder()

	// when
	req, err := http.NewRequestWithContext(ctx, http.MethodDelete, "/api/v1/todo/1", nil)
	require.NoError(t, err)
	r.ServeHTTP(w, req)
	respBytes := readBytes(t, w.Body)

	// then
	assert.Equal(t, http.StatusInternalServerError, w.Code, "status code should be 500")
	validateErrorResponse(t, respBytes, "internal_server_error", "Internal Server Error")
}
