package domain

import (
	"fmt"
	"time"
)

// ErrTodoNotFound is returned when a requested todo does not exist.
var ErrTodoNotFound = fmt.Errorf("todo not found")

// Todo represents a single todo item belonging to a user.
type Todo struct {
	ID         int    `validate:"required,gt=0"`
	UserID     int    `validate:"required,gt=0"`
	Text       string `validate:"required,max=255"`
	IsComplete bool
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

// NewTodo creates a validated Todo. Returns an error if validation fails.
func NewTodo(id int, userID int, text string, isComplete bool, createdAt, updatedAt time.Time) (*Todo, error) {
	m := &Todo{
		ID:         id,
		UserID:     userID,
		Text:       text,
		IsComplete: isComplete,
		CreatedAt:  createdAt,
		UpdatedAt:  updatedAt,
	}

	if err := ValidateStruct(m); err != nil {
		return nil, fmt.Errorf("validate todo model: %w", err)
	}

	return m, nil
}

// CreateTodoInput holds the parameters required to create a single todo.
type CreateTodoInput struct {
	UserID int    `validate:"required,gt=0"`
	Text   string `validate:"required,max=255"`
}

// NewCreateTodoInput creates a validated CreateTodoInput. Returns an error if validation fails.
func NewCreateTodoInput(userID int, text string) (*CreateTodoInput, error) {
	m := &CreateTodoInput{
		UserID: userID,
		Text:   text,
	}
	if err := ValidateStruct(m); err != nil {
		return nil, fmt.Errorf("validate create todo input: %w", err)
	}

	return m, nil
}

// CreateTodoOutput holds the result of a single todo creation.
type CreateTodoOutput struct {
	Todo *Todo `validate:"required"`
}

// NewCreateTodoOutput creates a validated CreateTodoOutput. Returns an error if validation fails.
func NewCreateTodoOutput(todo *Todo) (*CreateTodoOutput, error) {
	m := &CreateTodoOutput{
		Todo: todo,
	}
	if err := ValidateStruct(m); err != nil {
		return nil, fmt.Errorf("validate create todo output: %w", err)
	}
	return m, nil
}

// CreateBulkTodosInput holds the parameters required to create multiple todos at once (1-100 items).
type CreateBulkTodosInput struct {
	UserID int                `validate:"required,gt=0"`
	Todos  []*CreateTodoInput `validate:"required,min=1,max=100,dive"`
}

// NewCreateBulkTodosInput creates a validated CreateBulkTodosInput. Returns an error if validation fails.
func NewCreateBulkTodosInput(userID int, todos []*CreateTodoInput) (*CreateBulkTodosInput, error) {
	m := &CreateBulkTodosInput{
		UserID: userID,
		Todos:  todos,
	}
	if err := ValidateStruct(m); err != nil {
		return nil, fmt.Errorf("validate create bulk todos input: %w", err)
	}
	return m, nil
}

// CreateBulkTodosOutput holds the result of a bulk todo creation.
type CreateBulkTodosOutput struct {
	Todos []*Todo `validate:"required,min=1,dive"`
}

// NewCreateBulkTodosOutput creates a validated CreateBulkTodosOutput. Returns an error if validation fails.
func NewCreateBulkTodosOutput(todos []*Todo) (*CreateBulkTodosOutput, error) {
	m := &CreateBulkTodosOutput{
		Todos: todos,
	}
	if err := ValidateStruct(m); err != nil {
		return nil, fmt.Errorf("validate create bulk todos output: %w", err)
	}
	return m, nil
}

// UpdateTodoInput holds the parameters required to update an existing todo.
type UpdateTodoInput struct {
	ID         int    `validate:"required,gt=0"`
	UserID     int    `validate:"required,gt=0"`
	Text       string `validate:"required,max=255"`
	IsComplete bool
}

// NewUpdateTodoInput creates a validated UpdateTodoInput. Returns an error if validation fails.
func NewUpdateTodoInput(id int, userID int, text string, isComplete bool) (*UpdateTodoInput, error) {
	m := &UpdateTodoInput{
		ID:         id,
		UserID:     userID,
		Text:       text,
		IsComplete: isComplete,
	}
	if err := ValidateStruct(m); err != nil {
		return nil, fmt.Errorf("validate update todo input: %w", err)
	}
	return m, nil
}

// UpdateTodoOutput holds the result of a todo update.
type UpdateTodoOutput struct {
	Todo *Todo `validate:"required"`
}

// NewUpdateTodoOutput creates a validated UpdateTodoOutput. Returns an error if validation fails.
func NewUpdateTodoOutput(todo *Todo) (*UpdateTodoOutput, error) {
	m := &UpdateTodoOutput{
		Todo: todo,
	}
	if err := ValidateStruct(m); err != nil {
		return nil, fmt.Errorf("validate update todo output: %w", err)
	}
	return m, nil
}

// DeleteTodoInput holds the parameters required to delete a todo.
type DeleteTodoInput struct {
	ID     int `validate:"required,gt=0"`
	UserID int `validate:"required,gt=0"`
}

// NewDeleteTodoInput creates a validated DeleteTodoInput. Returns an error if validation fails.
func NewDeleteTodoInput(id int, userID int) (*DeleteTodoInput, error) {
	m := &DeleteTodoInput{
		ID:     id,
		UserID: userID,
	}
	if err := ValidateStruct(m); err != nil {
		return nil, fmt.Errorf("validate delete todo input: %w", err)
	}
	return m, nil
}
