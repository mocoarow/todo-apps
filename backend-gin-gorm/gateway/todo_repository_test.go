package gateway_test

import (
	"context"
	"math/rand"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/mocoarow/todo-apps/backend-gin-gorm/domain"
	"github.com/mocoarow/todo-apps/backend-gin-gorm/gateway"
)

const (
	timeMargin = 2 * time.Second // Margin for timestamp comparison
)

// cleanupTodoTable deletes all entries in the todo table for a specific user.
func cleanupTodoTable(t *testing.T, userID int) {
	t.Helper()
	if err := db.Exec("DELETE FROM todo WHERE user_id = ?", userID).Error; err != nil {
		t.Fatalf("Failed to delete from table todo: %v", err)
	}
}

// FindTodos Tests

func TestTodoRepository_FindTodos_shouldReturnEmptyList_whenNoTodosExist(t *testing.T) {
	t.Parallel()
	ctx := context.Background()
	userID := rand.Intn(1000000) //nolint:gosec

	// given
	cleanupTodoTable(t, userID)
	repo := gateway.NewTodoRepository(db)

	// when
	todos, err := repo.FindTodos(ctx, userID)
	require.NoError(t, err, "FindTodos() should not return an error")

	// then
	require.Empty(t, todos, "FindTodos() should return an empty list")
}

func TestTodoRepository_FindTodos_shouldReturnSingleTodo_whenOneTodoExists(t *testing.T) {
	t.Parallel()
	ctx := context.Background()
	userID := rand.Intn(1000000) //nolint:gosec

	// given
	cleanupTodoTable(t, userID)
	repo := gateway.NewTodoRepository(db)

	// - Insert test data
	input, err := domain.NewCreateTodoInput(userID, "Test Todo")
	require.NoError(t, err, "Failed to create input")
	_, err = repo.CreateTodo(ctx, input)
	require.NoError(t, err, "Failed to insert test data")

	// when
	todos, err := repo.FindTodos(ctx, userID)
	require.NoError(t, err, "FindTodos() should not return an error")

	// then
	assert.Len(t, todos, 1, "FindTodos() should return 1 todo")
}

func TestTodoRepository_FindTodos_shouldReturnMultipleTodos_whenMultipleTodosExist(t *testing.T) {
	t.Parallel()
	ctx := context.Background()
	userID := rand.Intn(1000000) //nolint:gosec

	// given
	cleanupTodoTable(t, userID)
	repo := gateway.NewTodoRepository(db)

	// - Insert test data
	texts := []string{
		"First Todo",
		"Second Todo",
		"Third Todo",
	}
	for _, text := range texts {
		input, err := domain.NewCreateTodoInput(userID, text)
		require.NoError(t, err, "Failed to create input")
		_, err = repo.CreateTodo(ctx, input)
		require.NoError(t, err, "Failed to insert test data")
	}

	// when
	todos, err := repo.FindTodos(ctx, userID)
	require.NoError(t, err, "FindTodos() should not return an error")

	// then
	assert.Len(t, todos, 3, "FindTodos() should return 3 todos")
}

func TestTodoRepository_FindTodos_shouldReturnTodosInInsertionOrder_whenMultipleTodosExist(t *testing.T) {
	t.Parallel()
	ctx := context.Background()
	userID := rand.Intn(1000000) //nolint:gosec

	// given
	cleanupTodoTable(t, userID)
	repo := gateway.NewTodoRepository(db)

	// - Insert test data
	texts := []string{
		"First Todo",
		"Second Todo",
		"Third Todo",
	}
	for _, text := range texts {
		input, err := domain.NewCreateTodoInput(userID, text)
		require.NoError(t, err, "Failed to create input")
		_, err = repo.CreateTodo(ctx, input)
		require.NoError(t, err, "Failed to insert test data")
	}

	// when
	todos, err := repo.FindTodos(ctx, userID)
	require.NoError(t, err, "FindTodos() should not return an error")

	// then
	assert.Len(t, todos, 3, "FindTodos() should return 3 todos")
	assert.Equal(t, texts[0], todos[0].Text, "First todo Text should match")
	assert.Equal(t, texts[1], todos[1].Text, "Second todo Text should match")
	assert.Equal(t, texts[2], todos[2].Text, "Third todo Text should match")
}

