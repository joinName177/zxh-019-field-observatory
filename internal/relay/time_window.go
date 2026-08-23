package observatory

import "time"

type TimeWindow struct {
	From time.Time
	To   time.Time
}

func NewTimeWindow(from, to time.Time) TimeWindow {
	return TimeWindow{From: from.UTC(), To: to.UTC()}
}

func (w TimeWindow) Valid() bool { return !w.From.IsZero() && !w.To.IsZero() && w.From.Before(w.To) }

func (w TimeWindow) Contains(at time.Time) bool {
	if !w.Valid() {
		return false
	}
	return !at.Before(w.From) && at.Before(w.To)
}

func FilterLayersByUpdated(layers []Layer, window TimeWindow) []Layer {
	if !window.Valid() {
		return nil
	}
	out := make([]Layer, 0, len(layers))
	for _, layer := range layers {
		if window.Contains(layer.UpdatedAt) {
			out = append(out, CloneLayer(layer))
		}
	}
	return out
}

func LatestLayer(layers []Layer) (Layer, bool) {
	if len(layers) == 0 {
		return Layer{}, false
	}
	latest := CloneLayer(layers[0])
	for _, candidate := range layers[1:] {
		if candidate.UpdatedAt.After(latest.UpdatedAt) {
			latest = CloneLayer(candidate)
		}
	}
	return latest, true
}
