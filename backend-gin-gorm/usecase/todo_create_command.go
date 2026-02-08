package usecase

import (
	"context"
	"fmt"

	"github.com/mocoarow/todo-apps/backend-gin-gorm/domain"
)

// TodoCreator defines the interface for creating todos in the repository.
type TodoCreator interface {
	CreateTodo(ctx context.Context, input *domain.CreateTodoInput) (*domain.Todo, error)
}

// CreateTodoCommand persists a new todo item via the repository.
type CreateTodoCommand struct {
	repo TodoCreator
}

// NewCreateTodoCommand returns a new CreateTodoCommand.
func NewCreateTodoCommand(repo TodoCreator) *CreateTodoCommand {
	return &CreateTodoCommand{
		repo: repo,
	}
}

// Execute creates a todo and returns the result.
func (u *CreateTodoCommand) Execute(ctx context.Context, input *domain.CreateTodoInput) (*domain.CreateTodoOutput, error) {
	todo, err := u.repo.CreateTodo(ctx, input)
	if err != nil {
		return nil, fmt.Errorf("create todo: %w", err)
	}

	output, err := domain.NewCreateTodoOutput(todo)
	if err != nil {
		return nil, fmt.Errorf("create todo output: %w", err)
	}

	return output, nil
}
