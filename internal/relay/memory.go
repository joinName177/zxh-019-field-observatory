package observatory

import "context"

type MemoryRepository struct{ items map[string]Workspace }

func NewMemoryRepository() *MemoryRepository { return &MemoryRepository{items: map[string]Workspace{}} }
func (r *MemoryRepository) Load(ctx context.Context, id string) (Workspace, error) {
	select {
	case <-ctx.Done():
		return Workspace{}, ctx.Err()
	default:
	}
	w, ok := r.items[id]
	if !ok {
		return Workspace{}, ErrWorkspaceNotFound
	}
	return CloneWorkspace(w), nil
}
func (r *MemoryRepository) Save(ctx context.Context, w Workspace) error {
	select {
	case <-ctx.Done():
		return ctx.Err()
	default:
	}
	r.items[w.ID] = CloneWorkspace(w)
	return nil
}
func (r *MemoryRepository) Delete(ctx context.Context, id string) error {
	select {
	case <-ctx.Done():
		return ctx.Err()
	default:
	}
	delete(r.items, id)
	return nil
}
