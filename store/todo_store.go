package store

import (
	"context"

	"github.com/LHaosss/todo-vue-grpc/svc/todo/domain"
	"gorm.io/gorm"
)

type TodoRegister struct {
	db gorm.DB
}

var _ (domain.TodoRegisterIntf) = (*TodoRegister)(nil)

func NewTodoRegister(db gorm.DB) *TodoRegister {
	return &TodoRegister{
		db: db,
	}
}

func (t *TodoRegister) GetTodos(context.Context) ([]domain.TodoIntf, error) {
	return nil, nil
}

func (t *TodoRegister) AddTodo(context.Context, domain.TodoIntf) error {
	return nil
}

func (t *TodoRegister) DeleteTodo(context.Context, string) error {
	return nil
}