// CreateTodo Tests
func TestTodoRepository_CreateTodo_shouldReturnValidTodo_whenTodoCreated(t *testing.T) {
	t.Parallel()
	ctx := context.Background()
	userID := rand.Intn(1000000) //nolint:gosec

	// given
	cleanupTodoTable(t, userID)
	repo := gateway.NewTodoRepository(db)

	// when
	text := "New Todo Item"
	before := time.Now().UTC()
	input, err := domain.NewCreateTodoInput(userID, text)
	require.NoError(t, err, "Failed to create input")
	todo, err := repo.CreateTodo(ctx, input)
	require.NoError(t, err, "CreateTodo() should not return an error")
	after := time.Now().UTC()

	// then
	assert.NotNil(t, todo, "CreateTodo() should return a todo")
	assert.Equal(t, text, todo.Text, "Todo text should match")
	assert.False(t, todo.IsComplete, "New todo should not be complete")
	assert.Positive(t, todo.ID, "Todo ID should be greater than 0")
	assert.NotZero(t, todo.CreatedAt, "CreatedAt should not be zero")
	assert.NotZero(t, todo.UpdatedAt, "UpdatedAt should not be zero")
	assert.True(t, todo.CreatedAt.Equal(todo.UpdatedAt), "CreatedAt and UpdatedAt should be equal for new todo")
	assert.True(t, todo.CreatedAt.After(before.Add(-timeMargin)) && todo.CreatedAt.Before(after.Add(timeMargin)), "CreatedAt should be within expected range")
	assert.True(t, todo.UpdatedAt.After(before.Add(-timeMargin)) && todo.UpdatedAt.Before(after.Add(timeMargin)), "UpdatedAt should be within expected range")
}

func TestTodoRepository_CreateTodo_shouldPersistTodoInDatabase_whenTodoCreated(t *testing.T) {
	t.Parallel()
	ctx := context.Background()
	userID := rand.Intn(1000000) //nolint:gosec

	// given
	cleanupTodoTable(t, userID)
	repo := gateway.NewTodoRepository(db)

	// when
	text := "Persisted Todo"
	input, err := domain.NewCreateTodoInput(userID, text)
	require.NoError(t, err, "Failed to create input")
	createdTodo, err := repo.CreateTodo(ctx, input)
	require.NoError(t, err, "CreateTodo() should not return an error")

	// then
	// - Fetch from database to verify persistence
	todos, err := repo.FindTodos(ctx, userID)
	require.NoError(t, err, "FindTodos() should not return an error")
	require.Len(t, todos, 1, "FindTodos() should return 1 todo")

	persistedTodo := todos[0]
	assert.Equal(t, createdTodo.ID, persistedTodo.ID, "Persisted todo ID should match created todo ID")
	assert.Equal(t, text, persistedTodo.Text, "Persisted todo Text should match created todo Text")
	assert.False(t, persistedTodo.IsComplete, "Persisted todo should not be complete")
}

// UpdateTodo Tests

func TestTodoRepository_UpdateTodo_shouldUpdateTextAndIsComplete_whenValidInputProvided(t *testing.T) {
	t.Parallel()
	ctx := context.Background()
	userID := rand.Intn(1000000) //nolint:gosec

	// given
	cleanupTodoTable(t, userID)
	repo := gateway.NewTodoRepository(db)

	// - Create a todo first
	createInput, err := domain.NewCreateTodoInput(userID, "Original Text")
	require.NoError(t, err, "Failed to create input")
	createdTodo, err := repo.CreateTodo(ctx, createInput)
	require.NoError(t, err, "Failed to create todo")

	// when
	time.Sleep(10 * time.Millisecond) // Ensure UpdatedAt will be different
	updateInput, err := domain.NewUpdateTodoInput(createdTodo.ID, userID, "Updated Text", true)
	require.NoError(t, err, "Failed to create update input")
	updatedTodo, err := repo.UpdateTodo(ctx, updateInput)
	require.NoError(t, err, "UpdateTodo() should not return an error")

	// then
	assert.NotNil(t, updatedTodo, "UpdateTodo() should return a todo")
	assert.Equal(t, createdTodo.ID, updatedTodo.ID, "Todo ID should not change")
	assert.Equal(t, "Updated Text", updatedTodo.Text, "Text should be updated")
	assert.True(t, updatedTodo.IsComplete, "IsComplete should be updated to true")
	assert.True(t, createdTodo.CreatedAt.Equal(updatedTodo.CreatedAt), "CreatedAt should not change")
	assert.True(t, updatedTodo.UpdatedAt.After(createdTodo.UpdatedAt), "UpdatedAt should be after the original UpdatedAt")
}

