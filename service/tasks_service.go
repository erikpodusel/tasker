package service

import (
	"fmt"
	"tasker/model"
	"time"
)

type TasksService struct{}

func (s TasksService) GetToday() time.Time {
	return time.Now()
}

/**
 * Get the current week days starting from Monday to Friday
 * @return []time.Time
 */
func (s TasksService) GetCurrentWeekDays() []time.Time {
	year, _ := time.Now().ISOWeek()
	month := time.Now().Month()

	t := time.Date(year, month, 0, 0, 0, 0, 0, time.UTC)

	// Roll back to Monday
	if wd := t.Weekday(); wd == time.Sunday {
		t = t.AddDate(0, 0, -6)
	} else {
		t = t.AddDate(0, 0, -int(wd)+1)
	}

	var days []time.Time

	for i := 0; i < 5; i++ {
		days = append(days, t.AddDate(0, 0, i))
	}

	return days
}

func (s TasksService) GetTasksForCurrentWeek() []model.Weekday {
	weekDays := s.GetCurrentWeekDays()
	tasks := make([]model.Weekday, len(weekDays))

	for i, day := range weekDays {
		task := model.Task{
			Id:           "1",
			Title:        "Task" + fmt.Sprint(i),
			Status:       "New",
			SystemStatus: "NEW",
		}

		tasks[i] = model.Weekday{
			Day:   day.Weekday().String(),
			Tasks: make([]model.Task, 1),
		}

		tasks[i].Tasks = append(tasks[i].Tasks, task)
	}

	return tasks
}
