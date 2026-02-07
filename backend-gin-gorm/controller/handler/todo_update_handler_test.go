package handler_test

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/mocoarow/todo-apps/backend-gin-gorm/domain"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func Test_TodoHandler_UpdateTodo_shouldReturn400_whenInvalidRequest(t *testing.T) {
	t.Parallel()

	// given
	userID := randomUserID()
	tests := []struct {
		name string
		body io.Reader
	}{
		{
			name: "empty request",
			body: bytes.NewBufferString("{}"),
		},
		{
			name: "empty text",
			body: bytes.NewBufferString(`{"text": ""}`),
		},
		{
			name: "nil request",
			body: nil,
		},
		{
			name: "251 characters text",
			body: bytes.NewBufferString(`{"text": "` + strings.Repeat("a", 251) + `"}`),
		},
		{
			name: "251 characters japanese text",
			body: bytes.NewBufferString(`{"text": "` + strings.Repeat("あ", 251) + `"}`),
		},
		{
			name: "251 characters emoji text",
			body: bytes.NewBufferString(`{"text": "` + strings.Repeat("😀", 251) + `"}`),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			ctx := context.Background()
			todoUsecase := NewMockTodoUsecase(t)
			r := initTodoRouter(t, ctx, todoUsecase, userID)
			w := httptest.NewRecorder()

			// when
			req, err := http.NewRequestWithContext(ctx, http.MethodPut, "/api/v1/todo/1", tt.body)
			require.NoError(t, err)
			r.ServeHTTP(w, req)
			respBytes := readBytes(t, w.Body)

			// then
			assert.Equal(t, http.StatusBadRequest, w.Code, "status code should be 400")
			validateErrorResponse(t, respBytes, "invalid_request", "request body is invalid")
		})
	}
}

func Test_TodoHandler_UpdateTodo_shouldReturn400_whenInvalidPath(t *testing.T) {
	t.Parallel()
	ctx := context.Background()

	// given
	userID := randomUserID()
	todoUsecase := NewMockTodoUsecase(t)
	r := initTodoRouter(t, ctx, todoUsecase, userID)
	w := httptest.NewRecorder()

	// when
	req, err := http.NewRequestWithContext(ctx, http.MethodPut, "/api/v1/todo/invalid", nil)
	require.NoError(t, err)
	r.ServeHTTP(w, req)
	respBytes := readBytes(t, w.Body)

	// then
	// - status code
	assert.Equal(t, http.StatusBadRequest, w.Code, "status code should be 400")
	validateErrorResponse(t, respBytes, "invalid_todo_id", "todo id must be a positive integer")
}

func Test_TodoHandler_UpdateTodo_shouldReturn404_whenTodoNotFound(t *testing.T) {
	t.Parallel()
	ctx := context.Background()

	// given
	userID := randomUserID()
	todoUsecase := NewMockTodoUsecase(t)
	todoUsecase.EXPECT().UpdateTodo(mock.Anything, &domain.UpdateTodoInput{
		ID:         1,
		UserID:     userID,
		Text:       "task 1",
		IsComplete: false,
	}).Return(nil, domain.ErrTodoNotFound).Once()
	r := initTodoRouter(t, ctx, todoUsecase, userID)
	w := httptest.NewRecorder()

	// when
	req, err := http.NewRequestWithContext(ctx, http.MethodPut, "/api/v1/todo/1", bytes.NewBufferString(`{"text": "task 1"}`))
	require.NoError(t, err)
	r.ServeHTTP(w, req)
	respBytes := readBytes(t, w.Body)

	// then
	assert.Equal(t, http.StatusNotFound, w.Code, "status code should be 404")
	validateErrorResponse(t, respBytes, "todo_not_found", "Not Found")
}

