package repository

import (
	"errors"
	"wb_lvl2_calendar/internal/models"
)

type CalendarRepository struct {
	lastIndex int
	Events    []models.Event
}

func NewCalendarRepository() CalendarRepository {
	events := make([]models.Event, 0)
	return CalendarRepository{lastIndex: 0, Events: events}
}

type ICalendarRepository interface {
	CreateEvent(event models.Event) int
	UpdateEvent(newEvent models.Event) error
	DeleteEvent(eventId int) error
}

func (cr *CalendarRepository) CreateEvent(event models.Event) int {
	event.EventId = cr.lastIndex
	cr.lastIndex++
	cr.Events = append(cr.Events, event)
	return event.EventId
}

func (cr *CalendarRepository) UpdateEvent(newEvent models.Event) error {
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

func (cr *CalendarRepository) getEvent(eventId int) (models.Event, error) {
	for _, val := range cr.Events {
		if val.EventId == eventId {
			return val, nil
		}
	}
	return models.Event{}, errors.New("event not found")
}

func (cr *CalendarRepository) getEventIndex(eventId int) (int, error) {
	for ind, val := range cr.Events {
		if val.EventId == eventId {
			return ind, nil
		}
	}
	return -1, errors.New("event not found")
}
