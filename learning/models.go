package main

import (
	"time"
)

// Event represents an RSVP event
type Event struct {
	ID          string
	Name        string
	Date        time.Time
	Description string
	Attendees   []Attendee
	CreatedAt   time.Time
}

// Attendee represents a person responding to an event
type Attendee struct {
	Name      string
	Status    string // "yes", "no", or "maybe"
	RSVPedAt  time.Time
}

// NewEvent creates a new event
func NewEvent(name string, date time.Time, description string) *Event {
	return &Event{
		ID:          generateID(),
		Name:        name,
		Date:        date,
		Description: description,
		Attendees:   []Attendee{},
		CreatedAt:   time.Now(),
	}
}

// FindAttendee searches for an attendee by name
func (e *Event) FindAttendee(name string) (int, *Attendee, bool) {
	for i, attendee := range e.Attendees {
		if attendee.Name == name {
			return i, &attendee, true
		}
	}
	return -1, nil, false
}

// RSVPResult holds aggregated RSVP statistics
type RSVPResult struct {
	YesCount   int
	NoCount    int
	MaybeCount int
}

// Total returns the total number of attendees
func (r *RSVPResult) Total() int {
	return r.YesCount + r.NoCount + r.MaybeCount
}
