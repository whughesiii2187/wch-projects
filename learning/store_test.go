package main

import (
	"testing"
	"time"
)

// TestAddEvent tests adding an event to the store
func TestAddEvent(t *testing.T) {
	store := NewEventStore()
	event := NewEvent("Birthday Party", time.Now(), "Celebrating!")

	store.AddEvent(event)

	retrieved, found := store.GetEvent(event.ID)
	if !found {
		t.Fatal("Event not found after adding")
	}

	if retrieved.Name != "Birthday Party" {
		t.Errorf("Expected name 'Birthday Party', got %s", retrieved.Name)
	}
}

// TestGetAllEvents tests retrieving all events
func TestGetAllEvents(t *testing.T) {
	store := NewEventStore()

	event1 := NewEvent("Event 1", time.Now(), "First")
	event2 := NewEvent("Event 2", time.Now(), "Second")

	store.AddEvent(event1)
	store.AddEvent(event2)

	events := store.GetAllEvents()
	if len(events) != 2 {
		t.Errorf("Expected 2 events, got %d", len(events))
	}
}

// TestAddRSVP tests adding an RSVP to an event
func TestAddRSVP(t *testing.T) {
	store := NewEventStore()
	event := NewEvent("Test Event", time.Now(), "Test")
	store.AddEvent(event)

	err := store.AddRSVP(event.ID, "Alice", "yes")
	if err != nil {
		t.Fatalf("Error adding RSVP: %v", err)
	}

	retrieved, _ := store.GetEvent(event.ID)
	if len(retrieved.Attendees) != 1 {
		t.Errorf("Expected 1 attendee, got %d", len(retrieved.Attendees))
	}

	if retrieved.Attendees[0].Name != "Alice" {
		t.Errorf("Expected attendee 'Alice', got %s", retrieved.Attendees[0].Name)
	}

	if retrieved.Attendees[0].Status != "yes" {
		t.Errorf("Expected status 'yes', got %s", retrieved.Attendees[0].Status)
	}
}

// TestUpdateRSVP tests updating an existing RSVP
func TestUpdateRSVP(t *testing.T) {
	store := NewEventStore()
	event := NewEvent("Test Event", time.Now(), "Test")
	store.AddEvent(event)

	store.AddRSVP(event.ID, "Bob", "yes")
	store.AddRSVP(event.ID, "Bob", "no")

	retrieved, _ := store.GetEvent(event.ID)
	if len(retrieved.Attendees) != 1 {
		t.Errorf("Expected 1 attendee after update, got %d", len(retrieved.Attendees))
	}

	if retrieved.Attendees[0].Status != "no" {
		t.Errorf("Expected updated status 'no', got %s", retrieved.Attendees[0].Status)
	}
}

// TestRemoveRSVP tests removing an attendee
func TestRemoveRSVP(t *testing.T) {
	store := NewEventStore()
	event := NewEvent("Test Event", time.Now(), "Test")
	store.AddEvent(event)

	store.AddRSVP(event.ID, "Charlie", "maybe")
	err := store.RemoveRSVP(event.ID, "Charlie")
	if err != nil {
		t.Fatalf("Error removing RSVP: %v", err)
	}

	retrieved, _ := store.GetEvent(event.ID)
	if len(retrieved.Attendees) != 0 {
		t.Errorf("Expected 0 attendees after removal, got %d", len(retrieved.Attendees))
	}
}

// TestGetRSVPResults tests aggregating RSVP statistics
func TestGetRSVPResults(t *testing.T) {
	store := NewEventStore()
	event := NewEvent("Test Event", time.Now(), "Test")
	store.AddEvent(event)

	store.AddRSVP(event.ID, "Alice", "yes")
	store.AddRSVP(event.ID, "Bob", "yes")
	store.AddRSVP(event.ID, "Charlie", "no")
	store.AddRSVP(event.ID, "Diana", "maybe")

	results, err := store.GetRSVPResults(event.ID)
	if err != nil {
		t.Fatalf("Error getting results: %v", err)
	}

	if results.YesCount != 2 {
		t.Errorf("Expected 2 'yes', got %d", results.YesCount)
	}

	if results.NoCount != 1 {
		t.Errorf("Expected 1 'no', got %d", results.NoCount)
	}

	if results.MaybeCount != 1 {
		t.Errorf("Expected 1 'maybe', got %d", results.MaybeCount)
	}

	if results.Total() != 4 {
		t.Errorf("Expected total 4, got %d", results.Total())
	}
}

// TestDeleteEvent tests removing an event
func TestDeleteEvent(t *testing.T) {
	store := NewEventStore()
	event := NewEvent("Test Event", time.Now(), "Test")
	store.AddEvent(event)

	err := store.DeleteEvent(event.ID)
	if err != nil {
		t.Fatalf("Error deleting event: %v", err)
	}

	_, found := store.GetEvent(event.ID)
	if found {
		t.Fatal("Event still found after deletion")
	}
}

// TestAddRSVPToNonexistentEvent tests error handling
func TestAddRSVPToNonexistentEvent(t *testing.T) {
	store := NewEventStore()

	err := store.AddRSVP("nonexistent", "Alice", "yes")
	if err == nil {
		t.Fatal("Expected error when adding RSVP to nonexistent event")
	}
}
