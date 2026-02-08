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

func Test_DeleteTodoCommand_Execute_shouldDeleteTodo_whenValidInput(t *testing.T) {
	t.Parallel()
	ctx := context.Background()
	userID := rand.Intn(1000000) //nolint:gosec

	// given
	cleanupTodoTable(t, userID)
	repo := gateway.NewTodoRepository(dbc.DB)
	cmd := usecase.NewDeleteTodoCommand(repo)

	createInput, err := domain.NewCreateTodoInput(userID, "to be deleted")
	require.NoError(t, err)
	created, err := repo.CreateTodo(ctx, createInput)
	require.NoError(t, err)

	deleteInput, err := domain.NewDeleteTodoInput(created.ID, userID)
	require.NoError(t, err)

	// when
	err = cmd.Execute(ctx, deleteInput)

	// then
	require.NoError(t, err)

	// DB からも削除されていることを確認
	todos, err := repo.FindTodos(ctx, userID)
	require.NoError(t, err)
	assert.Empty(t, todos)
}

func Test_DeleteTodoCommand_Execute_shouldReturnError_whenTodoNotFound(t *testing.T) {
	t.Parallel()
	ctx := context.Background()
	userID := rand.Intn(1000000) //nolint:gosec

	// given
	cleanupTodoTable(t, userID)
	repo := gateway.NewTodoRepository(dbc.DB)
	cmd := usecase.NewDeleteTodoCommand(repo)

	deleteInput, err := domain.NewDeleteTodoInput(999999999, userID)
	require.NoError(t, err)

	// when
	err = cmd.Execute(ctx, deleteInput)

	// then
	require.Error(t, err)
	assert.ErrorIs(t, err, domain.ErrTodoNotFound)
}

func Test_DeleteTodoCommand_Execute_shouldReturnError_whenUserIDDoesNotMatch(t *testing.T) {
	t.Parallel()
	ctx := context.Background()
	userID := rand.Intn(1000000) //nolint:gosec
	otherUserID := userID + 1

	// given
	cleanupTodoTable(t, userID)
	cleanupTodoTable(t, otherUserID)
	repo := gateway.NewTodoRepository(dbc.DB)
	cmd := usecase.NewDeleteTodoCommand(repo)

	createInput, err := domain.NewCreateTodoInput(userID, "protected")
	require.NoError(t, err)
	created, err := repo.CreateTodo(ctx, createInput)
	require.NoError(t, err)

	// 別ユーザーの ID で削除を試みる
	deleteInput, err := domain.NewDeleteTodoInput(created.ID, otherUserID)
	require.NoError(t, err)

	// when
	err = cmd.Execute(ctx, deleteInput)

	// then
	require.Error(t, err)
	require.ErrorIs(t, err, domain.ErrTodoNotFound)

	// 元の todo が削除されていないことを確認
	todos, err := repo.FindTodos(ctx, userID)
	require.NoError(t, err)
	assert.Len(t, todos, 1)
	assert.Equal(t, "protected", todos[0].Text)
}

func Test_DeleteTodoCommand_Execute_shouldNotAffectOtherTodos_whenOneTodoDeleted(t *testing.T) {
	t.Parallel()
	ctx := context.Background()
	userID := rand.Intn(1000000) //nolint:gosec

	// given
	cleanupTodoTable(t, userID)
	repo := gateway.NewTodoRepository(dbc.DB)
	cmd := usecase.NewDeleteTodoCommand(repo)

	input1, err := domain.NewCreateTodoInput(userID, "task1")
	require.NoError(t, err)
	todo1, err := repo.CreateTodo(ctx, input1)
	require.NoError(t, err)

	input2, err := domain.NewCreateTodoInput(userID, "task2")
	require.NoError(t, err)
	_, err = repo.CreateTodo(ctx, input2)
	require.NoError(t, err)

	deleteInput, err := domain.NewDeleteTodoInput(todo1.ID, userID)
	require.NoError(t, err)

	// when
	err = cmd.Execute(ctx, deleteInput)

	// then
	require.NoError(t, err)

	// task2 だけが残っていることを確認
	todos, err := repo.FindTodos(ctx, userID)
	require.NoError(t, err)
	require.Len(t, todos, 1)
	assert.Equal(t, "task2", todos[0].Text)
}
