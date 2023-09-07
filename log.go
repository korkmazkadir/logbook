package logbook

import "time"

type Log struct {
	ID      uint64    `json:"id"`
	Time    time.Time `json:"time"`
	Content string    `json:"content"`
}
