package logger

import (
	"time"
)

type Event struct {
	input       string
	description string
	timestamp   time.Time
}

type EventList []Event

func createEvent(input, description string) Event {
	return Event{
		input:       input,
		description: description,
		timestamp:   time.Now(),
	}
}
