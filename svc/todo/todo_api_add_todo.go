package todo

import (
	"context"

	todo "github.com/LHaosss/todo-vue-grpc/proto/gen/go/todo/v1"
)

func (t *TodoServiceHandler) AddTodo(context.Context, *todo.AddTodoRequest, *todo.AddTodoResponse) error {
	return nil
}