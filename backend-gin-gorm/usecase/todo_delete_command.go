package usecase

import (
	"context"
	"fmt"

	"github.com/mocoarow/todo-apps/backend-gin-gorm/domain"
)

// DeleteTodoCommand removes a todo item from the repository.
type DeleteTodoCommand struct {
	repo domain.TodoRepository
}

// NewDeleteTodoCommand returns a new DeleteTodoCommand.
func NewDeleteTodoCommand(repo domain.TodoRepository) *DeleteTodoCommand {
	return &DeleteTodoCommand{
		repo: repo,
	}
}

// Execute deletes the specified todo item.
func (u *DeleteTodoCommand) Execute(ctx context.Context, input *domain.DeleteTodoInput) error {
	err := u.repo.DeleteTodo(ctx, input)
	if err != nil {
		return fmt.Errorf("delete todo: %w", err)
	}

	return nil
}
