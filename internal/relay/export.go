package observatory

import (
	"context"
	"encoding/json"
)

type Exporter struct{ query *QueryService }

func NewExporter(q *QueryService) *Exporter { return &Exporter{query: q} }
func (e *Exporter) JSON(ctx context.Context, wid string) ([]byte, error) {
	w, err := e.query.Workspace(ctx, wid)
	if err != nil {
		return nil, err
	}
	return EncodeWorkspace(w)
}
func (e *Exporter) Summary(ctx context.Context, wid string) (map[string]any, error) {
	ls, err := e.query.List(ctx, wid)
	if err != nil {
		return nil, err
	}
	out := map[string]any{"workspace": wid, "layers": len(ls), "length": 0.0}
	total := 0.0
	for _, l := range ls {
		total += Length(l)
	}
	out["length"] = total
	return out, nil
}
func MarshalSummary(v map[string]any) ([]byte, error) { return json.Marshal(v) }
