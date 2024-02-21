package handler

import (
	"html/template"
	"log"
	"net/http"
	"tasker/service"
)

type TasksHandler struct {
	TasksService service.TasksService
}

type Weekdays struct {
	Weekdays []string
}

func (h TasksHandler) ShowTasks(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.New("").ParseFiles("./view/layout/index.html", "./view/page/tasks.html")
	if err != nil {
		log.Fatal(err)
		return
	}

	data := Weekdays{}
	weekDays := h.TasksService.GetCurrentWeekDays()

	for _, day := range weekDays {
		data.Weekdays = append(data.Weekdays, day.Weekday().String())
	}

	tmpl.ExecuteTemplate(w, "base", data)
}
