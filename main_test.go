package main

import (
	"fmt"
	"microserviceGo/requests_responses"
	"microserviceGo/service"
	"testing"
	"time"
)

func BenchmarkCreateTodoGo(b *testing.B) {
	todoService := service.NewTodoService()

	b.N = 10655712
	for i := 0; i < b.N; i++ {
		start := time.Now()

		createTodoRequest := requests_responses.CreateTodoRequest{
			ID:        i,
			Title:     fmt.Sprintf("Задача %d", i),
			Completed: false,
		}

		_, _ = todoService.CreateTodo(createTodoRequest)

		elapsed := time.Since(start)
		b.ReportMetric(float64(elapsed.Nanoseconds())/float64(time.Millisecond), "ms")
	}
}
