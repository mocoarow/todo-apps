package usecase

import (
	"context"
	"fmt"

	"github.com/mocoarow/todo-apps/backend-gin-gorm/domain"
)

type TodoCreateBulkCommandTxManager interface {
	WithTransaction(ctx context.Context, fn func(todoRepo domain.TodoRepository) ([]domain.Todo, error)) ([]domain.Todo, error)
}

type CreateBulkTodosCommand struct {
	txManager TodoCreateBulkCommandTxManager
}

func NewCreateBulkTodosCommand(txManager TodoCreateBulkCommandTxManager) *CreateBulkTodosCommand {
	return &CreateBulkTodosCommand{
		txManager: txManager,
	}
}

func (u *CreateBulkTodosCommand) Execute(ctx context.Context, input *domain.CreateBulkTodosInput) (*domain.CreateBulkTodosOutput, error) {
	todos, err := u.txManager.WithTransaction(ctx, func(todoRepo domain.TodoRepository) ([]domain.Todo, error) {
		var todos []domain.Todo
		for _, todoInput := range input.Todos {
			todo, err := todoRepo.CreateTodo(ctx, &todoInput)
			if err != nil {
				return nil, fmt.Errorf("create todo: %w", err)
			}
			if todo == nil {
				return nil, fmt.Errorf("created todo is nil")
			}

			todos = append(todos, *todo)
		}
		return todos, nil
	})
	if err != nil {
		return nil, fmt.Errorf("do in transaction: %w", err)
	}

	output, err := domain.NewCreateBulkTodosOutput(todos)
	if err != nil {
		return nil, fmt.Errorf("create bulk todos output: %w", err)
	}
	return output, nil
}
