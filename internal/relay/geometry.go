package observatory

import "math"

func Distance(a, b Point) float64 { return math.Hypot(a.X-b.X, a.Y-b.Y) }
func Length(l Layer) float64 {
	total := 0.0
	for i := 1; i < len(l.Points); i++ {
		total += Distance(l.Points[i-1], l.Points[i])
	}
	return total
}
func Simplify(l Layer, tolerance float64) Layer {
	if tolerance <= 0 || len(l.Points) < 3 {
		return CloneLayer(l)
	}
	out := Layer{ID: l.ID, Name: l.Name, Tags: cloneStrings(l.Tags), UpdatedAt: l.UpdatedAt}
	out.Points = append(out.Points, l.Points[0])
	for i := 1; i < len(l.Points)-1; i++ {
		if Distance(l.Points[i], out.Points[len(out.Points)-1]) >= tolerance {
			out.Points = append(out.Points, l.Points[i])
		}
	}
	out.Points = append(out.Points, l.Points[len(l.Points)-1])
	return out
}
func Centroid(l Layer) Point {
	if len(l.Points) == 0 {
		return Point{}
	}
	sum := Point{}
	for _, p := range l.Points {
		sum.X += p.X
		sum.Y += p.Y
	}
	n := float64(len(l.Points))
	return Point{sum.X / n, sum.Y / n}
}
