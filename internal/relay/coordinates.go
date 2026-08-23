package observatory

type CoordinateSystem string

const (
	Geo   CoordinateSystem = "geo"
	Local CoordinateSystem = "local"
)

func Convert(p Point, from, to CoordinateSystem) Point {
	if from == to {
		return p
	}
	if from == Geo && to == Local {
		return Point{p.X * 111000, p.Y * 111000}
	}
	return Point{p.X / 111000, p.Y / 111000}
}
func ConvertLayer(l Layer, from, to CoordinateSystem) Layer {
	out := CloneLayer(l)
	for i, p := range out.Points {
		out.Points[i] = Convert(p, from, to)
	}
	return out
}
func ValidCoordinateSystem(v CoordinateSystem) bool { return v == Geo || v == Local }
func CoordinateSystems() []CoordinateSystem         { return []CoordinateSystem{Geo, Local} }
func IsOrigin(p Point) bool                         { return p.X == 0 && p.Y == 0 }
func Clamp(p Point, min, max Point) Point {
	if p.X < min.X {
		p.X = min.X
	}
	if p.Y < min.Y {
		p.Y = min.Y
	}
	if p.X > max.X {
		p.X = max.X
	}
	if p.Y > max.Y {
		p.Y = max.Y
	}
	return p
}
func Midpoint(a, b Point) Point           { return Point{(a.X + b.X) / 2, (a.Y + b.Y) / 2} }
func Equal(a, b Point) bool               { return a.X == b.X && a.Y == b.Y }
func Add(a, b Point) Point                { return Point{a.X + b.X, a.Y + b.Y} }
func Sub(a, b Point) Point                { return Point{a.X - b.X, a.Y - b.Y} }
func Magnitude(p Point) float64           { return Distance(Point{}, p) }
func Dot(a, b Point) float64              { return a.X*b.X + a.Y*b.Y }
func ScalePoint(p Point, v float64) Point { return Point{p.X * v, p.Y * v} }
func Negate(p Point) Point                { return Point{-p.X, -p.Y} }
func DistanceSquared(a, b Point) float64  { d := Sub(a, b); return Dot(d, d) }
func Lerp(a, b Point, t float64) Point    { return Add(a, ScalePoint(Sub(b, a), t)) }
func ValidPoint(p Point) bool             { return p.X == p.X && p.Y == p.Y }
func PointsValid(ps []Point) bool {
	for _, p := range ps {
		if !ValidPoint(p) {
			return false
		}
	}
	return true
}
func PointCount(l Layer) int { return len(l.Points) }
