package model

type Task struct {
	Id           string
	Status       string
	SystemStatus string
	Title        string
}

type Weekday struct {
	Day   string
	Tasks []Task
}
