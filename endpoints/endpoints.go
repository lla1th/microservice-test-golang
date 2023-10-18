package endpoints

import (
	"github.com/gorilla/mux"
	"microserviceGo/handlers"
)

func RegisterTodoEndpoints(r *mux.Router, todoHandler handlers.TodoHandler) {
	r.Methods("POST").Path("/todos").HandlerFunc(todoHandler.CreateTodoHandler)
	r.Methods("GET").Path("/todos/{id}").HandlerFunc(todoHandler.GetTodoHandler)
}
