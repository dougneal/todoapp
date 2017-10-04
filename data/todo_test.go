package data

import (
	"os"
	//"fmt"
	"path/filepath"
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMain(t *testing.T) {
	_, thisFile, _, _ := runtime.Caller(0)

	os.Setenv(
		"DATA_DIR",
		filepath.Join(
			filepath.Dir(thisFile),
			"fixtures",
		),
	)
}

func TestListTodos(t *testing.T) {
	todos, err := ListTodos()
	assert.Nil(t, err)
	assert.NotNil(t, todos)
	assert.Equal(t, len(todos), 2)
}

func TestGetTodo(t *testing.T) {
	todo := Todo{ID: 1}
	err := todo.Get()
	assert.Nil(t, err)
	assert.Equal(t, todo.Text, "test todo text")
}

func TestTodo(t *testing.T) {
	todo := Todo{ID: 1}
	err := todo.Get()
	assert.Nil(t, err)
	assert.Equal(t, todo.Text, "test todo text")
	todo.Text = "test todo text - updated"
	err = todo.Upsert()
	assert.Nil(t, err)
	assert.Equal(t, todo.Text, "test todo text - updated")
	todo2 := Todo{ID: 1}
	err = todo2.Get()
	assert.Nil(t, err)
	assert.Equal(t, todo2.Text, "test todo text - updated")
}
