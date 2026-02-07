package usecase

import (
	"context"
	"fmt"

	"github.com/mocoarow/todo-apps/backend-gin-gorm/domain"
)

type UpdateTodoCommand struct {
	repo domain.TodoRepository
}

func NewUpdateTodoCommand(repo domain.TodoRepository) *UpdateTodoCommand {
	return &UpdateTodoCommand{
		repo: repo,
	}
}

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
