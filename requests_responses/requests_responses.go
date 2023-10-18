package requests_responses

type CreateTodoRequest struct {
	ID        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

type CreateTodoResponse struct {
	ID int `json:"id"`
}

type GetTodoRequest struct {
	ID int `json:"id"`
}

type GetTodoResponse struct {
	ID        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}
