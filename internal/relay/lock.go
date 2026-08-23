package observatory

import (
	"context"
	"sync"
)

type Lease struct{ mu sync.Mutex }

func (l *Lease) Acquire(ctx context.Context) error {
	select {
	case <-ctx.Done():
		return ctx.Err()
	default:
	}
	l.mu.Lock()
	return nil
}
func (l *Lease) Release() { l.mu.Unlock() }
func (l *Lease) With(ctx context.Context, fn func() error) error {
	if err := l.Acquire(ctx); err != nil {
		return err
	}
	defer l.Release()
	return fn()
}
