package store

type TodoStore struct {
	TodoID string
	Index string
	Task string
}

func (t *TodoStore) GetTodoID() string {
	if t.TodoID == "" {
		return ""
	}
	return t.TodoID
}

func (t *TodoStore) GetIndex() string {
	if t.Index == "" {
		return ""
	}
	return t.Index
}

func (t *TodoStore) GetTask() string {
	if t.Task == "" {
		return ""
	}
	return t.Task
}