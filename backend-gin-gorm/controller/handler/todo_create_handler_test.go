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

func Test_TodoHandler_CreateTodo_shouldReturn400_whenInvalidRequest(t *testing.T) {
	t.Parallel()
	ctx := context.Background()

	// given
	userID := randomUserID()
	todoUsecase := NewMockTodoUsecase(t)
	r := initTodoRouter(t, ctx, todoUsecase, userID)

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
			w := httptest.NewRecorder()
			// when
			req, err := http.NewRequestWithContext(ctx, http.MethodPost, "/api/v1/todo", tt.body)
			require.NoError(t, err)
			r.ServeHTTP(w, req)
			respBytes := readBytes(t, w.Body)

			// then
			assert.Equal(t, http.StatusBadRequest, w.Code, "status code should be 400")
			validateErrorResponse(t, respBytes, "invalid_request", "request body is invalid")
		})
	}
}

func Test_TodoHandler_CreateTodo_shouldReturn201_whenValidRequest(t *testing.T) {
	t.Parallel()

	// given
	userID := randomUserID()

	tests := []struct {
		name         string
		body         io.Reader
		expectedText string
	}{
		{
			name:         "normal text",
			body:         bytes.NewBufferString(`{"text": "task 1"}`),
			expectedText: "task 1",
		},
		{
			name:         "japanese text",
			body:         bytes.NewBufferString(`{"text": "タスク1"}`),
			expectedText: "タスク1",
		},
		{
			name:         "unicode text",
			body:         bytes.NewBufferString(`{"text": "😀"}`),
			expectedText: "😀",
		},
		{
			name:         "250 characters text",
			body:         bytes.NewBufferString(`{"text": "` + strings.Repeat("a", 250) + `"}`),
			expectedText: strings.Repeat("a", 250),
		},
		{
			name:         "250 characters japanese text",
			body:         bytes.NewBufferString(`{"text": "` + strings.Repeat("あ", 250) + `"}`),
			expectedText: strings.Repeat("あ", 250),
		},
		{
			name:         "250 characters emoji text",
			body:         bytes.NewBufferString(`{"text": "` + strings.Repeat("😀", 250) + `"}`),
			expectedText: strings.Repeat("😀", 250),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			ctx := context.Background()
			todoUsecase := NewMockTodoUsecase(t)
			todoUsecase.EXPECT().CreateTodo(mock.Anything, &domain.CreateTodoInput{
				UserID: userID,
				Text:   tt.expectedText,
			}).Return(&domain.CreateTodoOutput{
				Todo: &domain.Todo{
					ID:         123,
					Text:       tt.expectedText,
					IsComplete: false,
				},
			}, nil)
			r := initTodoRouter(t, ctx, todoUsecase, userID)
			w := httptest.NewRecorder()
			// when
			req, err := http.NewRequestWithContext(ctx, http.MethodPost, "/api/v1/todo", tt.body)
			require.NoError(t, err)
			r.ServeHTTP(w, req)
			respBytes := readBytes(t, w.Body)

			// then
			// - status code
			assert.Equal(t, http.StatusCreated, w.Code, "status code should be 201")

			jsonObj := parseJSON(t, respBytes)

			// - id
			idExpr := parseExpr(t, "$.id")
			id := idExpr.Get(jsonObj)
			require.Len(t, id, 1, "response should have one id")
			assert.Equal(t, int64(123), id[0], "id should be 123")

			// - text
			textExpr := parseExpr(t, "$.text")
			text := textExpr.Get(jsonObj)
			require.Len(t, text, 1, "response should have one text")
			assert.Equal(t, tt.expectedText, text[0])

			// - isComplete
			isCompleteExpr := parseExpr(t, "$.isComplete")
			isComplete := isCompleteExpr.Get(jsonObj)
			require.Len(t, isComplete, 1, "response should have one isComplete")
			assert.Equal(t, false, isComplete[0], "isComplete should be false")
		})
	}
}

func Test_TodoHandler_CreateTodo_shouldReturn500_whenUsecaseReturnsError(t *testing.T) {
	t.Parallel()
	ctx := context.Background()

	// given
	userID := randomUserID()
	todoUsecase := NewMockTodoUsecase(t)
	todoUsecase.EXPECT().CreateTodo(mock.Anything, mock.Anything).Return(nil, assert.AnError).Once()
	r := initTodoRouter(t, ctx, todoUsecase, userID)
	w := httptest.NewRecorder()
	// when
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, "/api/v1/todo", bytes.NewBufferString(`{"text": "task 1"}`))
	require.NoError(t, err)
	r.ServeHTTP(w, req)
	respBytes := readBytes(t, w.Body)

	// then
	assert.Equal(t, http.StatusInternalServerError, w.Code, "status code should be 500")
	validateErrorResponse(t, respBytes, "internal_server_error", "Internal Server Error")
}

func Test_TodoHandler_CreateTodo_shouldReturn500_whenUsecaseReturnsInvalidTodo(t *testing.T) {
	t.Parallel()
	ctx := context.Background()

	// given
	userID := randomUserID()
	todoUsecase := NewMockTodoUsecase(t)
	todoUsecase.EXPECT().CreateTodo(mock.Anything, mock.Anything).Return(&domain.CreateTodoOutput{}, nil).Once()
	r := initTodoRouter(t, ctx, todoUsecase, userID)
	w := httptest.NewRecorder()
	// when
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, "/api/v1/todo", bytes.NewBufferString(`{"text": "task 1"}`))
	require.NoError(t, err)
	r.ServeHTTP(w, req)
	respBytes := readBytes(t, w.Body)

	// then
	assert.Equal(t, http.StatusInternalServerError, w.Code, "status code should be 500")
	validateErrorResponse(t, respBytes, "internal_server_error", "Internal Server Error")
}
