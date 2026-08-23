package observatory

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"time"
)

type ImportManifest struct {
	LayerID    string    `json:"layer_id"`
	Source     string    `json:"source"`
	Checksum   string    `json:"checksum"`
	PointCount int       `json:"point_count"`
	ImportedAt time.Time `json:"imported_at"`
}

func NewImportManifest(source string, layer Layer) (ImportManifest, error) {
	if !layer.Valid() {
		return ImportManifest{}, fmt.Errorf("layer %q is invalid", layer.ID)
	}
	hash := sha256.New()
	for _, point := range layer.Points {
		fmt.Fprintf(hash, "%.6f,%.6f;", point.X, point.Y)
	}
	return ImportManifest{LayerID: layer.ID, Source: source, Checksum: hex.EncodeToString(hash.Sum(nil)), PointCount: len(layer.Points), ImportedAt: time.Now().UTC()}, nil
}

func (m ImportManifest) Matches(layer Layer) bool {
	if m.LayerID != layer.ID || m.PointCount != len(layer.Points) {
		return false
	}
	other, err := NewImportManifest(m.Source, layer)
	return err == nil && other.Checksum == m.Checksum
}
