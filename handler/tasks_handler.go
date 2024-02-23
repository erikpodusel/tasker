package handler

import (
	"html/template"
	"log"
	"net/http"
	"tasker/model"
	"tasker/service"
)

type TasksHandler struct {
	TasksService service.TasksService
}

func (h TasksHandler) ShowTasks(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.New("").ParseFiles("./view/layout/index.html", "./view/page/tasks.html")
	if err != nil {
		log.Fatal(err)
		return
	}

	response := map[string][]model.Weekday{
		"Response": h.TasksService.GetTasksForCurrentWeek(),
	}

	tmpl.ExecuteTemplate(w, "base", response)
}
