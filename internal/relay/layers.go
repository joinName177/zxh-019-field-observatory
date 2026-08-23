package observatory

import "context"

func PutLayer(ctx context.Context, r Repository, wid string, l Layer) error {
	w, err := r.Load(ctx, wid)
	if err != nil {
		return err
	}
	if err := ValidateLayer(l); err != nil {
		return err
	}
	w.Layers[l.ID] = CloneLayer(l)
	w.Revision++
	return r.Save(ctx, w)
}
func RemoveLayer(ctx context.Context, r Repository, wid, lid string) error {
	w, err := r.Load(ctx, wid)
	if err != nil {
		return err
	}
	if _, ok := w.Layers[lid]; !ok {
		return ErrLayerNotFound
	}
	delete(w.Layers, lid)
	w.Revision++
	return r.Save(ctx, w)
}
func RenameLayerInWorkspace(ctx context.Context, r Repository, wid, lid, name string) (Layer, error) {
	w, err := r.Load(ctx, wid)
	if err != nil {
		return Layer{}, err
	}
	l, ok := w.Layers[lid]
	if !ok {
		return Layer{}, ErrLayerNotFound
	}
	l = RenameLayer(l, name)
	w.Layers[lid] = l
	w.Revision++
	return l, r.Save(ctx, w)
}
