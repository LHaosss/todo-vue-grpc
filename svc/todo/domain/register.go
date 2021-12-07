// 定义store实现的方法的interface
package domain

import (
	"context"
)

type TodoRegisterIntf interface {
	GetTodos(context.Context) ([]TodoIntf, error) 
	AddTodo(context.Context, TodoIntf) error
	DeleteTodo(context.Context, string) error
}
