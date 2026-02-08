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

func Test_CreateTodoCommand_Execute_shouldReturnCreatedTodo_whenValidInput(t *testing.T) {
	t.Parallel()
	ctx := context.Background()
	userID := rand.Intn(1000000) //nolint:gosec

	// given
	cleanupTodoTable(t, userID)
	repo := gateway.NewTodoRepository(dbc.DB)
	cmd := usecase.NewCreateTodoCommand(repo)

	input, err := domain.NewCreateTodoInput(userID, "buy milk")
	require.NoError(t, err)

	// when
	output, err := cmd.Execute(ctx, input)

	// then
	require.NoError(t, err)
	require.NotNil(t, output)
	assert.Equal(t, "buy milk", output.Todo.Text)
	assert.Equal(t, userID, output.Todo.UserID)
	assert.False(t, output.Todo.IsComplete)
	assert.Positive(t, output.Todo.ID)

	// DB にも永続化されていることを確認
	todos, err := repo.FindTodos(ctx, userID)
	require.NoError(t, err)
	assert.Len(t, todos, 1)
	assert.Equal(t, "buy milk", todos[0].Text)
}

func Test_CreateTodoCommand_Execute_shouldReturnError_whenRepoFails(t *testing.T) {
	t.Parallel()
	ctx := context.Background()
	userID := rand.Intn(1000000) //nolint:gosec

	// given
	cleanupTodoTable(t, userID)
	repo := gateway.NewTodoRepository(dbc.DB)
	cmd := usecase.NewCreateTodoCommand(repo)

	// コンストラクタをバイパスし、varchar(255) を超えるテキストで DB 制約違反を起こす
	input := &domain.CreateTodoInput{
		UserID: userID,
		Text:   strings.Repeat("a", 256),
	}

	// when
	output, err := cmd.Execute(ctx, input)

	// then
	require.Error(t, err)
	assert.Nil(t, output)
	assert.Contains(t, err.Error(), "create todo")

	// DB にレコードが残っていないことを確認
	todos, err := repo.FindTodos(ctx, userID)
	require.NoError(t, err)
	assert.Empty(t, todos)
}
