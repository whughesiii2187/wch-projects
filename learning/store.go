package main

import (
	"fmt"
	"sync"
	"time"
)

// EventStore manages all events and RSVPs
type EventStore struct {
	events map[string]*Event
	mu     sync.RWMutex
}

// NewEventStore creates a new event store
func NewEventStore() *EventStore {
	return &EventStore{
		events: make(map[string]*Event),
	}
}

// AddEvent adds a new event to the store
func (es *EventStore) AddEvent(event *Event) {
	es.mu.Lock()
	defer es.mu.Unlock()
	es.events[event.ID] = event
}

// GetEvent retrieves an event by ID
func (es *EventStore) GetEvent(id string) (*Event, bool) {
	es.mu.RLock()
	defer es.mu.RUnlock()
	event, found := es.events[id]
	return event, found
}

// GetAllEvents returns all events
func (es *EventStore) GetAllEvents() []*Event {
	es.mu.RLock()
	defer es.mu.RUnlock()
	events := make([]*Event, 0, len(es.events))
	for _, event := range es.events {
		events = append(events, event)
	}
	return events
}

// AddRSVP adds or updates an attendee's RSVP
func (es *EventStore) AddRSVP(eventID, attendeeName, status string) error {
	es.mu.Lock()
	defer es.mu.Unlock()

	event, found := es.events[eventID]
	if !found {
		return fmt.Errorf("event not found")
	}

	// Check if attendee already exists
	if idx, _, found := event.FindAttendee(attendeeName); found {
		// Update existing attendee
		event.Attendees[idx].Status = status
		event.Attendees[idx].RSVPedAt = time.Now()
		return nil
	}

	// Add new attendee
	attendee := Attendee{
		Name:     attendeeName,
		Status:   status,
		RSVPedAt: time.Now(),
	}
	event.Attendees = append(event.Attendees, attendee)
	return nil
}

// RemoveRSVP removes an attendee from an event
func (es *EventStore) RemoveRSVP(eventID, attendeeName string) error {
	es.mu.Lock()
	defer es.mu.Unlock()

	event, found := es.events[eventID]
	if !found {
		return fmt.Errorf("event not found")
	}

	idx, _, found := event.FindAttendee(attendeeName)
	if !found {
		return fmt.Errorf("attendee not found")
	}

	// Remove attendee by slicing
	event.Attendees = append(event.Attendees[:idx], event.Attendees[idx+1:]...)
	return nil
}

// GetRSVPResults aggregates RSVP statistics for an event
func (es *EventStore) GetRSVPResults(eventID string) (*RSVPResult, error) {
	es.mu.RLock()
	defer es.mu.RUnlock()

	event, found := es.events[eventID]
	if !found {
		return nil, fmt.Errorf("event not found")
	}

	results := &RSVPResult{}
	for _, attendee := range event.Attendees {
		switch attendee.Status {
		case "yes":
			results.YesCount++
		case "no":
			results.NoCount++
		case "maybe":
			results.MaybeCount++
		}
	}

	return results, nil
}

// DeleteEvent removes an event from the store
func (es *EventStore) DeleteEvent(eventID string) error {
	es.mu.Lock()
	defer es.mu.Unlock()

	if _, found := es.events[eventID]; !found {
		return fmt.Errorf("event not found")
	}

	delete(es.events, eventID)
	return nil
}
