package usecase

import (
	"context"
	"fmt"

	"github.com/mocoarow/todo-apps/backend-gin-gorm/domain"
)

// TodoFinder defines the interface for fetching todos from the repository.
type TodoFinder interface {
	FindTodos(ctx context.Context, userID int) ([]domain.Todo, error)
}

// FindTodosQuery fetches todos for a specific user from the repository.
type FindTodosQuery struct {
	repo TodoFinder
}

// NewFindTodosQuery returns a new FindTodosQuery.
func NewFindTodosQuery(repo TodoFinder) *FindTodosQuery {
	return &FindTodosQuery{
		repo: repo,
	}
}

// Execute retrieves all todos for the given userID.
func (q *FindTodosQuery) Execute(ctx context.Context, userID int) ([]domain.Todo, error) {
	todos, err := q.repo.FindTodos(ctx, userID)
	if err != nil {
		return nil, fmt.Errorf("find todos: %w", err)
	}
	return todos, nil
}
