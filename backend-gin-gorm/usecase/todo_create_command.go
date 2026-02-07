package usecase

import (
	"context"
	"fmt"

	"github.com/mocoarow/todo-apps/backend-gin-gorm/domain"
)

// CreateTodoCommand persists a new todo item via the repository.
type CreateTodoCommand struct {
	repo domain.TodoRepository
}

// NewCreateTodoCommand returns a new CreateTodoCommand.
func NewCreateTodoCommand(repo domain.TodoRepository) *CreateTodoCommand {
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
