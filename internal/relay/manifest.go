package observatory

import "time"

type Manifest struct {
	WorkspaceID string
	Revision    int
	GeneratedAt time.Time
	LayerIDs    []string
}

func BuildManifest(w Workspace) Manifest {
	ids := make([]string, 0, len(w.Layers))
	for id := range w.Layers {
		ids = append(ids, id)
	}
	return Manifest{WorkspaceID: w.ID, Revision: w.Revision, GeneratedAt: time.Now().UTC(), LayerIDs: ids}
}
func (m Manifest) Has(id string) bool {
	for _, x := range m.LayerIDs {
		if x == id {
			return true
		}
	}
	return false
}
func (m Manifest) Count() int { return len(m.LayerIDs) }
