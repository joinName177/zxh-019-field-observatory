package observatory

type Viewport struct {
	Center              Point
	Width, Height, Zoom float64
}

func DefaultViewport() Viewport { return Viewport{Width: 800, Height: 600, Zoom: 1} }
func (v Viewport) Visible(p Point) bool {
	halfW := v.Width / (2 * v.Zoom)
	halfH := v.Height / (2 * v.Zoom)
	return p.X >= v.Center.X-halfW && p.X <= v.Center.X+halfW && p.Y >= v.Center.Y-halfH && p.Y <= v.Center.Y+halfH
}
func (v Viewport) Pan(dx, dy float64) Viewport { v.Center.X += dx; v.Center.Y += dy; return v }
func (v Viewport) ZoomBy(f float64) Viewport {
	if f > 0 {
		v.Zoom *= f
	}
	return v
}
func (v Viewport) Fit(min, max Point) Viewport {
	v.Center = Point{(min.X + max.X) / 2, (min.Y + max.Y) / 2}
	v.Width = max.X - min.X
	v.Height = max.Y - min.Y
	if v.Width <= 0 {
		v.Width = 1
	}
	if v.Height <= 0 {
		v.Height = 1
	}
	return v
}
