package tasks

type TaskCreateRequest struct {
	Description string `json:"description" validate:"required"`
	Status      string `json:"status"`
	Epic        string `json:"epic"`
	Daily       bool   `json:"daily"`
}
