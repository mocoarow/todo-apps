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

func Test_FindTodosQuery_Execute_shouldReturnEmptyList_whenNoTodosExist(t *testing.T) {
	t.Parallel()
	ctx := context.Background()
	userID := rand.Intn(1000000) //nolint:gosec

	// given
	cleanupTodoTable(t, userID)
	repo := gateway.NewTodoRepository(dbc.DB)
	query := usecase.NewFindTodosQuery(repo)

	// when
	todos, err := query.Execute(ctx, userID)

	// then
	require.NoError(t, err)
	assert.Empty(t, todos)
}

func Test_FindTodosQuery_Execute_shouldReturnTodos_whenTodosExist(t *testing.T) {
	t.Parallel()
	ctx := context.Background()
	userID := rand.Intn(1000000) //nolint:gosec

	// given
	cleanupTodoTable(t, userID)
	repo := gateway.NewTodoRepository(dbc.DB)
	query := usecase.NewFindTodosQuery(repo)

	for _, text := range []string{"task1", "task2", "task3"} {
		input, err := domain.NewCreateTodoInput(userID, text)
		require.NoError(t, err)
		_, err = repo.CreateTodo(ctx, input)
		require.NoError(t, err)
	}

	// when
	todos, err := query.Execute(ctx, userID)

	// then
	require.NoError(t, err)
	assert.Len(t, todos, 3)
	assert.Equal(t, "task1", todos[0].Text)
	assert.Equal(t, "task2", todos[1].Text)
	assert.Equal(t, "task3", todos[2].Text)
}

func Test_FindTodosQuery_Execute_shouldNotReturnOtherUserTodos(t *testing.T) {
	t.Parallel()
	ctx := context.Background()
	userID := rand.Intn(1000000) //nolint:gosec
	otherUserID := userID + 1

	// given
	cleanupTodoTable(t, userID)
	cleanupTodoTable(t, otherUserID)
	repo := gateway.NewTodoRepository(dbc.DB)
	query := usecase.NewFindTodosQuery(repo)

	// 別ユーザーの todo を作成
	input, err := domain.NewCreateTodoInput(otherUserID, "other user task")
	require.NoError(t, err)
	_, err = repo.CreateTodo(ctx, input)
	require.NoError(t, err)

	// when
	todos, err := query.Execute(ctx, userID)

	// then
	require.NoError(t, err)
	assert.Empty(t, todos)
}
