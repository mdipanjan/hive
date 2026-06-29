package session

import "time"

type Status int

const (
	StatusIdle Status = iota
	StatusActive
	StatusRunning
	StatusReady
	StatusCompleted
	StatusFailed
	StatusUnknown
)

func (s Status) String() string {
	switch s {
	case StatusActive:
		return "active"
	case StatusRunning:
		return "running"
	case StatusReady:
		return "ready"
	case StatusCompleted:
		return "completed"
	case StatusFailed:
		return "failed"
	case StatusUnknown:
		return "unknown"
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
