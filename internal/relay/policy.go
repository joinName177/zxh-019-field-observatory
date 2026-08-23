package observatory

import "context"

type Policy struct{ MaxLayers, MaxPoints int }

func DefaultPolicy() Policy { return Policy{MaxLayers: 1000, MaxPoints: 100000} }
func (p Policy) Check(w Workspace) error {
	if p.MaxLayers > 0 && len(w.Layers) > p.MaxLayers {
		return ErrInvalidLayer
	}
	points := 0
	for _, l := range w.Layers {
		points += len(l.Points)
	}
	if p.MaxPoints > 0 && points > p.MaxPoints {
		return ErrInvalidLayer
	}
	return nil
}
func ApplyPolicy(ctx context.Context, p Policy, r Repository, id string) error {
	w, err := r.Load(ctx, id)
	if err != nil {
		return err
	}
	return p.Check(w)
}
