package domain_test

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/mocoarow/todo-apps/backend-gin-gorm/domain"
)

// NewTodo tests
func TestNewTodo_shouldReturnTodo_whenValidInput(t *testing.T) {
	t.Parallel()

	// given
	now := time.Now()

	// when
	todo, err := domain.NewTodo(1, 2, "Test todo", false, now, now)

	// then
	require.NoError(t, err, "expected no error for valid Todo")
	assert.NotNil(t, todo, "expected non-nil Todo")
	assert.Equal(t, 1, todo.ID, "expected ID to match")
	assert.Equal(t, "Test todo", todo.Text, "expected Text to match")
	assert.False(t, todo.IsComplete, "expected IsComplete to match")
	assert.Equal(t, now, todo.CreatedAt, "expected CreatedAt to match")
	assert.Equal(t, now, todo.UpdatedAt, "expected UpdatedAt to match")
}

func TestNewTodo_shouldReturnError_whenInvalidInput(t *testing.T) {
	t.Parallel()

	// given
	now := time.Now()

	tests := []struct {
		name   string
		id     int
		userID int
		text   string
	}{
		{
			name:   "ID is zero",
			id:     0,
			userID: 1,
			text:   "Test todo",
		},
		{
			name:   "ID is negative",
			id:     -1,
			userID: 1,
			text:   "Test todo",
		},
		{
			name:   "USER ID is zero",
			id:     1,
			userID: 0,
			text:   "Test todo",
		},
		{
			name:   "User ID is negative",
			id:     1,
			userID: -1,
			text:   "Test todo",
		},
		{
			name:   "text is empty",
			id:     1,
			userID: 1,
			text:   "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			// when
			todo, err := domain.NewTodo(tt.id, tt.userID, tt.text, false, now, now)

			// then
			require.Error(t, err, "expected error for invalid input")
			assert.Nil(t, todo, "expected nil Todo")
			assert.Contains(t, err.Error(), "validate todo model", "error should mention validation")
		})
	}
}

// NewCreateTodoInput tests
func TestNewCreateTodoInput_shouldReturnInput_whenValidInput(t *testing.T) {
	t.Parallel()

	// when
	input, err := domain.NewCreateTodoInput(1, "Test todo")

	// then
	require.NoError(t, err, "expected no error for valid CreateTodoInput")
	assert.NotNil(t, input, "expected non-nil CreateTodoInput")
	assert.Equal(t, 1, input.UserID, "expected UserID to match")
	assert.Equal(t, "Test todo", input.Text, "expected Text to match")
}

func TestNewCreateTodoInput_shouldReturnError_whenInvalidInput(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name   string
		userID int
		text   string
	}{
		{
			name:   "UserID is zero",
			userID: 0,
			text:   "Test todo",
		},
		{
			name:   "UserID is negative",
			userID: -1,
			text:   "Test todo",
		},
		{
			name:   "text is empty",
			userID: 1,
			text:   "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			// when
			input, err := domain.NewCreateTodoInput(tt.userID, tt.text)

			// then
			require.Error(t, err, "expected error for invalid input")
			assert.Nil(t, input, "expected nil CreateTodoInput")
			assert.Contains(t, err.Error(), "validate create todo input", "error should mention validation")
		})
	}
}

// NewCreateTodoOutput tests
func TestNewCreateTodoOutput_shouldReturnOutput_whenValidInput(t *testing.T) {
	t.Parallel()

	// given
	now := time.Now()
	todo, err := domain.NewTodo(1, 2, "Test todo", false, now, now)
	require.NoError(t, err)

	// when
	output, err := domain.NewCreateTodoOutput(todo)

	// then
	require.NoError(t, err, "expected no error for valid CreateTodoOutput")
	assert.NotNil(t, output, "expected non-nil CreateTodoOutput")
	assert.Equal(t, todo, output.Todo, "expected Todo to match")
}

func TestNewCreateTodoOutput_shouldReturnError_whenTodoIsNil(t *testing.T) {
	t.Parallel()

	// when
	output, err := domain.NewCreateTodoOutput(nil)

	// then
	require.Error(t, err, "expected error for nil Todo")
	assert.Nil(t, output, "expected nil CreateTodoOutput")
	assert.Contains(t, err.Error(), "validate create todo output: Key: 'CreateTodoOutput.Todo' Error:Field validation for 'Todo' failed on the 'required' tag", "error should mention validation")
}

