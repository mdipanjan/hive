package session

import "time"

type Status int

const (
	StatusIdle Status = iota
	StatusRunning
	StatusWaiting
)

type Session struct {
	Name         string
	Tool         string
	Path         string
	Status       Status
	CreatedAt    time.Time
	LastActivity time.Time
}
