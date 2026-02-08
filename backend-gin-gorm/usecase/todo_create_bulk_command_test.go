package usecase_test

import (
	"context"
	"math/rand"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/mocoarow/todo-apps/backend-gin-gorm/domain"
	"github.com/mocoarow/todo-apps/backend-gin-gorm/gateway"
	"github.com/mocoarow/todo-apps/backend-gin-gorm/usecase"
)

func Test_CreateBulkTodosCommand_Execute_shouldCreateAllTodos_whenAllSucceed(t *testing.T) {
	t.Parallel()
	ctx := context.Background()
	userID := rand.Intn(1000000) //nolint:gosec

	// given
	cleanupTodoTable(t, userID)
	txManager := gateway.NewTodoCreateBulkCommandTxManager(dbc)
	cmd := usecase.NewCreateBulkTodosCommand(txManager)

	input, err := domain.NewCreateBulkTodosInput(userID, []domain.CreateTodoInput{
		{UserID: userID, Text: "task1"},
		{UserID: userID, Text: "task2"},
		{UserID: userID, Text: "task3"},
	})
	require.NoError(t, err)

	// when
	output, err := cmd.Execute(ctx, input)

	// then
	require.NoError(t, err)
	require.NotNil(t, output)
	assert.Len(t, output.Todos, 3)
	assert.Equal(t, "task1", output.Todos[0].Text)
	assert.Equal(t, "task2", output.Todos[1].Text)
	assert.Equal(t, "task3", output.Todos[2].Text)

	// DB にも永続化されていることを確認
	repo := gateway.NewTodoRepository(dbc.DB)
	todos, err := repo.FindTodos(ctx, userID)
	require.NoError(t, err)
	assert.Len(t, todos, 3)
}

func Test_CreateBulkTodosCommand_Execute_shouldRollbackAll_whenMiddleTodoFails(t *testing.T) {
	t.Parallel()
	ctx := context.Background()
	userID := rand.Intn(1000000) //nolint:gosec

	// given
	cleanupTodoTable(t, userID)
	txManager := gateway.NewTodoCreateBulkCommandTxManager(dbc)
	cmd := usecase.NewCreateBulkTodosCommand(txManager)

	// コンストラクタをバイパスし、varchar(255) を超えるテキストで DB 制約違反を起こす
	input := &domain.CreateBulkTodosInput{
		UserID: userID,
		Todos: []domain.CreateTodoInput{
			{UserID: userID, Text: "task1"},
			{UserID: userID, Text: strings.Repeat("a", 256)},
			{UserID: userID, Text: "task3"},
		},
	}

	// when
	output, err := cmd.Execute(ctx, input)

	// then
	require.Error(t, err)
	assert.Nil(t, output)

	// 1件目もロールバックされ、DB にレコードが残っていないことを確認
	repo := gateway.NewTodoRepository(dbc.DB)
	todos, err := repo.FindTodos(ctx, userID)
	require.NoError(t, err)
	assert.Empty(t, todos)
}
