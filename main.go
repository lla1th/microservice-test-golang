package main

import (
	"github.com/gorilla/mux"
	"log"
	"microserviceGo/endpoints"
	"microserviceGo/handlers"
	"microserviceGo/service"
	"net/http"
)

func main() {
	router := mux.NewRouter()

	todoService := service.NewTodoService()
	todoHandler := handlers.NewTodoHandler(todoService)
	endpoints.RegisterTodoEndpoints(router, todoHandler)

	http.Handle("/", router)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
