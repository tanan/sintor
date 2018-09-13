package domain

import "time"

type MonitorID int

type Method int

type Monitor struct {
	ID       MonitorID
	Request  Request
	Schedule Schedule
}

type MonitorResultID int

type MonitorResult struct {
	ID        MonitorResultID
	MonitorID MonitorID
	Date      time.Time
	Response  Response
}
