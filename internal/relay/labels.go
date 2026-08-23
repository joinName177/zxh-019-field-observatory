package observatory

type Label struct {
	Text string
	At   Point
	Size float64
}

func MakeLabel(l Layer) Label                  { return Label{Text: l.Name, At: Centroid(l), Size: 12} }
func LabelForPoint(p Point, text string) Label { return Label{Text: text, At: p, Size: 12} }
func ShiftLabel(l Label, dx, dy float64) Label { l.At.X += dx; l.At.Y += dy; return l }
func ResizeLabel(l Label, size float64) Label {
	if size > 0 {
		l.Size = size
	}
	return l
}
