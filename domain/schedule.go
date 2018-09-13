package domain

import "time"

type ScheduleID int

type Schedule struct {
	ID ScheduleID
	Interval time.Duration
}
