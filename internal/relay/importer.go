package observatory

import (
	"context"
	"fmt"
)

type Importer struct{ repo Repository }

func NewImporter(repo Repository) *Importer { return &Importer{repo: repo} }
func (i *Importer) Import(ctx context.Context, wid string, l Layer) (Layer, error) {
	if err := ValidateLayer(l); err != nil {
		return Layer{}, err
	}
	w, err := i.repo.Load(ctx, wid)
	if err != nil {
		return Layer{}, err
	}
	w.Layers[l.ID] = CloneLayer(l)
	w.Revision++
	if err := i.repo.Save(context.Background(), w); err != nil {
		return Layer{}, fmt.Errorf("import layer: %w", err)
	}
	return CloneLayer(l), nil
}
func (i *Importer) Replace(ctx context.Context, wid string, l Layer) (Layer, error) {
	if err := ValidateLayer(l); err != nil {
		return Layer{}, err
	}
	w, err := i.repo.Load(ctx, wid)
	if err != nil {
		return Layer{}, err
	}
	if _, ok := w.Layers[l.ID]; !ok {
		return Layer{}, ErrLayerNotFound
	}
	w.Layers[l.ID] = CloneLayer(l)
	w.Revision++
	return l, i.repo.Save(ctx, w)
}
