package todo

import (
	"encoding/json"
	"errors"
	"os"
	"time"
)

type item struct {
	Task      string
	Completed bool
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Todos []item

func (list *Todos) AddTodo(task string) {
	todo := item{
		Task:      task,
		Completed: false,
		CreatedAt: time.Now(),
		UpdatedAt: time.Time{},
	}

	*list = append(*list, todo)
}

func (list *Todos) MarkComplete(index int) error {
	ls := *list
	if index <= 0 || index > len(ls) {
		return errors.New("invalid choice")
	}

	ls[index-1].Completed = true
	ls[index-1].UpdatedAt = time.Now()

	return nil
}

func (list *Todos) DeleteTodo(index int) error {
	ls := *list
	if index <= 0 || index > len(ls) {
		return errors.New("todo doesn't exist")
	}

	*list = append(ls[:index-1], ls[index:]...)

	return nil
}

func (list *Todos) Load(filename string) error {
	file, err := os.ReadFile(filename)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return nil
		}
		return err
	}
	if len(file) == 0 {
		return err
	}

	err = json.Unmarshal(file, list)
	if err != nil {
		return err
	}
	return nil
}

func (list *Todos) StoreTodo(filename string) error {
	data, err := json.Marshal(list)
	if err != nil {
		return err
	}

	return os.WriteFile(filename, data, 0644)
}

func (t *Todos) Print() {

}
