package observatory

import "context"

type Job func(context.Context) error
type Queue struct{ jobs chan Job }

func NewQueue(size int) *Queue {
	if size < 1 {
		size = 1
	}
	return &Queue{jobs: make(chan Job, size)}
}
func (q *Queue) Submit(ctx context.Context, j Job) error {
	select {
	case q.jobs <- j:
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}
func (q *Queue) RunOne(ctx context.Context) error {
	select {
	case j := <-q.jobs:
		return j(ctx)
	case <-ctx.Done():
		return ctx.Err()
	}
}
func (q *Queue) Len() int { return len(q.jobs) }
