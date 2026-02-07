package usecase

import (
	"context"
	"fmt"
	"log/slog"

	"github.com/mocoarow/todo-apps/backend-gin-gorm/domain"
)

type TodoUsecase struct {
	findTodosQuery         *FindTodosQuery
	createTodoCommand      *CreateTodoCommand
	createBulkTodosCommand *CreateBulkTodosCommand
	updateTodoCommand      *UpdateTodoCommand
	deleteTodoCommand      *DeleteTodoCommand
	logger                 *slog.Logger
}

func NewTodoUsecase(repo domain.TodoRepository, createBulkCommandTxManager TodoCreateBulkCommandTxManager) *TodoUsecase {
	findTodosQuery := NewFindTodosQuery(repo)
	createTodoCommand := NewCreateTodoCommand(repo)
	createBulkTodosCommand := NewCreateBulkTodosCommand(createBulkCommandTxManager)
	updateTodoCommand := NewUpdateTodoCommand(repo)
	deleteTodoCommand := NewDeleteTodoCommand(repo)
	return &TodoUsecase{
		findTodosQuery:         findTodosQuery,
		createTodoCommand:      createTodoCommand,
		createBulkTodosCommand: createBulkTodosCommand,
		updateTodoCommand:      updateTodoCommand,
		deleteTodoCommand:      deleteTodoCommand,
		logger:                 slog.Default().With(slog.String(domain.LoggerNameKey, domain.AppName+"-TodoUsecase")),
	}
}

func (u *TodoUsecase) FindTodos(ctx context.Context, userID int) ([]domain.Todo, error) {
	ctx, span := tracer.Start(ctx, "FindTodos")
	defer span.End()
	u.logger.InfoContext(ctx, "FindTodos called")

	todos, err := u.findTodosQuery.Execute(ctx, userID)
	if err != nil {
		return nil, fmt.Errorf("execute find todos query: %w", err)
	}
	return todos, nil
}

func (u *TodoUsecase) CreateTodo(ctx context.Context, input *domain.CreateTodoInput) (*domain.CreateTodoOutput, error) {
	output, err := u.createTodoCommand.Execute(ctx, input)
	if err != nil {
		return nil, fmt.Errorf("execute create todo command: %w", err)
	}
	return output, nil
}

func (u *TodoUsecase) CreateBulkTodos(ctx context.Context, input *domain.CreateBulkTodosInput) (*domain.CreateBulkTodosOutput, error) {
	output, err := u.createBulkTodosCommand.Execute(ctx, input)
	if err != nil {
		return nil, fmt.Errorf("execute create bulk todos command: %w", err)
	}
	return output, nil
}

func (u *TodoUsecase) UpdateTodo(ctx context.Context, input *domain.UpdateTodoInput) (*domain.UpdateTodoOutput, error) {
	output, err := u.updateTodoCommand.Execute(ctx, input)
	if err != nil {
		return nil, fmt.Errorf("execute update todo command: %w", err)
	}
	return output, nil
}

func (u *TodoUsecase) DeleteTodo(ctx context.Context, input *domain.DeleteTodoInput) error {
	err := u.deleteTodoCommand.Execute(ctx, input)
	if err != nil {
		return fmt.Errorf("execute delete todo command: %w", err)
	}
	return nil
}
