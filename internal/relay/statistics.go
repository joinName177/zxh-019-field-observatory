package observatory

type Statistics struct {
	Layers, Points, Tags int
	TotalLength          float64
}

func ComputeStatistics(layers []Layer) Statistics {
	s := Statistics{}
	for _, l := range layers {
		s.Layers++
		s.Points += len(l.Points)
		s.Tags += len(l.Tags)
		s.TotalLength += Length(l)
	}
	return s
}
func (s Statistics) AveragePoints() float64 {
	if s.Layers == 0 {
		return 0
	}
	return float64(s.Points) / float64(s.Layers)
}
func (s Statistics) Empty() bool { return s.Layers == 0 }
