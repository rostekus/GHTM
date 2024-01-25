package domain

import "time"

type (
	TodoPriority int
)

const (
	Low TodoPriority = iota
	Medium
	High
)

type Todo struct {
	EndDate     time.Time
	Title       string
	Description string
	Id          string
	Priority    TodoPriority
	IsDone      bool
}
