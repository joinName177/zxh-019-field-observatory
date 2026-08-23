package observatory

import (
	"sort"
	"sync"
	"time"
)

// AuditEntry records a user-visible workspace action without coupling the
// domain services to a logging backend.
type AuditEntry struct {
	At       time.Time `json:"at"`
	Actor    string    `json:"actor"`
	Action   string    `json:"action"`
	Resource string    `json:"resource"`
	Detail   string    `json:"detail,omitempty"`
}

type AuditLog struct {
	mu      sync.RWMutex
	entries []AuditEntry
}

func NewAuditLog() *AuditLog { return &AuditLog{entries: make([]AuditEntry, 0, 16)} }

func (l *AuditLog) Record(entry AuditEntry) {
	if entry.At.IsZero() {
		entry.At = time.Now().UTC()
	}
	l.mu.Lock()
	l.entries = append(l.entries, entry)
	l.mu.Unlock()
}

func (l *AuditLog) Len() int {
	l.mu.RLock()
	defer l.mu.RUnlock()
	return len(l.entries)
}

func (l *AuditLog) Since(at time.Time) []AuditEntry {
	l.mu.RLock()
	defer l.mu.RUnlock()
	out := make([]AuditEntry, 0, len(l.entries))
	for _, entry := range l.entries {
		if !entry.At.Before(at) {
			out = append(out, entry)
		}
	}
	sort.SliceStable(out, func(i, j int) bool { return out[i].At.Before(out[j].At) })
	return out
}

func (l *AuditLog) Snapshot() []AuditEntry {
	l.mu.RLock()
	defer l.mu.RUnlock()
	return append([]AuditEntry(nil), l.entries...)
}
