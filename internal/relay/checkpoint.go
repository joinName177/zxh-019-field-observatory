package observatory

import "time"

type Checkpoint struct {
	Revision  int
	At        time.Time
	Workspace Workspace
}
type Checkpoints struct{ items []Checkpoint }

func NewCheckpoints() *Checkpoints { return &Checkpoints{items: []Checkpoint{}} }
func (c *Checkpoints) Save(w Workspace) {
	c.items = append(c.items, Checkpoint{w.Revision, time.Now().UTC(), CloneWorkspace(w)})
}
func (c *Checkpoints) Latest() (Checkpoint, bool) {
	if len(c.items) == 0 {
		return Checkpoint{}, false
	}
	return c.items[len(c.items)-1], true
}
func (c *Checkpoints) At(revision int) (Checkpoint, bool) {
	for i := len(c.items) - 1; i >= 0; i-- {
		if c.items[i].Revision == revision {
			return c.items[i], true
		}
	}
	return Checkpoint{}, false
}
func (c *Checkpoints) Count() int { return len(c.items) }
