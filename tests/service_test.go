package service_test

import (
	"strconv"
	"testing"
	"time"

	"wb_lvl2_calendar/internal/model"
	"wb_lvl2_calendar/internal/repository"
	"wb_lvl2_calendar/internal/service"
)

func TestCreateEvent(t *testing.T) {
	repo := repository.NewCalendarRepository()
	s := service.NewCalendarService(repo)

	eventInput := model.EventInput{
		EventId:     0,
		UserId:      1,
		Date:        "2025-10-24",
		Description: "Test event",
	}

	id, err := s.CreateEvent(eventInput)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if id != 0 {
		t.Errorf("expected event id 0, got %d", id)
	}

	if len(repo.Events) != 1 {
		t.Errorf("expected 1 event in repo, got %d", len(repo.Events))
	}
}

func TestUpdateEvent(t *testing.T) {
	repo := repository.NewCalendarRepository()
	s := service.NewCalendarService(repo)

	repo.CreateEvent(model.Event{
		EventId:     0,
		UserId:      1,
		Date:        time.Now(),
		Description: "Old",
	})

	input := model.EventInput{
		EventId:     0,
		UserId:      1,
		Date:        "2025-10-25",
		Description: "Updated",
	}

	err := s.UpdateEvent(input)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if repo.Events[0].Description != "Updated" {
		t.Errorf("expected 'Updated', got '%s'", repo.Events[0].Description)
	}
}

func TestDeleteEvent(t *testing.T) {
	repo := repository.NewCalendarRepository()
	s := service.NewCalendarService(repo)

	id := repo.CreateEvent(model.Event{EventId: 0, UserId: 1, Date: time.Now(), Description: "delete me"})

	err := s.DeleteEvent(strconv.Itoa(id))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if len(repo.Events) != 0 {
		t.Errorf("expected 0 events, got %d", len(repo.Events))
	}

	err = s.DeleteEvent("notNumber")
	if err == nil {
		t.Errorf("expected error for invalid id, got nil")
	}
}

func TestGetEventsForDay(t *testing.T) {
	repo := repository.NewCalendarRepository()
	s := service.NewCalendarService(repo)

	date := time.Date(2025, 10, 24, 0, 0, 0, 0, time.UTC)
	repo.CreateEvent(model.Event{EventId: 0, UserId: 1, Date: date, Description: "Today"})
	repo.CreateEvent(model.Event{EventId: 1, UserId: 1, Date: date.AddDate(0, 0, 1), Description: "Tomorrow"})

	events, err := s.GetEventsForDay("1", "2025-10-24")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(events) != 1 {
		t.Errorf("expected 1 event for day, got %d", len(events))
	}
}

func TestGetEventsForWeek(t *testing.T) {
	repo := repository.NewCalendarRepository()
	s := service.NewCalendarService(repo)

	date := time.Date(2025, 10, 20, 0, 0, 0, 0, time.UTC)
	repo.CreateEvent(model.Event{EventId: 0, UserId: 1, Date: date, Description: "Week event"})
	repo.CreateEvent(model.Event{EventId: 1, UserId: 1, Date: date.AddDate(0, 0, 10), Description: "Next week"})

	events, err := s.GetEventsForWeek("1", "2025-10-22")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(events) != 1 {
		t.Errorf("expected 1 event for week, got %d", len(events))
	}
}

func TestGetEventsForMonth(t *testing.T) {
	repo := repository.NewCalendarRepository()
	s := service.NewCalendarService(repo)

	date := time.Date(2025, 10, 10, 0, 0, 0, 0, time.UTC)
	repo.CreateEvent(model.Event{EventId: 0, UserId: 1, Date: date, Description: "This month"})
	repo.CreateEvent(model.Event{EventId: 1, UserId: 1, Date: date.AddDate(0, 1, 0), Description: "Next month"})

	events, err := s.GetEventsForMonth("1", "2025-10-01")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(events) != 1 {
		t.Errorf("expected 1 event for month, got %d", len(events))
	}
}