func TestTodoRepository_UpdateTodo_shouldPersistChangesInDatabase_whenTodoUpdated(t *testing.T) {
	t.Parallel()
	ctx := context.Background()
	userID := rand.Intn(1000000) //nolint:gosec

	// given
	cleanupTodoTable(t, userID)
	repo := gateway.NewTodoRepository(db)

	// - Create a todo first
	createInput, err := domain.NewCreateTodoInput(userID, "Original Text")
	require.NoError(t, err, "Failed to create input")
	createdTodo, err := repo.CreateTodo(ctx, createInput)
	require.NoError(t, err, "Failed to create todo")

	// when
	updateInput, err := domain.NewUpdateTodoInput(createdTodo.ID, userID, "Updated Text", true)
	require.NoError(t, err, "Failed to create update input")
	_, err = repo.UpdateTodo(ctx, updateInput)
	require.NoError(t, err, "UpdateTodo() should not return an error")

	// then - Fetch from database to verify persistence
	todos, err := repo.FindTodos(ctx, userID)
	require.NoError(t, err, "FindTodos() should not return an error")
	require.Len(t, todos, 1, "FindTodos() should return 1 todo")

	persistedTodo := todos[0]
	assert.Equal(t, createdTodo.ID, persistedTodo.ID, "Persisted todo ID should match")
	assert.Equal(t, "Updated Text", persistedTodo.Text, "Persisted todo Text should match updated text")
	assert.True(t, persistedTodo.IsComplete, "Persisted todo should be complete")
}

func TestTodoRepository_UpdateTodo_shouldReturnError_whenTodoNotFound(t *testing.T) {
	t.Parallel()
	ctx := context.Background()
	userID := rand.Intn(1000000) //nolint:gosec

	// given
	cleanupTodoTable(t, userID)
	repo := gateway.NewTodoRepository(db)

	// when - Try to update a non-existent todo
	nonExistentID := 999999999
	updateInput, err := domain.NewUpdateTodoInput(nonExistentID, userID, "Updated Text", true)
	require.NoError(t, err, "Failed to create update input")
	_, err = repo.UpdateTodo(ctx, updateInput)

	// then
	assert.ErrorIs(t, err, domain.ErrTodoNotFound, "UpdateTodo() should return ErrTodoNotFound when todo not found")
}

func TestTodoRepository_UpdateTodo_shouldReturnError_whenUserIDDoesNotMatch(t *testing.T) {
	t.Parallel()
	ctx := context.Background()
	userID := rand.Intn(1000000) //nolint:gosec
	differentUserID := userID + 1

	// given
	cleanupTodoTable(t, userID)
	cleanupTodoTable(t, differentUserID)
	repo := gateway.NewTodoRepository(db)

	// Create a todo for userID
	createInput, err := domain.NewCreateTodoInput(userID, "Original Text")
	require.NoError(t, err, "Failed to create input")
	createdTodo, err := repo.CreateTodo(ctx, createInput)
	require.NoError(t, err, "Failed to create todo")

	// when - Try to update the todo with a different user ID
	updateInput, err := domain.NewUpdateTodoInput(createdTodo.ID, differentUserID, "Updated Text", true)
	require.NoError(t, err, "Failed to create update input")
	_, err = repo.UpdateTodo(ctx, updateInput)

	// then
	require.ErrorIs(t, err, domain.ErrTodoNotFound, "UpdateTodo() should return ErrTodoNotFound when userID does not match")

	// Verify the original todo was not updated
	todos, err := repo.FindTodos(ctx, userID)
	require.NoError(t, err, "FindTodos() should not return an error")
	require.Len(t, todos, 1, "FindTodos() should return 1 todo")
	assert.Equal(t, "Original Text", todos[0].Text, "Todo text should not be updated")
	assert.False(t, todos[0].IsComplete, "Todo should not be marked as complete")
}

// DeleteTodo Tests

func TestTodoRepository_DeleteTodo_shouldDeleteTodo_whenValidInputProvided(t *testing.T) {
	t.Parallel()
	ctx := context.Background()
	userID := rand.Intn(1000000) //nolint:gosec

	// given
	cleanupTodoTable(t, userID)
	repo := gateway.NewTodoRepository(db)

	// Create a todo first
	createInput, err := domain.NewCreateTodoInput(userID, "Todo to delete")
	require.NoError(t, err, "Failed to create input")
	createdTodo, err := repo.CreateTodo(ctx, createInput)
	require.NoError(t, err, "Failed to create todo")

	// when
	deleteInput, err := domain.NewDeleteTodoInput(createdTodo.ID, userID)
	require.NoError(t, err, "Failed to create delete input")
	err = repo.DeleteTodo(ctx, deleteInput)
	require.NoError(t, err, "DeleteTodo() should not return an error")

	// then - Verify the todo was deleted
	todos, err := repo.FindTodos(ctx, userID)
	require.NoError(t, err, "FindTodos() should not return an error")
	assert.Empty(t, todos, "FindTodos() should return an empty list after deletion")
}

