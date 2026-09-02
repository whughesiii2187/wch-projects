package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	store := NewEventStore()
	
	fmt.Println("=== RSVP System ===")
	fmt.Println("Commands: create-event, list-events, view-event, rsvp, remove-rsvp, results, quit")
	fmt.Println()

	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("> ")
		if !scanner.Scan() {
			break
		}

		input := strings.TrimSpace(scanner.Text())
		if input == "" {
			continue
		}

		parts := strings.Fields(input)
		command := parts[0]

		switch command {
		case "create-event":
			handleCreateEvent(scanner, store)
		case "list-events":
			handleListEvents(store)
		case "view-event":
			if len(parts) < 2 {
				fmt.Println("Usage: view-event <event-id>")
				continue
			}
			handleViewEvent(parts[1], store)
		case "rsvp":
			if len(parts) < 4 {
				fmt.Println("Usage: rsvp <event-id> <attendee-name> <status> (yes/no/maybe)")
				continue
			}
			handleRSVP(parts[1], parts[2], parts[3], store)
		case "remove-rsvp":
			if len(parts) < 3 {
				fmt.Println("Usage: remove-rsvp <event-id> <attendee-name>")
				continue
			}
			handleRemoveRSVP(parts[1], parts[2], store)
		case "results":
			if len(parts) < 2 {
				fmt.Println("Usage: results <event-id>")
				continue
			}
			handleResults(parts[1], store)
		case "quit":
			fmt.Println("Goodbye!")
			return
		default:
			fmt.Println("Unknown command. Try: create-event, list-events, view-event, rsvp, remove-rsvp, results, quit")
		}
	}
}

func handleCreateEvent(scanner *bufio.Scanner, store *EventStore) {
	fmt.Print("Event name: ")
	scanner.Scan()
	name := strings.TrimSpace(scanner.Text())

	fmt.Print("Event date (YYYY-MM-DD): ")
	scanner.Scan()
	dateStr := strings.TrimSpace(scanner.Text())

	date, err := time.Parse("2006-01-02", dateStr)
	if err != nil {
		fmt.Printf("Invalid date format: %v\n", err)
		return
	}

	fmt.Print("Event description: ")
	scanner.Scan()
	description := strings.TrimSpace(scanner.Text())

	event := NewEvent(name, date, description)
	store.AddEvent(event)
	fmt.Printf("Event created with ID: %s\n", event.ID)
}

func handleListEvents(store *EventStore) {
	events := store.GetAllEvents()
	if len(events) == 0 {
		fmt.Println("No events found.")
		return
	}

	fmt.Println("\n=== Events ===")
	for _, event := range events {
		fmt.Printf("[%s] %s - %s\n", event.ID, event.Name, event.Date.Format("2006-01-02"))
	}
	fmt.Println()
}

func handleViewEvent(eventID string, store *EventStore) {
	event, found := store.GetEvent(eventID)
	if !found {
		fmt.Println("Event not found.")
		return
	}

	fmt.Printf("\n=== %s ===\n", event.Name)
	fmt.Printf("ID: %s\n", event.ID)
	fmt.Printf("Date: %s\n", event.Date.Format("2006-01-02"))
	fmt.Printf("Description: %s\n", event.Description)
	fmt.Printf("Total attendees: %d\n", len(event.Attendees))

	if len(event.Attendees) > 0 {
		fmt.Println("\nAttendees:")
		for _, attendee := range event.Attendees {
			fmt.Printf("  - %s: %s\n", attendee.Name, attendee.Status)
		}
	}
	fmt.Println()
}

func handleRSVP(eventID, attendeeName, status string, store *EventStore) {
	// Validate status
	if status != "yes" && status != "no" && status != "maybe" {
		fmt.Println("Status must be: yes, no, or maybe")
		return
	}

	err := store.AddRSVP(eventID, attendeeName, status)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	fmt.Printf("%s has RSVP'd '%s' to the event.\n", attendeeName, status)
}

func handleRemoveRSVP(eventID, attendeeName string, store *EventStore) {
	err := store.RemoveRSVP(eventID, attendeeName)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	fmt.Printf("%s has been removed from the event.\n", attendeeName)
}

func handleResults(eventID string, store *EventStore) {
	results, err := store.GetRSVPResults(eventID)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	fmt.Printf("\n=== RSVP Results ===\n")
	fmt.Printf("Yes: %d\n", results.YesCount)
	fmt.Printf("No: %d\n", results.NoCount)
	fmt.Printf("Maybe: %d\n", results.MaybeCount)
	fmt.Printf("Total: %d\n\n", results.Total())
}
