// 定义domain的struct和interface
package domain


type Todo struct {
	Index string
	Task string
}

type TodoMapper interface {
	toDomainInterface
	fromDomainInterface
}

type toDomainInterface interface {
	ToDomainTodo() TodoIntf
}

type fromDomainInterface interface {
	FromDomainTodo(TodoIntf)
}

type TodoIntf interface {
	GetIndex() string
	GetTask() string
}

func (t *Todo) GetIndex() string {
	if t.Index == "" {
		return ""
	}
	return t.Index
}

func (t *Todo) GetTask() string {
	if t.Task == "" {
		return ""
	}
	return t.Task
}