func TestTodoRepository_DeleteTodo_shouldNotAffectOtherTodos_whenOneTodoDeleted(t *testing.T) {
	t.Parallel()
	ctx := context.Background()
	userID := rand.Intn(1000000) //nolint:gosec

	// given
	cleanupTodoTable(t, userID)
	repo := gateway.NewTodoRepository(db)

	// Create multiple todos
	todo1Input, err := domain.NewCreateTodoInput(userID, "Todo 1")
	require.NoError(t, err, "Failed to create input")
	todo1, err := repo.CreateTodo(ctx, todo1Input)
	require.NoError(t, err, "Failed to create todo 1")

	todo2Input, err := domain.NewCreateTodoInput(userID, "Todo 2")
	require.NoError(t, err, "Failed to create input")
	todo2, err := repo.CreateTodo(ctx, todo2Input)
	require.NoError(t, err, "Failed to create todo 2")

	// when - Delete the first todo
	deleteInput, err := domain.NewDeleteTodoInput(todo1.ID, userID)
	require.NoError(t, err, "Failed to create delete input")
	err = repo.DeleteTodo(ctx, deleteInput)
	require.NoError(t, err, "DeleteTodo() should not return an error")

	// then - Verify only the first todo was deleted
	todos, err := repo.FindTodos(ctx, userID)
	require.NoError(t, err, "FindTodos() should not return an error")
	require.Len(t, todos, 1, "FindTodos() should return 1 todo")
	assert.Equal(t, todo2.ID, todos[0].ID, "Remaining todo should be todo 2")
	assert.Equal(t, "Todo 2", todos[0].Text, "Remaining todo text should match")
}

func TestTodoRepository_DeleteTodo_shouldReturnError_whenTodoNotFound(t *testing.T) {
	t.Parallel()
	ctx := context.Background()
	userID := rand.Intn(1000000) //nolint:gosec

	// given
	cleanupTodoTable(t, userID)
	repo := gateway.NewTodoRepository(db)

	// when - Try to delete a non-existent todo
	nonExistentID := 999999999
	deleteInput, err := domain.NewDeleteTodoInput(nonExistentID, userID)
	require.NoError(t, err, "Failed to create delete input")
	err = repo.DeleteTodo(ctx, deleteInput)

	// then
	assert.ErrorIs(t, err, domain.ErrTodoNotFound, "DeleteTodo() should return ErrTodoNotFound when todo not found")
}

func TestTodoRepository_DeleteTodo_shouldReturnError_whenUserIDDoesNotMatch(t *testing.T) {
	t.Parallel()
	ctx := context.Background()
	userID := rand.Intn(1000000) //nolint:gosec
	differentUserID := userID + 1

	// given
	cleanupTodoTable(t, userID)
	cleanupTodoTable(t, differentUserID)
	repo := gateway.NewTodoRepository(db)

	// Create a todo for userID
	createInput, err := domain.NewCreateTodoInput(userID, "Todo to protect")
	require.NoError(t, err, "Failed to create input")
	createdTodo, err := repo.CreateTodo(ctx, createInput)
	require.NoError(t, err, "Failed to create todo")

	// when - Try to delete the todo with a different user ID
	deleteInput, err := domain.NewDeleteTodoInput(createdTodo.ID, differentUserID)
	require.NoError(t, err, "Failed to create delete input")
	err = repo.DeleteTodo(ctx, deleteInput)

	// then
	require.ErrorIs(t, err, domain.ErrTodoNotFound, "DeleteTodo() should return ErrTodoNotFound when userID does not match")

	// Verify the original todo was not deleted
	todos, err := repo.FindTodos(ctx, userID)
	require.NoError(t, err, "FindTodos() should not return an error")
	require.Len(t, todos, 1, "FindTodos() should return 1 todo")
	assert.Equal(t, createdTodo.ID, todos[0].ID, "Todo should not be deleted")
	assert.Equal(t, "Todo to protect", todos[0].Text, "Todo text should match")
}
