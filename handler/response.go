package handler

import (
	"time"
)

type Task struct {
	ID       int64     `json:"id" db:"id"`
	Title    string    `json:"title" db:"title"`
	Status   string    `json:"status" db:"status"`
	Created  time.Time `json:"created" db:"created"`
	Modified time.Time `json:"modified" db:"modified"`
}

type T struct {
	Time time.Time `json:"time"`
}

func (t *T) Now() time.Time {
	t.Time = time.Now()
	return t.Time
}
