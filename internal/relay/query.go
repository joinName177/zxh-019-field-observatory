package observatory

import "context"

type QueryService struct{ repo Repository }

func NewQueryService(repo Repository) *QueryService { return &QueryService{repo: repo} }
func (q *QueryService) Workspace(ctx context.Context, id string) (Workspace, error) {
	return q.repo.Load(ctx, id)
}
func (q *QueryService) Layer(ctx context.Context, wid, lid string) (Layer, error) {
	w, err := q.repo.Load(ctx, wid)
	if err != nil {
		return Layer{}, err
	}
	l, ok := w.Layers[lid]
	if !ok {
		return Layer{}, ErrLayerNotFound
	}
	return CloneLayer(l), nil
}
func (q *QueryService) List(ctx context.Context, wid string) ([]Layer, error) {
	w, err := q.repo.Load(ctx, wid)
	if err != nil {
		return nil, err
	}
	out := make([]Layer, 0, len(w.Layers))
	for _, l := range w.Layers {
		out = append(out, CloneLayer(l))
	}
	return out, nil
}
