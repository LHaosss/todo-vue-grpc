package todo

import (
	"context"

	todo "github.com/LHaosss/todo-vue-grpc/proto/gen/go/todo/v1"
)


func (t *TodoServiceHandler) DeleteTodo(context.Context, *todo.DeleteTodoRequest, *todo.DeleteTodoResponse) error {
	return nil
}