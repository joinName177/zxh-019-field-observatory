package observatory

import "context"

func Duplicate(ctx context.Context, r Repository, wid, lid, newID string) (Layer, error) {
	w, err := r.Load(ctx, wid)
	if err != nil {
		return Layer{}, err
	}
	l, ok := w.Layers[lid]
	if !ok {
		return Layer{}, ErrLayerNotFound
	}
	l = CloneLayer(l)
	l.ID = newID
	w.Layers[newID] = l
	w.Revision++
	return l, r.Save(ctx, w)
}
func Tag(ctx context.Context, r Repository, wid, lid, tag string) error {
	w, err := r.Load(ctx, wid)
	if err != nil {
		return err
	}
	l, ok := w.Layers[lid]
	if !ok {
		return ErrLayerNotFound
	}
	l.Tags = MergeTags(l.Tags, []string{tag})
	w.Layers[lid] = l
	w.Revision++
	return r.Save(ctx, w)
}
func Untag(ctx context.Context, r Repository, wid, lid, tag string) error {
	w, err := r.Load(ctx, wid)
	if err != nil {
		return err
	}
	l, ok := w.Layers[lid]
	if !ok {
		return ErrLayerNotFound
	}
	out := l.Tags[:0]
	for _, x := range l.Tags {
		if x != tag {
			out = append(out, x)
		}
	}
	l.Tags = out
	w.Layers[lid] = l
	w.Revision++
	return r.Save(ctx, w)
}
func Touch(ctx context.Context, r Repository, wid, lid string) error {
	w, err := r.Load(ctx, wid)
	if err != nil {
		return err
	}
	l, ok := w.Layers[lid]
	if !ok {
		return ErrLayerNotFound
	}
	l.UpdatedAt = SystemClock{}.Now()
	w.Layers[lid] = l
	w.Revision++
	return r.Save(ctx, w)
}
