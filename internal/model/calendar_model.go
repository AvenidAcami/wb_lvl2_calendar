package model

import "time"

type Event struct {
	EventId     int       `json:"event_id"`
	UserId      int       `json:"user_id"`
	Date        time.Time `json:"event_date"`
	Description string    `json:"description"`
}

type EventInput struct {
	EventId     int    `json:"event_id"`
	UserId      int    `json:"user_id"`
	Date        string `json:"event_date"`
	Description string `json:"description"`
}
