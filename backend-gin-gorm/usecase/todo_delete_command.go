package usecase

import (
	"context"
	"fmt"

	"github.com/mocoarow/todo-apps/backend-gin-gorm/domain"
)

// TodoDeleter defines the interface for deleting todos from the repository.
type TodoDeleter interface {
	DeleteTodo(ctx context.Context, input *domain.DeleteTodoInput) error
}

// DeleteTodoCommand removes a todo item from the repository.
type DeleteTodoCommand struct {
	repo TodoDeleter
}

// NewDeleteTodoCommand returns a new DeleteTodoCommand.
func NewDeleteTodoCommand(repo TodoDeleter) *DeleteTodoCommand {
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
