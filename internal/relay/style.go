package observatory

type Style struct {
	Stroke, Fill Color
	Width        float64
	Visible      bool
}

func DefaultStyle() Style {
	return Style{Stroke: RGB(30, 90, 180), Fill: Transparent(), Width: 2, Visible: true}
}
func (s Style) Hidden() bool { return !s.Visible }
func (s Style) Thicker(v float64) Style {
	if v > 0 {
		s.Width = v
	}
	return s
}
func (s Style) Hide() Style { s.Visible = false; return s }
func (s Style) Show() Style { s.Visible = true; return s }
