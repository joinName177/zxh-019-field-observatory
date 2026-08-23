package observatory

import "context"

type Backup struct {
	source Repository
	target Repository
}

func NewBackup(s, t Repository) *Backup { return &Backup{source: s, target: t} }
func (b *Backup) Copy(ctx context.Context, id string) error {
	w, err := b.source.Load(ctx, id)
	if err != nil {
		return err
	}
	return b.target.Save(ctx, CloneWorkspace(w))
}
func (b *Backup) Restore(ctx context.Context, id string) error { return b.Copy(ctx, id) }
