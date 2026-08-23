package observatory

import "context"

type Closer interface{ Close() error }
type Cleanup struct{ items []Closer }

func NewCleanup() *Cleanup      { return &Cleanup{items: []Closer{}} }
func (c *Cleanup) Add(v Closer) { c.items = append(c.items, v) }
func (c *Cleanup) Close() error {
	var first error
	for i := len(c.items) - 1; i >= 0; i-- {
		if err := c.items[i].Close(); err != nil && first == nil {
			first = err
		}
	}
	return first
}
func CloseContext(ctx context.Context, c Closer) error {
	select {
	case <-ctx.Done():
		return ctx.Err()
	default:
		return c.Close()
	}
}
