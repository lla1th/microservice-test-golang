package service

import (
	"errors"
	"microserviceGo/requests_responses"
)

type TodoService interface {
	CreateTodo(request requests_responses.CreateTodoRequest) (requests_responses.CreateTodoResponse, error)
	GetTodoByID(request requests_responses.GetTodoRequest) (requests_responses.GetTodoResponse, error)
}

type todoServiceImpl struct {
	todos map[int]Todo // Пример: картотека задач
}

func NewTodoService() TodoService {
	return &todoServiceImpl{
		todos: make(map[int]Todo),
	}
}

func (t *todoServiceImpl) CreateTodo(request requests_responses.CreateTodoRequest) (requests_responses.CreateTodoResponse, error) {
	// Создаем новую задачу из запроса
	newTodo := Todo{
		ID:        request.ID,
		Title:     request.Title,
		Completed: request.Completed,
	}

	// Сохраняем задачу в картотеке
	t.todos[newTodo.ID] = newTodo

	// Возвращаем созданную задачу в ответе
	return requests_responses.CreateTodoResponse{ID: newTodo.ID}, nil
}

func (t *todoServiceImpl) GetTodoByID(request requests_responses.GetTodoRequest) (requests_responses.GetTodoResponse, error) {
	todo, found := t.todos[request.ID]
	if !found {
		return requests_responses.GetTodoResponse{}, errors.New("задача не найдена")
	}

	return requests_responses.GetTodoResponse(todo), nil
}
