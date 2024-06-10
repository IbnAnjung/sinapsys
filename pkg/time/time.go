package time

import "time"

type Time interface {
	Now() time.Time
	GetLastTimeOnDay() time.Time
	GetDefaultLoc() *time.Location
}

type timeService struct {
	loc *time.Location
}

func NewTimeService() Time {
	loc, _ := time.LoadLocation("Asia/Jakarta")
	return &timeService{
		loc,
	}
}

func (t *timeService) Now() time.Time {
	return time.Now().In(t.loc)
}

func (t *timeService) GetLastTimeOnDay() time.Time {
	now := t.Now()
	return time.Date(now.Year(), now.Month(), now.Day(), 23, 59, 59, 59, t.loc)
}

func (t *timeService) GetDefaultLoc() *time.Location {
	loc, _ := time.LoadLocation("Asia/Jakarta")
	return loc
}
