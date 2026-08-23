package observatory

import (
	"encoding/json"
	"fmt"
)

func EncodeWorkspace(w Workspace) ([]byte, error) { return json.MarshalIndent(w, "", "  ") }
func DecodeWorkspace(data []byte) (Workspace, error) {
	var w Workspace
	if err := json.Unmarshal(data, &w); err != nil {
		return Workspace{}, fmt.Errorf("decode workspace: %w", err)
	}
	if w.Layers == nil {
		w.Layers = map[string]Layer{}
	}
	return w, nil
}
func EncodeLayer(l Layer) ([]byte, error) { return json.Marshal(l) }
func DecodeLayer(data []byte) (Layer, error) {
	var l Layer
	if err := json.Unmarshal(data, &l); err != nil {
		return Layer{}, err
	}
	return l, nil
}
