package handler

import (
	"html/template"
	"log"
	"net/http"
	"tasker/service"
	"time"
)

type TasksHandler struct {
	TasksService service.TasksService
}

type Day struct {
	Today time.Time
}

func (h TasksHandler) ShowTasks(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.New("").ParseFiles("./view/layout/index.html", "./view/page/tasks.html")
	if err != nil {
		log.Fatal(err)
		return
	}

	data := map[string]string{
		"Today": h.TasksService.GetToday().Weekday().String(),
	}

	tmpl.ExecuteTemplate(w, "base", data)
}
