package api

import (
	"encoding/json"
	"net/http"

	"github.com/dougneal/playground/data"
)

func TodoList(rw http.ResponseWriter, r *http.Request) {
	todos, err := data.ListTodos()

	//TODO common error handling
	if err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}

	e := json.NewEncoder(rw)
	e.Encode(todos)
}

func TodoCreate(rw http.ResponseWriter, r *http.Request) {
}

func TodoGet(rw http.ResponseWriter, r *http.Request) {
}

func TodoUpdate(rw http.ResponseWriter, r *http.Request) {
}

func TodoDelete(rw http.ResponseWriter, r *http.Request) {
}
