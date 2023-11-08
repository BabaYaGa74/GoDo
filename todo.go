package todo

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/alexeyco/simpletable"
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

func (list *Todos) MarkAll() {
	ls := *list
	len := len(ls)
	for i := 0; i < len; i++ {
		ls[i].Completed = false
		ls[i].UpdatedAt = time.Now()
		ls[i].CreatedAt = time.Now()
	}
}

func (list *Todos) DeleteTodo(index int) error {
	ls := *list
	if index <= 0 || index > len(ls) {
		return errors.New("todo doesn't exist")
	}

	*list = append(ls[:index-1], ls[index:]...)

	return nil
}

func (list *Todos) DeleteAll() {
	*list = make(Todos, 0)
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
	println()
	println()

	table := simpletable.New()

	table.Header = &simpletable.Header{
		Cells: []*simpletable.Cell{
			{Align: simpletable.AlignCenter, Text: "#"},
			{Align: simpletable.AlignCenter, Text: "Task"},
			{Align: simpletable.AlignCenter, Text: "Done?"},
			{Align: simpletable.AlignRight, Text: "CreatedAt"},
			{Align: simpletable.AlignRight, Text: "CompletedAt"},
		},
	}

	var cells [][]*simpletable.Cell
	for idx, item := range *t {
		idx++
		task := blue(item.Task)
		done := blue("no")
		if item.Completed {
			task = green(fmt.Sprintf("\u2705 %s", item.Task))
			done = green("yes")
		}
		cells = append(cells, []*simpletable.Cell{
			{Text: fmt.Sprintf("%d", idx)},
			{Text: task},
			{Text: done},
			{Text: item.CreatedAt.Format(time.RFC822)},
			{Text: item.UpdatedAt.Format(time.RFC822)},
		})
	}

	table.Body = &simpletable.Body{Cells: cells}

	table.Footer = &simpletable.Footer{Cells: []*simpletable.Cell{
		{Align: simpletable.AlignCenter, Span: 5, Text: t.CountPending()},
	}}

	table.SetStyle(simpletable.StyleUnicode)

	table.Println()
	println(red(t.timeLeft()))
	println()
	println()

}

func (list *Todos) CountPending() string {
	pending := 0

	for _, item := range *list {
		if !item.Completed {
			pending++
		}
	}
	if list.Is_empty() {
		return gray(fmt.Sprint("Nothing To do!"))
	}

	if pending == 0 {
		return green(fmt.Sprint("No todos available!"))
	}
	return red(fmt.Sprintf("You have %d pending things to do", pending))

}

func (list *Todos) Is_empty() bool {
	return len(*list) == 0
}

func (list *Todos) timeLeft() string {
	ls := *list
	if len(ls) == 0 {
		return "Time Left: 0"
	}
	item := ls[0]
	diff := time.Until(item.CreatedAt)
	out := time.Time{}.Add(diff)
	timer := fmt.Sprint("Time Left: ", out.Format("15:04:05"))
	return timer
}
