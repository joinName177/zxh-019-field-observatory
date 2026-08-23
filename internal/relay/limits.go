package observatory

type Limits struct {
	MaxName          int
	MaxTags          int
	MaxPointDistance float64
}

func DefaultLimits() Limits { return Limits{MaxName: 120, MaxTags: 32, MaxPointDistance: 1000000} }
func (l Limits) CheckLayer(v Layer) error {
	if l.MaxName > 0 && len(v.Name) > l.MaxName {
		return ErrInvalidLayer
	}
	if l.MaxTags > 0 && len(v.Tags) > l.MaxTags {
		return ErrInvalidLayer
	}
	for i := 1; i < len(v.Points); i++ {
		if l.MaxPointDistance > 0 && Distance(v.Points[i-1], v.Points[i]) > l.MaxPointDistance {
			return ErrInvalidLayer
		}
	}
	return nil
}
