package observatory

import "context"

type RenderFrame struct {
	WorkspaceID          string
	Layers               []Layer
	BoundsMin, BoundsMax Point
}
type Renderer struct{ query *QueryService }

func NewRenderer(q *QueryService) *Renderer { return &Renderer{query: q} }
func (r *Renderer) Frame(ctx context.Context, wid string) (RenderFrame, error) {
	layers, err := r.query.List(ctx, wid)
	if err != nil {
		return RenderFrame{}, err
	}
	f := RenderFrame{WorkspaceID: wid, Layers: layers}
	first := true
	for _, l := range layers {
		a, b := l.Bounds()
		if first {
			f.BoundsMin, f.BoundsMax = a, b
			first = false
		}
		if a.X < f.BoundsMin.X {
			f.BoundsMin.X = a.X
		}
		if a.Y < f.BoundsMin.Y {
			f.BoundsMin.Y = a.Y
		}
		if b.X > f.BoundsMax.X {
			f.BoundsMax.X = b.X
		}
		if b.Y > f.BoundsMax.Y {
			f.BoundsMax.Y = b.Y
		}
	}
	return f, nil
}
func (r *Renderer) HitTest(ctx context.Context, wid string, p Point) (string, error) {
	layers, err := r.query.List(ctx, wid)
	if err != nil {
		return "", err
	}
	for _, l := range layers {
		a, b := l.Bounds()
		if p.X >= a.X && p.X <= b.X && p.Y >= a.Y && p.Y <= b.Y {
			return l.ID, nil
		}
	}
	return "", ErrLayerNotFound
}
