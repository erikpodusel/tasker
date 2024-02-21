package service

import "time"

type TasksService struct{}

func (s TasksService) GetToday() time.Time {
	return time.Now()
}

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
