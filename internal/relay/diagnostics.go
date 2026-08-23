package observatory

import "context"

type Diagnostic struct {
	Code, Message string
	LayerID       string
}

func Diagnose(ctx context.Context, r Repository, id string) []Diagnostic {
	w, err := r.Load(ctx, id)
	if err != nil {
		return []Diagnostic{{Code: "workspace", Message: err.Error()}}
	}
	out := []Diagnostic{}
	for lid, l := range w.Layers {
		if err := ValidateLayer(l); err != nil {
			out = append(out, Diagnostic{Code: "layer", Message: err.Error(), LayerID: lid})
		}
		if !TagsValid(l.Tags) {
			out = append(out, Diagnostic{Code: "tags", Message: "duplicate or empty tags", LayerID: lid})
		}
	}
	return out
}
func Healthy(ds []Diagnostic) bool { return len(ds) == 0 }
