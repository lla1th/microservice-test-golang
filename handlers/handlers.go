package handlers

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"microserviceGo/requests_responses"
	"microserviceGo/service"
	"net/http"
	"strconv"
)

type TodoHandler struct {
	todoService service.TodoService
}

func NewTodoHandler(service service.TodoService) TodoHandler {
	return TodoHandler{todoService: service}
}

func (h *TodoHandler) CreateTodoHandler(w http.ResponseWriter, r *http.Request) {
	var req requests_responses.CreateTodoRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	createdTodo, err := h.todoService.CreateTodo(req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(createdTodo)
	if err != nil {
		http.Error(w, err.Error(), http.StatusServiceUnavailable)
		return
	}
}

func (h *TodoHandler) GetTodoHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	idStr := params["id"]
	id, err := strconv.Atoi(idStr)
	req := requests_responses.GetTodoRequest{ID: id}

	todos, err := h.todoService.GetTodoByID(req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(todos)
	if err != nil {
		http.Error(w, err.Error(), http.StatusServiceUnavailable)
		return
	}
}
