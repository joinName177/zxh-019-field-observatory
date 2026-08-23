package observatory

import "fmt"

func MergeLayers(base, overlay Layer) (Layer, error) {
	if base.ID == "" || overlay.ID == "" {
		return Layer{}, ErrInvalidLayer
	}
	if base.ID != overlay.ID {
		return Layer{}, fmt.Errorf("layer ids differ")
	}
	out := CloneLayer(base)
	out.Points = append(out.Points, clonePoints(overlay.Points)...)
	out.Tags = MergeTags(out.Tags, overlay.Tags)
	if overlay.Name != "" {
		out.Name = overlay.Name
	}
	return out, nil
}
func SplitLayer(l Layer, size int) []Layer {
	if size < 1 {
		size = 1
	}
	out := []Layer{}
	for start := 0; start < len(l.Points); start += size {
		end := start + size
		if end > len(l.Points) {
			end = len(l.Points)
		}
		part := CloneLayer(l)
		part.ID = fmt.Sprintf("%s-%d", l.ID, len(out)+1)
		part.Points = clonePoints(l.Points[start:end])
		out = append(out, part)
	}
	return out
}
