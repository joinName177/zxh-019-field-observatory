package observatory

import "context"

type ImportQueue struct {
	queue    *Queue
	importer *Importer
}

func NewImportQueue(i *Importer, size int) *ImportQueue {
	return &ImportQueue{queue: NewQueue(size), importer: i}
}
func (q *ImportQueue) Enqueue(ctx context.Context, wid string, l Layer) error {
	return q.queue.Submit(ctx, func(jobCtx context.Context) error { _, err := q.importer.Import(jobCtx, wid, l); return err })
}
func (q *ImportQueue) Process(ctx context.Context) error { return q.queue.RunOne(ctx) }
func (q *ImportQueue) Pending() int                      { return q.queue.Len() }
