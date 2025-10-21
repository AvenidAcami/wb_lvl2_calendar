package repository

import (
	"errors"
	"time"
	"wb_lvl2_calendar/internal/model"
)

type CalendarRepository struct {
	lastIndex int
	Events    []model.Event
}

func NewCalendarRepository() *CalendarRepository {
	events := make([]model.Event, 0)
	return &CalendarRepository{lastIndex: 0, Events: events}
}

type ICalendarRepository interface {
	CreateEvent(event model.Event) int
	UpdateEvent(newEvent model.Event) error
	DeleteEvent(eventId int) error
	GetEventsForDay(userId int, date time.Time) []model.Event
	GetEventsForWeek(userId int, date time.Time) []model.Event
	GetEventsForMonth(userId int, date time.Time) []model.Event
}

func (cr *CalendarRepository) CreateEvent(event model.Event) int {
	event.EventId = cr.lastIndex
	cr.lastIndex++
	cr.Events = append(cr.Events, event)
	return event.EventId
}

func (cr *CalendarRepository) UpdateEvent(newEvent model.Event) error {
	oldEvent, err := cr.getEvent(newEvent.EventId)
	if err != nil {
		return err
	}
	cr.Events[oldEvent.EventId] = newEvent
	return nil
}

func (cr *CalendarRepository) DeleteEvent(eventId int) error {
	ind, err := cr.getEventIndex(eventId)
	if err != nil {
		return err
	}
	cr.Events[ind] = cr.Events[len(cr.Events)-1]
	cr.Events = cr.Events[:len(cr.Events)-1]
	return nil
}

func (cr *CalendarRepository) getEvent(eventId int) (model.Event, error) {
	for _, val := range cr.Events {
		if val.EventId == eventId {
			return val, nil
		}
	}
	return model.Event{}, errors.New("event not found")
}

func (cr *CalendarRepository) GetEventsForDay(userId int, date time.Time) []model.Event {
	events := make([]model.Event, 0)
	for _, val := range cr.Events {
		if val.Date.Day() == date.Day() {
			events = append(events, val)
		}
	}
	return events
}

func (cr *CalendarRepository) GetEventsForWeek(userId int, date time.Time) []model.Event {
	events := make([]model.Event, 0)
	for _, val := range cr.Events {
		if weekNumber(val.Date) == weekNumber(date) {
			events = append(events, val)
		}
	}
	return events
}

func (cr *CalendarRepository) GetEventsForMonth(userId int, date time.Time) []model.Event {
	events := make([]model.Event, 0)
	for _, val := range cr.Events {
		if val.Date.Month() == date.Month() {
			events = append(events, val)
		}
	}
	return events
}

func (cr *CalendarRepository) getEventIndex(eventId int) (int, error) {
	for ind, val := range cr.Events {
		if val.EventId == eventId {
			return ind, nil
		}
	}
	return -1, errors.New("event not found")
}

func weekNumber(t time.Time) int {
	_, week := t.ISOWeek()
	return week
}
