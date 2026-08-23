package observatory

import (
	"sync"
	"time"
)

type EventLog struct {
	mu     sync.Mutex
	events []Change
}

func NewEventLog() *EventLog { return &EventLog{events: []Change{}} }
func (l *EventLog) Append(c Change) {
	l.mu.Lock()
	defer l.mu.Unlock()
	if c.At.IsZero() {
		c.At = time.Now().UTC()
	}
	l.events = append(l.events, c)
}
func (l *EventLog) Recent(n int) []Change {
	l.mu.Lock()
	defer l.mu.Unlock()
	if n > len(l.events) {
		n = len(l.events)
	}
	out := make([]Change, n)
	copy(out, l.events[len(l.events)-n:])
	return out
}
func (l *EventLog) Count() int { l.mu.Lock(); defer l.mu.Unlock(); return len(l.events) }
