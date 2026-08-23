package observatory

import "sync"

type Notification struct{ Kind, Message string }
type Notifier struct {
	mu    sync.Mutex
	items []Notification
}

func NewNotifier() *Notifier { return &Notifier{items: []Notification{}} }
func (n *Notifier) Send(v Notification) {
	n.mu.Lock()
	defer n.mu.Unlock()
	n.items = append(n.items, v)
}
func (n *Notifier) Drain() []Notification {
	n.mu.Lock()
	defer n.mu.Unlock()
	out := append([]Notification(nil), n.items...)
	n.items = nil
	return out
}
func (n *Notifier) Pending() int { n.mu.Lock(); defer n.mu.Unlock(); return len(n.items) }
