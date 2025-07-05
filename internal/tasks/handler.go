package tasks

import (
	"encoding/json"
	"net/http"
	"task-tracker/config"
	"task-tracker/pkg/req"
)

type TaskHandlerDeps struct {
	Config         *config.Config
	TaskRepository *TaskRepository
}

type TaskHandler struct {
	TaskRepository *TaskRepository
}

func NewTaskHandler(router *http.ServeMux, deps TaskHandlerDeps) {
	handler := &TaskHandler{
		TaskRepository: deps.TaskRepository,
	}
	router.HandleFunc("POST /tasks", handler.Create())
}

func (handler *TaskHandler) Create() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		body, err := req.HandleBody[TaskCreateRequest](w, r)
		if err != nil {
			return
		}
		task := NewTask(body.Description, body.Status, body.Epic, body.Daily)
		createdTask, err := handler.TaskRepository.Create(task)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(createdTask)
	}
}
