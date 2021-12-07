package todo

import (
	"context"

	todo "github.com/LHaosss/todo-vue-grpc/proto/gen/go/todo/v1"
)

func (t *TodoServiceHandler) GetTodos(context.Context, *todo.GetTodosRequest, *todo.GetTodosResponse) error {
	return nil
}