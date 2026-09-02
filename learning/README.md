# RSVP System

A simple yet complete CRUD application for managing event RSVPs, built in Go.

## Project Structure

```
rsvp-system/
├── main.go          # CLI entry point and command handlers
├── models.go        # Data models (Event, Attendee, RSVPResult)
├── store.go         # In-memory event store with CRUD operations
├── utils.go         # Utility functions (ID generation)
└── go.mod           # Go module definition
```

## Features

- **Create Events**: Add new events with name, date, and description
- **Manage RSVPs**: Add, update, or remove attendee responses (yes/no/maybe)
- **View Events**: List all events or get details of a specific event
- **Get Results**: Aggregate RSVP statistics for any event
- **Thread-Safe**: Uses mutexes in the store for concurrent access

## Key Go Learning Points

This scaffold teaches you:

1. **Struct Organization**: Separating data models, business logic, and CLI handling
2. **Pointers & References**: Understanding when to use `*Event` vs `Event`
3. **Error Handling**: Returning errors explicitly with `error` type
4. **Concurrency Primitives**: Using `sync.RWMutex` for thread-safe access
5. **Methods**: Receiver functions on types (e.g., `func (e *Event) FindAttendee`)
6. **Interfaces**: Time to add these for persistence (database, JSON files)
7. **Slicing**: Managing attendee lists with append/indexing
8. **Maps**: Storing events by ID for fast lookup

## Setup

```bash
# Navigate to the project directory
cd rsvp-system

# Build the binary
go build

# Run the application
./rsvp-system
```

## Usage

Once running, you'll see a CLI prompt:

```
=== RSVP System ===
Commands: create-event, list-events, view-event, rsvp, remove-rsvp, results, quit

> create-event
Event name: Team Lunch
Event date (YYYY-MM-DD): 2024-12-20
Event description: Monthly team gathering at the cafe
Event created with ID: a7k2m9z1

> list-events
=== Events ===
[a7k2m9z1] Team Lunch - 2024-12-20

> rsvp a7k2m9z1 Alice yes
Alice has RSVP'd 'yes' to the event.

> rsvp a7k2m9z1 Bob maybe
Bob has RSVP'd 'maybe' to the event.

> rsvp a7k2m9z1 Charlie no
Charlie has RSVP'd 'no' to the event.

> view-event a7k2m9z1
=== Team Lunch ===
ID: a7k2m9z1
Date: 2024-12-20
Description: Monthly team gathering at the cafe
Total attendees: 3

Attendees:
  - Alice: yes
  - Bob: maybe
  - Charlie: no

> results a7k2m9z1
=== RSVP Results ===
Yes: 1
No: 1
Maybe: 1
Total: 3

> quit
Goodbye!
```

## Known Gaps in the Current Code

- **`DeleteEvent` isn't reachable from the CLI.** It's implemented in `store.go:122` but `main.go`'s command switch never calls it — there's no `delete-event` case. Wiring this up is a good first exercise before anything bigger.
- **`FindAttendee` (`models.go:37`) returns a pointer to a copy, not the slice element.** `for i, attendee := range e.Attendees { return i, &attendee, ... }` takes the address of the loop variable, so `&attendee` never points into `e.Attendees`. Every current caller works around this by mutating through the returned `idx` instead of the pointer (see `store.go:61` and `store.go:92`) — but the pointer is a trap for the next person who uses it directly. Fix: `return i, &e.Attendees[i], true`.
- **RSVP status validation lives only in `main.go:133`, not in the store.** `store.AddRSVP` will happily accept any string as a status. If you build a REST API or any other caller on top of `EventStore`, it bypasses that check entirely.

## Next Steps to Extend This

### 1. **Add Persistence**
- Extract a `Storage` interface with the methods `EventStore` already has: `AddEvent(*Event)`, `GetEvent(id string) (*Event, bool)`, `GetAllEvents() []*Event`, `AddRSVP(eventID, name, status string) error`, `RemoveRSVP(eventID, name string) error`, `GetRSVPResults(eventID string) (*RSVPResult, error)`, `DeleteEvent(eventID string) error`.
- Rename the current struct to `MemoryStore` and have it implement that interface — no logic changes needed, just the extraction.
- Add a `JSONStore` that loads from / saves to a file with `encoding/json`. Watch out: `Event.Date` and `Event.CreatedAt` are `time.Time` — JSON marshals them as RFC3339 strings by default, which is fine, but if you move to SQLite you'll need to convert explicitly (`time.Parse`/`.Format("2006-01-02")`, same layout `handleCreateEvent` already uses).
- For SQLite, `modernc.org/sqlite` avoids cgo. Schema: an `events` table plus an `attendees` table with `event_id` as a foreign key, mirroring the `Event.Attendees []Attendee` slice.

