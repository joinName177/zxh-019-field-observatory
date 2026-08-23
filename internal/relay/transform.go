package observatory

import "math"

type Transform struct{ Scale, Rotate, OffsetX, OffsetY float64 }

func IdentityTransform() Transform { return Transform{Scale: 1} }
func (t Transform) Apply(p Point) Point {
	s := t.Scale
	if s == 0 {
		s = 1
	}
	c := math.Cos(t.Rotate)
	si := math.Sin(t.Rotate)
	return Point{X: s*(p.X*c-p.Y*si) + t.OffsetX, Y: s*(p.X*si+p.Y*c) + t.OffsetY}
}
func (t Transform) ApplyLayer(l Layer) Layer {
	out := CloneLayer(l)
	for i, p := range out.Points {
		out.Points[i] = t.Apply(p)
	}
	return out
}
func (t Transform) Inverse(p Point) Point {
	s := t.Scale
	if s == 0 {
		s = 1
	}
	x := (p.X - t.OffsetX) / s
	y := (p.Y - t.OffsetY) / s
	c := math.Cos(-t.Rotate)
	si := math.Sin(-t.Rotate)
	return Point{x*c - y*si, x*si + y*c}
}
func Translate(x, y float64) Transform { return Transform{Scale: 1, OffsetX: x, OffsetY: y} }
func ScaleBy(v float64) Transform {
	if v == 0 {
		v = 1
	}
	return Transform{Scale: v}
}
func RotateBy(v float64) Transform { return Transform{Scale: 1, Rotate: v} }
