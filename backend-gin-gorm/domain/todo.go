package domain

import (
	"fmt"
	"time"
)

var ErrTodoNotFound = fmt.Errorf("todo not found")

type Todo struct {
	ID         int    `validate:"required,gt=0"`
	UserID     int    `validate:"required,gt=0"`
	Text       string `validate:"required,max=255"`
	IsComplete bool
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

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

type CreateTodoInput struct {
	UserID int    `validate:"required,gt=0"`
	Text   string `validate:"required,max=255"`
}

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

type CreateTodoOutput struct {
	Todo *Todo `validate:"required"`
}

func NewCreateTodoOutput(todo *Todo) (*CreateTodoOutput, error) {
	m := &CreateTodoOutput{
		Todo: todo,
	}
	if err := ValidateStruct(m); err != nil {
		return nil, fmt.Errorf("validate create todo output: %w", err)
	}
	return m, nil
}

type CreateBulkTodosInput struct {
	UserID int                `validate:"required,gt=0"`
	Todos  []*CreateTodoInput `validate:"required,min=1,max=100,dive"`
}

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

type CreateBulkTodosOutput struct {
	Todos []*Todo `validate:"required,min=1,dive"`
}

func NewCreateBulkTodosOutput(todos []*Todo) (*CreateBulkTodosOutput, error) {
	m := &CreateBulkTodosOutput{
		Todos: todos,
	}
	if err := ValidateStruct(m); err != nil {
		return nil, fmt.Errorf("validate create bulk todos output: %w", err)
	}
	return m, nil
}

type UpdateTodoInput struct {
	ID         int    `validate:"required,gt=0"`
	UserID     int    `validate:"required,gt=0"`
	Text       string `validate:"required,max=255"`
	IsComplete bool
}

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

type UpdateTodoOutput struct {
	Todo *Todo `validate:"required"`
}

func NewUpdateTodoOutput(todo *Todo) (*UpdateTodoOutput, error) {
	m := &UpdateTodoOutput{
		Todo: todo,
	}
	if err := ValidateStruct(m); err != nil {
		return nil, fmt.Errorf("validate update todo output: %w", err)
	}
	return m, nil
}

type DeleteTodoInput struct {
	ID     int `validate:"required,gt=0"`
	UserID int `validate:"required,gt=0"`
}

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