func Test_TodoHandler_UpdateTodo_shouldReturn200_whenValidRequest(t *testing.T) {
	t.Parallel()

	// given
	userID := randomUserID()
	tests := []struct {
		name               string
		body               io.Reader
		expectedText       string
		expectedIsComplete bool
	}{
		{
			name:               "normal text",
			body:               bytes.NewBufferString(`{"text": "task 1"}`),
			expectedText:       "task 1",
			expectedIsComplete: false,
		},
		{
			name:               "japanese text",
			body:               bytes.NewBufferString(`{"text": "タスク1"}`),
			expectedText:       "タスク1",
			expectedIsComplete: false,
		},
		{
			name:               "unicode text",
			body:               bytes.NewBufferString(`{"text": "😀"}`),
			expectedText:       "😀",
			expectedIsComplete: false,
		},
		{
			name:               "250 characters text",
			body:               bytes.NewBufferString(`{"text": "` + strings.Repeat("a", 250) + `"}`),
			expectedText:       strings.Repeat("a", 250),
			expectedIsComplete: false,
		},
		{
			name:               "250 characters japanese text",
			body:               bytes.NewBufferString(`{"text": "` + strings.Repeat("あ", 250) + `"}`),
			expectedText:       strings.Repeat("あ", 250),
			expectedIsComplete: false,
		},
		{
			name:               "250 characters emoji text",
			body:               bytes.NewBufferString(`{"text": "` + strings.Repeat("😀", 250) + `"}`),
			expectedText:       strings.Repeat("😀", 250),
			expectedIsComplete: false,
		},
		{
			name:               "complete todo",
			body:               bytes.NewBufferString(`{"text": "task 1", "isComplete": true}`),
			expectedText:       "task 1",
			expectedIsComplete: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			ctx := context.Background()
			// given
			todoUsecase := NewMockTodoUsecase(t)
			todoUsecase.EXPECT().UpdateTodo(mock.Anything, &domain.UpdateTodoInput{
				ID:         1,
				UserID:     userID,
				Text:       tt.expectedText,
				IsComplete: tt.expectedIsComplete,
			}).Return(&domain.UpdateTodoOutput{
				Todo: &domain.Todo{
					ID:         1,
					Text:       tt.expectedText,
					IsComplete: tt.expectedIsComplete,
				},
			}, nil).Once()
			r := initTodoRouter(t, ctx, todoUsecase, userID)
			w := httptest.NewRecorder()

			// when
			req, err := http.NewRequestWithContext(ctx, http.MethodPut, "/api/v1/todo/1", tt.body)
			require.NoError(t, err)
			r.ServeHTTP(w, req)
			respBytes := readBytes(t, w.Body)

			// then
			// - status code
			assert.Equal(t, http.StatusOK, w.Code, "status code should be 200")

			jsonObj := parseJSON(t, respBytes)

			// - id
			idExpr := parseExpr(t, "$.id")
			id := idExpr.Get(jsonObj)
			require.Len(t, id, 1, "response should have one id")
			assert.Positive(t, id[0], "id should be greater than 0")

			// - text
			textExpr := parseExpr(t, "$.text")
			text := textExpr.Get(jsonObj)
			require.Len(t, text, 1, "response should have one text")
			assert.Equal(t, tt.expectedText, text[0])

			// - isComplete
			isCompleteExpr := parseExpr(t, "$.isComplete")
			isComplete := isCompleteExpr.Get(jsonObj)
			require.Len(t, isComplete, 1, "response should have one isComplete")
			assert.Equal(t, tt.expectedIsComplete, isComplete[0])
		})
	}
}

func Test_TodoHandler_UpdateTodo_shouldReturn500_whenUsecaseReturnsError(t *testing.T) {
	t.Parallel()
	ctx := context.Background()

	// given
	userID := randomUserID()
	todoUsecase := NewMockTodoUsecase(t)
	todoUsecase.EXPECT().UpdateTodo(mock.Anything, &domain.UpdateTodoInput{
		ID:         1,
		UserID:     userID,
		Text:       "task 1",
		IsComplete: false,
	}).Return(nil, assert.AnError).Once()
	r := initTodoRouter(t, ctx, todoUsecase, userID)
	w := httptest.NewRecorder()
	// when
	req, err := http.NewRequestWithContext(ctx, http.MethodPut, "/api/v1/todo/1", bytes.NewBufferString(`{"text": "task 1"}`))
	require.NoError(t, err)
	r.ServeHTTP(w, req)
	respBytes := readBytes(t, w.Body)

	// then
	// - status code
	assert.Equal(t, http.StatusInternalServerError, w.Code, "status code should be 500")
	validateErrorResponse(t, respBytes, "internal_server_error", "Internal Server Error")
}

func Test_TodoHandler_UpdateTodo_shouldReturn500_whenUsecaseReturnsInvalidTodo(t *testing.T) {
	t.Parallel()
	ctx := context.Background()

	// given
	userID := randomUserID()
	todoUsecase := NewMockTodoUsecase(t)
	todoUsecase.EXPECT().UpdateTodo(mock.Anything, &domain.UpdateTodoInput{
		ID:         1,
		UserID:     userID,
		Text:       "task 1",
		IsComplete: false,
	}).Return(&domain.UpdateTodoOutput{}, nil).Once()
	r := initTodoRouter(t, ctx, todoUsecase, userID)
	w := httptest.NewRecorder()
	// when
	req, err := http.NewRequestWithContext(ctx, http.MethodPut, "/api/v1/todo/1", bytes.NewBufferString(`{"text": "task 1"}`))
	require.NoError(t, err)
	r.ServeHTTP(w, req)
	respBytes := readBytes(t, w.Body)

	// then
	assert.Equal(t, http.StatusInternalServerError, w.Code, "status code should be 500")
	validateErrorResponse(t, respBytes, "internal_server_error", "Internal Server Error")
}
