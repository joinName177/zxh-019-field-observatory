package observatory

import "context"

type Catalog struct{ repo Repository }

func NewCatalog(r Repository) *Catalog { return &Catalog{repo: r} }
func (c *Catalog) Create(ctx context.Context, id, name string) (Workspace, error) {
	w := Workspace{ID: id, Name: NormalizeName(name), Layers: map[string]Layer{}}
	if err := ValidateWorkspace(w); err != nil {
		return Workspace{}, err
	}
	return w, c.repo.Save(ctx, w)
}
func (c *Catalog) Rename(ctx context.Context, id, name string) (Workspace, error) {
	w, err := c.repo.Load(ctx, id)
	if err != nil {
		return Workspace{}, err
	}
	w.Name = NormalizeName(name)
	w.Revision++
	return w, c.repo.Save(ctx, w)
}
func (c *Catalog) Remove(ctx context.Context, id string) error { return c.repo.Delete(ctx, id) }
