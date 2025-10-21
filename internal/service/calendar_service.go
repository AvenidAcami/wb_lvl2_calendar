package service

import (
	"errors"
	"fmt"
	"strconv"
	"time"
	"wb_lvl2_calendar/internal/model"
	"wb_lvl2_calendar/internal/repository"
)

type CalendarService struct {
	repo repository.ICalendarRepository
}

type ICalendarService interface {
	CreateEvent(event model.EventInput) (int, error)
	UpdateEvent(newEvent model.EventInput) error
	DeleteEvent(eventId string) error
	GetEventsForDay(userId string, date string) ([]model.Event, error)
	GetEventsForWeek(userId string, date string) ([]model.Event, error)
	GetEventsForMonth(userId string, date string) ([]model.Event, error)
}

func NewCalendarService(repo repository.ICalendarRepository) *CalendarService {
	return &CalendarService{repo: repo}
}

func (cs *CalendarService) CreateEvent(event model.EventInput) (int, error) {
	t, err := cs.stringToTime(event.Date)
	if err != nil {
		return -1, err
	}
	formattedEvent := model.Event{
		EventId:     event.EventId,
		UserId:      event.EventId,
		Date:        t,
		Description: event.Description}
	return cs.repo.CreateEvent(formattedEvent), nil
}

func (cs *CalendarService) UpdateEvent(newEvent model.EventInput) error {
	t, err := cs.stringToTime(newEvent.Date)
	if err != nil {
		return err
	}
	formattedEvent := model.Event{
		EventId:     newEvent.EventId,
		UserId:      newEvent.EventId,
		Date:        t,
		Description: newEvent.Description}
	return cs.repo.UpdateEvent(formattedEvent)
}

func (cs *CalendarService) DeleteEvent(eventId string) error {
	str, err := strconv.Atoi(eventId)
	if err != nil {
		return errors.New("id must be integer")
	}
	return cs.repo.DeleteEvent(str)
}

func (cs *CalendarService) GetEventsForDay(userId string, date string) ([]model.Event, error) {
	str, err := strconv.Atoi(userId)
	if err != nil {
		return []model.Event{}, errors.New("id must be integer")
	}
	t, err := cs.stringToTime(date)
	if err != nil {
		return []model.Event{}, err
	}
	return cs.repo.GetEventsForDay(str, t), nil
}

func (cs *CalendarService) GetEventsForWeek(userId string, date string) ([]model.Event, error) {
	str, err := strconv.Atoi(userId)
	if err != nil {
		return []model.Event{}, errors.New("id must be integer")
	}
	t, err := cs.stringToTime(date)
	if err != nil {
		return []model.Event{}, err
	}
	return cs.repo.GetEventsForWeek(str, t), nil
}

func (cs *CalendarService) GetEventsForMonth(userId string, date string) ([]model.Event, error) {
	str, err := strconv.Atoi(userId)
	if err != nil {
		return []model.Event{}, errors.New("id must be integer")
	}
	t, err := cs.stringToTime(date)
	if err != nil {
		return []model.Event{}, err
	}
	return cs.repo.GetEventsForMonth(str, t), nil
}

func (cs *CalendarService) stringToTime(date string) (time.Time, error) {
	t, err := time.Parse("2006-01-02", date)
	if err != nil {
		fmt.Println("Parse error:", err)
		return t, errors.New("wrong date")
	}
	return t, nil
}
