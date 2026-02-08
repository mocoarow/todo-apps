package usecase

import (
	"context"
	"fmt"

	"github.com/mocoarow/todo-apps/backend-gin-gorm/domain"
)

// TodoUpdater defines the interface for updating todos in the repository.
type TodoUpdater interface {
	UpdateTodo(ctx context.Context, input *domain.UpdateTodoInput) (*domain.Todo, error)
}

// UpdateTodoCommand updates an existing todo item in the repository.
type UpdateTodoCommand struct {
	repo TodoUpdater
}

// NewUpdateTodoCommand returns a new UpdateTodoCommand.
func NewUpdateTodoCommand(repo TodoUpdater) *UpdateTodoCommand {
	return &UpdateTodoCommand{
		repo: repo,
	}
}

// Execute updates the todo and returns the updated result.
func (u *UpdateTodoCommand) Execute(ctx context.Context, input *domain.UpdateTodoInput) (*domain.UpdateTodoOutput, error) {
	todo, err := u.repo.UpdateTodo(ctx, input)
	if err != nil {
		return nil, fmt.Errorf("update todo: %w", err)
	}

	output, err := domain.NewUpdateTodoOutput(todo)
	if err != nil {
		return nil, fmt.Errorf("create updated todo output: %w", err)
	}

	return output, nil
}
