package tasks

import (
	"encoding/json"
	"net/http"
	"strconv"
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
	router.HandleFunc("DELETE /tasks/{id}", handler.Delete())
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

func (handler *TaskHandler) Delete() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		idString := r.PathValue("id")
		id, err := strconv.ParseUint(idString, 10, 32)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		_, err = handler.TaskRepository.GetById(uint(id))
		if err != nil {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}
		err = handler.TaskRepository.Delete(uint(id))
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(200)
		json.NewEncoder(w).Encode(nil)
	}
}
