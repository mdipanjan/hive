package session

import "time"

type Status int

const (
	StatusIdle Status = iota
	StatusRunning
	StatusWaiting
)

func (s Status) String() string {
	switch s {
	case StatusRunning:
		return "running"
	case StatusWaiting:
		return "waiting"
	default:
		return "idle"
	}
}

type Session struct {
	Name         string
	Tool         string
	Path         string
	Status       Status
	CreatedAt    time.Time
	LastActivity time.Time
}
