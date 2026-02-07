package usecase

import (
	"context"
	"fmt"

	"github.com/mocoarow/todo-apps/backend-gin-gorm/domain"
)

type FindTodosQuery struct {
	repo domain.TodoRepository
}

func NewFindTodosQuery(repo domain.TodoRepository) *FindTodosQuery {
	return &FindTodosQuery{
		repo: repo,
	}
}

func (q *FindTodosQuery) Execute(ctx context.Context, userID int) ([]domain.Todo, error) {
	todos, err := q.repo.FindTodos(ctx, userID)
	if err != nil {
		return nil, fmt.Errorf("find todos: %w", err)
	}
	return todos, nil
}
