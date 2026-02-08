package usecase_test

import (
	"context"
	"math/rand"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/mocoarow/todo-apps/backend-gin-gorm/domain"
	"github.com/mocoarow/todo-apps/backend-gin-gorm/gateway"
	"github.com/mocoarow/todo-apps/backend-gin-gorm/usecase"
)

func Test_UpdateTodoCommand_Execute_shouldReturnUpdatedTodo_whenValidInput(t *testing.T) {
	t.Parallel()
	ctx := context.Background()
	userID := rand.Intn(1000000) //nolint:gosec

	// given
	cleanupTodoTable(t, userID)
	repo := gateway.NewTodoRepository(dbc.DB)
	cmd := usecase.NewUpdateTodoCommand(repo)

	createInput, err := domain.NewCreateTodoInput(userID, "original")
	require.NoError(t, err)
	created, err := repo.CreateTodo(ctx, createInput)
	require.NoError(t, err)

	updateInput, err := domain.NewUpdateTodoInput(created.ID, userID, "updated", true)
	require.NoError(t, err)

	// when
	output, err := cmd.Execute(ctx, updateInput)

	// then
	require.NoError(t, err)
	require.NotNil(t, output)
	assert.Equal(t, "updated", output.Todo.Text)
	assert.True(t, output.Todo.IsComplete)

	// DB にも反映されていることを確認
	todos, err := repo.FindTodos(ctx, userID)
	require.NoError(t, err)
	require.Len(t, todos, 1)
	assert.Equal(t, "updated", todos[0].Text)
	assert.True(t, todos[0].IsComplete)
}

func Test_UpdateTodoCommand_Execute_shouldReturnError_whenTodoNotFound(t *testing.T) {
	t.Parallel()
	ctx := context.Background()
	userID := rand.Intn(1000000) //nolint:gosec

	// given
	cleanupTodoTable(t, userID)
	repo := gateway.NewTodoRepository(dbc.DB)
	cmd := usecase.NewUpdateTodoCommand(repo)

	updateInput, err := domain.NewUpdateTodoInput(999999999, userID, "updated", true)
	require.NoError(t, err)

	// when
	output, err := cmd.Execute(ctx, updateInput)

	// then
	require.Error(t, err)
	assert.Nil(t, output)
	assert.ErrorIs(t, err, domain.ErrTodoNotFound)
}

func Test_UpdateTodoCommand_Execute_shouldReturnError_whenUserIDDoesNotMatch(t *testing.T) {
	t.Parallel()
	ctx := context.Background()
	userID := rand.Intn(1000000) //nolint:gosec
	otherUserID := userID + 1

	// given
	cleanupTodoTable(t, userID)
	cleanupTodoTable(t, otherUserID)
	repo := gateway.NewTodoRepository(dbc.DB)
	cmd := usecase.NewUpdateTodoCommand(repo)

	createInput, err := domain.NewCreateTodoInput(userID, "original")
	require.NoError(t, err)
	created, err := repo.CreateTodo(ctx, createInput)
	require.NoError(t, err)

	// 別ユーザーの ID で更新を試みる
	updateInput, err := domain.NewUpdateTodoInput(created.ID, otherUserID, "hacked", true)
	require.NoError(t, err)

	// when
	output, err := cmd.Execute(ctx, updateInput)

	// then
	require.Error(t, err)
	assert.Nil(t, output)
	require.ErrorIs(t, err, domain.ErrTodoNotFound)

	// 元の todo が変更されていないことを確認
	todos, err := repo.FindTodos(ctx, userID)
	require.NoError(t, err)
	require.Len(t, todos, 1)
	assert.Equal(t, "original", todos[0].Text)
	assert.False(t, todos[0].IsComplete)
}
