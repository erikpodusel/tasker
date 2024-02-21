package main

import (
	"log"
	"net/http"
	"os"
	"tasker/handler"
	"tasker/service"
)

func main() {
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	tasksService := service.TasksService{}
	tasksHandler := handler.TasksHandler{TasksService: tasksService}

	http.HandleFunc("/", tasksHandler.ShowTasks)

	port := os.Getenv("PORT")
	if port == "" {
		port = ":8080"
	}

	log.Println("Server started at", port)
	log.Fatal(http.ListenAndServe(port, nil))
}
