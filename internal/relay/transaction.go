package observatory

import "context"

type Transaction struct {
	repo      Repository
	workspace Workspace
	done      bool
}

func Begin(ctx context.Context, r Repository, id string) (*Transaction, error) {
	w, err := r.Load(ctx, id)
	if err != nil {
		return nil, err
	}
	return &Transaction{repo: r, workspace: CloneWorkspace(w)}, nil
}
func (t *Transaction) Apply(fn func(*Workspace) error) error {
	if t.done {
		return ErrCancelled
	}
	return fn(&t.workspace)
}
func (t *Transaction) Commit(ctx context.Context) error {
	if t.done {
		return ErrCancelled
	}
	t.done = true
	return t.repo.Save(ctx, t.workspace)
}
func (t *Transaction) Rollback() { t.done = true }
