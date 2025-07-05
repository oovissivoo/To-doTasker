package main

import (
	"fmt"
	"net/http"
	"task-tracker/config"
	"task-tracker/internal/tasks"
	"task-tracker/pkg/db"
)

func main() {
	conf := config.LoadConfig()
	db := db.NewDb(conf)
	router := http.NewServeMux()

	//Repositories
	taskRepository := tasks.NewTaskRepository(db)

	//Handlers
	tasks.NewTaskHandler(router, tasks.TaskHandlerDeps{
		Config:         conf,
		TaskRepository: taskRepository,
	})

	server := http.Server{
		Addr:    ":8081",
		Handler: router,
	}
	fmt.Println("Server is listening")
	server.ListenAndServe()
}
