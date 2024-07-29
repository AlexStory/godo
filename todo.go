package main

import "fmt"

type Todo struct {
	Id    int    `json:"id"`
	Title string `json:"title"`
	Done  bool   `json:"done"`
}

type Todos struct {
	todos []Todo
}

func NewTodos() Todos {
	return Todos{
		todos: []Todo{},
	}
}

func (t *Todos) Add(title string) {
	newTodo := Todo{
		Id:    t.nextId(),
		Title: title,
		Done:  false,
	}
	t.todos = append(t.todos, newTodo)
	Write(*t)
}

func (t *Todos) Toggle(id int) error {
	for i, todo := range t.todos {
		if todo.Id == id {
			t.todos[i].Done = !todo.Done
			Write(*t)
			return nil
		}
	}
	return fmt.Errorf("Todo not found")
}

func (t *Todos) Update(id int, title string) error {
	for i, todo := range t.todos {
		if todo.Id == id {
			t.todos[i].Title = title
			Write(*t)
			return nil
		}
	}
	return fmt.Errorf("Todo not found")
}

func (t *Todos) nextId() int {
	maxId := 0
	for _, todo := range t.todos {
		if todo.Id > maxId {
			maxId = todo.Id
		}
	}
	return maxId + 1
}

func (t *Todo) String() string {
	if t.Done {
		return fmt.Sprintf("%d: [x] %s", t.Id, t.Title)
	}
	return fmt.Sprintf("%d: [ ] %s", t.Id, t.Title)
}
