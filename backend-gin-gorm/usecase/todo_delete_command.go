package usecase

import (
	"context"
	"fmt"

	"github.com/mocoarow/todo-apps/backend-gin-gorm/domain"
)

type DeleteTodoCommand struct {
	repo domain.TodoRepository
}

func NewDeleteTodoCommand(repo domain.TodoRepository) *DeleteTodoCommand {
	return &DeleteTodoCommand{
		repo: repo,
	}
}

func (u *DeleteTodoCommand) Execute(ctx context.Context, input *domain.DeleteTodoInput) error {
	err := u.repo.DeleteTodo(ctx, input)
	if err != nil {
		return fmt.Errorf("delete todo: %w", err)
	}

	return nil
}