### 2. **Build a REST API**
- Add `json:"..."` struct tags to `Event`, `Attendee`, and `RSVPResult` in `models.go` — there are none right now, so `encoding/json` will use the Go field names (`ID`, `Name`, `Date`...) as-is, which is usually not what you want for a public API.
- Suggested routes, mapped 1:1 to existing `handle*` functions in `main.go`: `POST /events` → `handleCreateEvent` logic, `GET /events` → `handleListEvents`, `GET /events/{id}` → `handleViewEvent`, `POST /events/{id}/rsvp` → `handleRSVP`, `DELETE /events/{id}/rsvp/{name}` → `handleRemoveRSVP`, `GET /events/{id}/results` → `handleResults`, `DELETE /events/{id}` → the currently-unwired `DeleteEvent`.
- Go 1.22+'s stdlib `net/http.ServeMux` supports path parameters (`"GET /events/{id}"`) natively now, so you don't need Gin or Echo just to get this working — worth trying that first since it's one less dependency, then compare against Gin if you want middleware/validation conveniences.
- Move the status validation from `handleRSVP` (main.go:133) into the store layer so the API gets the same guarantee the CLI has.

### 3. **Add Testing**
- `store_test.go` already covers the happy paths (add/get/update/remove RSVP, results, delete, one not-found case). What's missing:
  - A concurrency test: spin up multiple goroutines calling `store.AddRSVP` on the same event and run `go test -race` to actually exercise the `sync.RWMutex`.
  - A test for invalid status strings against `store.AddRSVP` directly (bypassing `main.go`'s validation) — this will currently pass when it probably shouldn't, which is the gap noted above.
  - A regression test for the `FindAttendee` pointer bug: get a pointer via `FindAttendee`, mutate through it, then assert the change is visible in `event.Attendees` — this fails today.
  - Convert `TestGetRSVPResults`-style tests into table-driven form (`[]struct{ name, status string; want RSVPResult }`) — the repo's own README asks this as a discussion question but doesn't have an example to point to yet.

### 4. **Improve CLI**
- `cobra` maps cleanly onto the existing commands — each `case` in `main.go`'s switch (`create-event`, `list-events`, `view-event`, `rsvp`, `remove-rsvp`, `results`, `quit`) becomes a `cobra.Command`, and the `handle*` functions become the `Run` bodies with almost no changes.
- Add the missing `delete-event <event-id>` command while you're in there.
- Improve error messages: right now `store.go` returns bare `fmt.Errorf("event not found")` / `fmt.Errorf("attendee not found")` with no context — consider `fmt.Errorf("event %q not found", eventID)` so CLI output is actionable.

### 5. **Add Web Frontend**
- Only worth doing after step 2 (REST API) exists to call.
- Serve static files with `http.FileServer` from the same Go binary; have the frontend `fetch()` the routes listed above.

## Code Highlights

### Thread-Safe Store
```go
func (es *EventStore) AddRSVP(eventID, attendeeName, status string) error {
    es.mu.Lock()          // Acquire write lock
    defer es.mu.Unlock()  // Release on return
    
    event, found := es.events[eventID]
    // ... business logic
}
```

### Error Handling
```go
event, found := store.GetEvent(eventID)
if !found {
    return fmt.Errorf("event not found")  // Explicit error
}
```

### Slicing & Removal
```go
// Find attendee index
idx, _, found := event.FindAttendee(attendeeName)
if found {
    // Remove by slicing
    event.Attendees = append(event.Attendees[:idx], event.Attendees[idx+1:]...)
}
```

## Questions to Answer as You Code

1. Why use `sync.RWMutex` instead of a simple `sync.Mutex`?
2. When should you use pointers vs values for function receivers?
3. How would you make the store database-agnostic with an interface?
4. What happens if two goroutines try to RSVP simultaneously?
5. How would you test the concurrent access?

Have fun building! 🚀
