package todo

import (
	"encoding/json"
	"errors"
	"os"
	"time"
)

type Element struct {
	Task        string
	Done        bool
	CreatedAt   time.Time
	CompletedAt time.Time
}

type Todos []Element

func (t *Todos) Add(task string) {
	todo := Element{
		Task:        task,
		Done:        false,
		CreatedAt:   time.Now(),
		CompletedAt: time.Time{},
	}
	*t = append(*t, todo)
}

func (t *Todos) Complete(number int) error {
	items := *t
	if number <= 0 || number > len(items) {
		return errors.New("invalid index number")
	}

	items[number-1].CompletedAt = time.Now()
	items[number-1].Done = true

	return nil
}

func (t *Todos) Delete(number int) error {
	items := *t
	if number <= 0 || number > len(items) {
		return errors.New("invalid index number")
	}
	*t = append(items[:number-1], items[number:]...)
	return nil
}

func (t *Todos) Load(filename string) error {
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
	err = json.Unmarshal(file, t)
	if err != nil {
		return err
	}
	return nil
}