// NewUpdateTodoInput tests
func TestNewUpdateTodoInput_shouldReturnInput_whenValidInput(t *testing.T) {
	t.Parallel()

	// when
	input, err := domain.NewUpdateTodoInput(1, 2, "Updated todo", true)

	// then
	require.NoError(t, err, "expected no error for valid UpdateTodoInput")
	assert.NotNil(t, input, "expected non-nil UpdateTodoInput")
	assert.Equal(t, 1, input.ID, "expected ID to match")
	assert.Equal(t, 2, input.UserID, "expected UserID to match")
	assert.Equal(t, "Updated todo", input.Text, "expected Text to match")
	assert.True(t, input.IsComplete, "expected IsComplete to match")
}

func TestNewUpdateTodoInput_shouldReturnError_whenInvalidInput(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name   string
		id     int
		userID int
		text   string
	}{
		{
			name:   "ID is zero",
			id:     0,
			userID: 1,
			text:   "Updated todo",
		},
		{
			name:   "UserID is zero",
			id:     1,
			userID: 0,
			text:   "Updated todo",
		},
		{
			name:   "text is empty",
			id:     1,
			userID: 2,
			text:   "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			// when
			input, err := domain.NewUpdateTodoInput(tt.id, tt.userID, tt.text, true)

			// then
			require.Error(t, err, "expected error for invalid input")
			assert.Nil(t, input, "expected nil UpdateTodoInput")
			assert.Contains(t, err.Error(), "validate update todo input", "error should mention validation")
		})
	}
}

// NewUpdateTodoOutput tests
func TestNewUpdateTodoOutput_shouldReturnOutput_whenValidInput(t *testing.T) {
	t.Parallel()

	// given
	now := time.Now()
	todo, err := domain.NewTodo(1, 2, "Test todo", false, now, now)
	require.NoError(t, err)

	// when
	output, err := domain.NewUpdateTodoOutput(todo)

	// then
	require.NoError(t, err, "expected no error for valid UpdateTodoOutput")
	assert.NotNil(t, output, "expected non-nil UpdateTodoOutput")
	assert.Equal(t, todo, output.Todo, "expected Todo to match")
}

func TestNewUpdateTodoOutput_shouldReturnError_whenTodoIsNil(t *testing.T) {
	t.Parallel()

	// when
	output, err := domain.NewUpdateTodoOutput(nil)

	// then
	require.Error(t, err, "expected error for nil Todo")
	assert.Nil(t, output, "expected nil UpdateTodoOutput")
	assert.Contains(t, err.Error(), "validate update todo output: Key: 'UpdateTodoOutput.Todo' Error:Field validation for 'Todo' failed on the 'required' tag", "error should mention validation")
}

// NewDeleteTodoInput tests
func TestNewDeleteTodoInput_shouldReturnInput_whenValidInput(t *testing.T) {
	t.Parallel()

	// when
	input, err := domain.NewDeleteTodoInput(1, 2)

	// then
	require.NoError(t, err, "expected no error for valid DeleteTodoInput")
	assert.NotNil(t, input, "expected non-nil DeleteTodoInput")
	assert.Equal(t, 1, input.ID, "expected ID to match")
	assert.Equal(t, 2, input.UserID, "expected UserID to match")
}

func TestNewDeleteTodoInput_shouldReturnError_whenInvalidInput(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name   string
		id     int
		userID int
	}{
		{
			name:   "ID is zero",
			id:     0,
			userID: 1,
		},
		{
			name:   "ID is negative",
			id:     -1,
			userID: 1,
		},
		{
			name:   "UserID is zero",
			id:     1,
			userID: 0,
		},
		{
			name:   "UserID is negative",
			id:     1,
			userID: -1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			// when
			input, err := domain.NewDeleteTodoInput(tt.id, tt.userID)

			// then
			require.Error(t, err, "expected error for invalid input")
			assert.Nil(t, input, "expected nil DeleteTodoInput")
			assert.Contains(t, err.Error(), "validate delete todo input", "error should mention validation")
		})
	}
}
