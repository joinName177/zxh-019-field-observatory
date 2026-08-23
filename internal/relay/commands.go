package observatory

import "context"

type Command interface {
	Execute(context.Context) error
	Undo(context.Context) error
}
type RenameCommand struct {
	repo                       Repository
	wid, lid, oldName, newName string
}

func NewRenameCommand(r Repository, wid, lid, name string) (*RenameCommand, error) {
	w, err := r.Load(context.Background(), wid)
	if err != nil {
		return nil, err
	}
	l, ok := w.Layers[lid]
	if !ok {
		return nil, ErrLayerNotFound
	}
	return &RenameCommand{repo: r, wid: wid, lid: lid, oldName: l.Name, newName: name}, nil
}
func (c *RenameCommand) Execute(ctx context.Context) error {
	_, err := RenameLayerInWorkspace(ctx, c.repo, c.wid, c.lid, c.newName)
	return err
}
func (c *RenameCommand) Undo(ctx context.Context) error {
	_, err := RenameLayerInWorkspace(ctx, c.repo, c.wid, c.lid, c.oldName)
	return err
}
