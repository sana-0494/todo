package api

import "time"

type todoCreateRequest struct {
	Title string `json:"title"`
}

type todoUpdateRequest struct {
	Id string `json:"id"`
}

type todoResponse struct {
	Title     string    `json:"title"`
	CreatedAt time.Time `json:"created_at"`
}

type todoDeleteResponse struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}
