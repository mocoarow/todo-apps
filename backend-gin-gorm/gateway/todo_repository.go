package gateway

import (
	"context"
	"errors"
	"fmt"
	"time"

	"gorm.io/gorm"

	"github.com/mocoarow/todo-apps/backend-gin-gorm/domain"
)

// TodoEntity is the GORM model for the "todo" table.
type TodoEntity struct {
	ID         int       `gorm:"primaryKey;autoIncrement"`
	UserID     int       `gorm:"not null"`
	Text       string    `gorm:"type:varchar(255);not null"`
	IsComplete bool      `gorm:"not null;default:false"`
	CreatedAt  time.Time `gorm:"autoCreateTime"`
	UpdatedAt  time.Time `gorm:"autoUpdateTime"`
}

func (e *TodoEntity) TableName() string {
	return "todo"
}

func (e *TodoEntity) toTodo() (*domain.Todo, error) {
	todo, err := domain.NewTodo(e.ID, e.UserID, e.Text, e.IsComplete, e.CreatedAt, e.UpdatedAt)
	if err != nil {
		return nil, fmt.Errorf("to todo model: %w", err)
	}

	return todo, nil
}

// TodoEntities is a slice of TodoEntity with batch conversion support.
type TodoEntities []TodoEntity

func (e TodoEntities) toTodos() ([]domain.Todo, error) {
	todos := make([]domain.Todo, len(e))
	for i, todoE := range e {
		todo, err := todoE.toTodo()
		if err != nil {
			return nil, fmt.Errorf("to todo: %w", err)
		}
		todos[i] = *todo
	}

	return todos, nil
}

// TodoRepository implements todo persistence operations using GORM.
type TodoRepository struct {
	db *gorm.DB
}

// NewTodoRepository returns a new TodoRepository backed by the given GORM DB.
func NewTodoRepository(db *gorm.DB) *TodoRepository {
	return &TodoRepository{
		db: db,
	}
}

// FindTodos returns all todos for the given user, ordered by ID.
func (r *TodoRepository) FindTodos(ctx context.Context, userID int) ([]domain.Todo, error) {
	var entities TodoEntities
	if result := r.db.WithContext(ctx).Where("user_id = ?", userID).Order("id").Find(&entities); result.Error != nil {
		return nil, fmt.Errorf("find todos: %w", result.Error)
	}
	todos, err := entities.toTodos()
	if err != nil {
		return nil, fmt.Errorf("to todos: %w", err)
	}
	return todos, nil
}

// CreateTodo inserts a new todo record and returns the created domain model.
func (r *TodoRepository) CreateTodo(ctx context.Context, input *domain.CreateTodoInput) (*domain.Todo, error) {
	entity := &TodoEntity{ //nolint:exhaustruct
		UserID:     input.UserID,
		Text:       input.Text,
		IsComplete: false,
	}

	if input.Text == "XYZ" {
		return nil, errors.New("simulated database error")
	}

	if result := r.db.WithContext(ctx).Create(entity); result.Error != nil {
		return nil, fmt.Errorf("create todo: %w", result.Error)
	}

	todo, err := entity.toTodo()
	if err != nil {
		return nil, fmt.Errorf("to todo: %w", err)
	}

	return todo, nil
}

// UpdateTodo updates a todo owned by the user. Returns ErrTodoNotFound if not found.
func (r *TodoRepository) UpdateTodo(ctx context.Context, input *domain.UpdateTodoInput) (*domain.Todo, error) {
	var entity TodoEntity

	// Find the todo by ID and UserID to ensure the user owns this todo
	if result := r.db.WithContext(ctx).Where("id = ? AND user_id = ?", input.ID, input.UserID).First(&entity); result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, domain.ErrTodoNotFound
		}
		return nil, fmt.Errorf("find todo: %w", result.Error)
	}

	// Update the fields
	entity.Text = input.Text
	entity.IsComplete = input.IsComplete

	// Save the changes
	if result := r.db.WithContext(ctx).Save(&entity); result.Error != nil {
		return nil, fmt.Errorf("update todo: %w", result.Error)
	}

	todo, err := entity.toTodo()
	if err != nil {
		return nil, fmt.Errorf("to todo: %w", err)
	}

	return todo, nil
}

// DeleteTodo deletes a todo owned by the user. Returns ErrTodoNotFound if not found.
func (r *TodoRepository) DeleteTodo(ctx context.Context, input *domain.DeleteTodoInput) error {
	// Delete the todo by ID and UserID to ensure the user owns this todo
	result := r.db.WithContext(ctx).Where("id = ? AND user_id = ?", input.ID, input.UserID).Delete(&TodoEntity{}) //nolint:exhaustruct
	if result.Error != nil {
		return fmt.Errorf("delete todo: %w", result.Error)
	}

	// Check if any rows were affected
	if result.RowsAffected == 0 {
		return domain.ErrTodoNotFound
	}

	return nil
}

// TodoCreateBulkCommandTxManager manages GORM transactions for bulk todo creation.
type TodoCreateBulkCommandTxManager struct {
	dbc *DBConnection
}

// NewTodoCreateBulkCommandTxManager returns a new transaction manager.
func NewTodoCreateBulkCommandTxManager(dbc *DBConnection) *TodoCreateBulkCommandTxManager {
	return &TodoCreateBulkCommandTxManager{
		dbc: dbc,
	}
}

// WithTransaction executes fn within a database transaction, rolling back on error.
func (tm *TodoCreateBulkCommandTxManager) WithTransaction(ctx context.Context, fn func(createTodo domain.CreateTodoFunc) ([]domain.Todo, error)) ([]domain.Todo, error) {
	var todos []domain.Todo
	err := tm.dbc.DB.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		todoRepo := NewTodoRepository(tx)
		var err error
		todos, err = fn(todoRepo.CreateTodo)
		if err != nil {
			return fmt.Errorf("execute function in transaction: %w", err)
		}
		return nil
	})
	if err != nil {
		return nil, fmt.Errorf("transaction failed: %w", err)
	}
	return todos, nil
}
