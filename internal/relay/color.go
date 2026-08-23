package observatory

type Color struct{ R, G, B, A uint8 }

func RGB(r, g, b uint8) Color           { return Color{r, g, b, 255} }
func Transparent() Color                { return Color{} }
func (c Color) Opaque() bool            { return c.A > 0 }
func (c Color) WithAlpha(a uint8) Color { c.A = a; return c }
func (c Color) Mix(o Color, ratio float64) Color {
	if ratio < 0 {
		ratio = 0
	}
	if ratio > 1 {
		ratio = 1
	}
	return Color{uint8(float64(c.R)*(1-ratio) + float64(o.R)*ratio), uint8(float64(c.G)*(1-ratio) + float64(o.G)*ratio), uint8(float64(c.B)*(1-ratio) + float64(o.B)*ratio), uint8(float64(c.A)*(1-ratio) + float64(o.A)*ratio)}
}
