package usecase

import (
	"context"
	"fmt"
	"log/slog"

	"github.com/mocoarow/todo-apps/backend-gin-gorm/domain"
)

// TodoRepository composes all todo CRUD interfaces.
type TodoRepository interface {
	TodoCreator
	TodoFinder
	TodoUpdater
	TodoDeleter
}

// TodoUsecase orchestrates todo CRUD operations via command/query objects.
type TodoUsecase struct {
	findTodosQuery         *FindTodosQuery
	createTodoCommand      *CreateTodoCommand
	createBulkTodosCommand *CreateBulkTodosCommand
	updateTodoCommand      *UpdateTodoCommand
	deleteTodoCommand      *DeleteTodoCommand
	logger                 *slog.Logger
}

// NewTodoUsecase returns a new TodoUsecase wired with the given repository and transaction manager.
func NewTodoUsecase(repo TodoRepository, createBulkCommandTxManager TodoCreateBulkCommandTxManager) *TodoUsecase {
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

// FindTodos returns all todos belonging to the given user.
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

// CreateTodo creates a single todo item.
func (u *TodoUsecase) CreateTodo(ctx context.Context, input *domain.CreateTodoInput) (*domain.CreateTodoOutput, error) {
	output, err := u.createTodoCommand.Execute(ctx, input)
	if err != nil {
		return nil, fmt.Errorf("execute create todo command: %w", err)
	}
	return output, nil
}

// CreateBulkTodos creates multiple todo items within a single transaction.
func (u *TodoUsecase) CreateBulkTodos(ctx context.Context, input *domain.CreateBulkTodosInput) (*domain.CreateBulkTodosOutput, error) {
	output, err := u.createBulkTodosCommand.Execute(ctx, input)
	if err != nil {
		return nil, fmt.Errorf("execute create bulk todos command: %w", err)
	}
	return output, nil
}

// UpdateTodo updates an existing todo item.
func (u *TodoUsecase) UpdateTodo(ctx context.Context, input *domain.UpdateTodoInput) (*domain.UpdateTodoOutput, error) {
	output, err := u.updateTodoCommand.Execute(ctx, input)
	if err != nil {
		return nil, fmt.Errorf("execute update todo command: %w", err)
	}
	return output, nil
}

// DeleteTodo removes a todo item.
func (u *TodoUsecase) DeleteTodo(ctx context.Context, input *domain.DeleteTodoInput) error {
	if err := u.deleteTodoCommand.Execute(ctx, input); err != nil {
		return fmt.Errorf("execute delete todo command: %w", err)
	}
	return nil
}
