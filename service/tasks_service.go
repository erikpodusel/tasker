package service

import "time"

type TasksService struct{}

func (s TasksService) GetToday() time.Time {
	return time.Now()
}
