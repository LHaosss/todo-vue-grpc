package todo

import "github.com/LHaosss/todo-vue-grpc/svc/todo/domain"

/* 定义TodoServiceHandler
包含实现service所需要的基础组件，例如store、redis、token等，按需引入
*/

type TodoServiceHandler struct {
	// 数据库
	TodoRegister domain.TodoRegisterIntf
}

// new方法生成TodoServiceHandler
func NewTodoServiceHandler(r domain.TodoRegisterIntf) *TodoServiceHandler {
	return &TodoServiceHandler{
		TodoRegister: r,
	}
}
