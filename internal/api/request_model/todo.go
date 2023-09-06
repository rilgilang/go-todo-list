package request_model

type CreateTodo struct {
	Title string `json:"title,omitempty"`
}

type UpdateTodo struct {
	Id     int  `json:"id,omitempty"`
	Status bool `json:"status,omitempty"`
}

type DeleteTodo struct {
	Id int `json:"id,omitempty"`
}
