package internal

import "time"

type TaskStatus string

const (
	STATUS_TODO        TaskStatus = "todo"
	STATUS_IN_PROGRESS TaskStatus = "in-progress"
	STATUS_DONE        TaskStatus = "done"
)

type Task struct {
	Id          int64      `json:"id"`
	Description string     `json:"description"`
	Status      TaskStatus `json:"status"`
	CreatedAt   time.Time  `json:"createdat"`
	UpdatedAt   time.Time  `json:"updatedat"`
}

func Newtask(id int64, description string) *Task {
	return &Task{
		Id:          id,
		Description: description,
		Status:      STATUS_TODO,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
}
