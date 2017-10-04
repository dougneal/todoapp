package data

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"time"
)

type Todo struct {
	ID        uint      `json:"-"`
	Text      string    `json:"text"`
	UpdatedAt time.Time `json:"updated_at"`
}

func ListTodos() ([]Todo, error) {
	file, err := os.Open(filepath.Join(dataDirPath(), "todo"))

	if err != nil {
		return nil, err
	}
	defer file.Close()

	filenames, err := file.Readdirnames(-1)
	if err != nil {
		return nil, err
	}

	todos := make([]Todo, 10, 10)
	for _, filename := range filenames {
		id, err := strconv.ParseUint(filename, 10, 32)
		if err == nil {
			t := Todo{
				ID: uint(id),
			}
			err = t.Get()
			if err != nil {
				return nil, err
			}

			todos = append(todos, t)
		}
	}

	return todos, nil
}

func (t *Todo) Get() error {
	filename := getTodoFilename(t.ID)

	fileInfo, err := os.Stat(filename)
	if err != nil {
		return err
	}

	t.UpdatedAt = fileInfo.ModTime()

	file, err := os.Open(filename)
	if err != nil {
		return err
	}

	defer file.Close()
	buf := make([]byte, 1024*16, 1024*16)
	_, err = file.Read(buf)
	if err != nil {
		return err
	}

	t.Text = string(buf)

	return nil
}

func (t *Todo) Upsert() error {
	file, err := os.Create(getTodoFilename(t.ID))
	if err != nil {
		return err
	}

	defer file.Close()
	_, err = file.Write([]byte(t.Text))
	if err != nil {
		return err
	}

	t.UpdatedAt = time.Now()
	return nil
}

func (t *Todo) Delete() error {
	return os.Remove(getTodoFilename(t.ID))
}

func getTodoFilename(id uint) string {
	return filepath.Join(
		dataDirPath(),
		"todo",
		fmt.Sprintf("%d", id),
	)
}
