package observatory

import (
	"bytes"
	"encoding/json"
)

func CloneJSON(v any) (any, error) {
	data, err := json.Marshal(v)
	if err != nil {
		return nil, err
	}
	var out any
	err = json.NewDecoder(bytes.NewReader(data)).Decode(&out)
	return out, err
}
func LayerMap(layers []Layer) map[string]Layer {
	out := map[string]Layer{}
	for _, l := range layers {
		out[l.ID] = CloneLayer(l)
	}
	return out
}
func LayerValues(m map[string]Layer) []Layer {
	out := make([]Layer, 0, len(m))
	for _, l := range m {
		out = append(out, CloneLayer(l))
	}
	return out
}
func WorkspaceFromLayers(id, name string, layers []Layer) Workspace {
	return Workspace{ID: id, Name: name, Layers: LayerMap(layers)}
}
