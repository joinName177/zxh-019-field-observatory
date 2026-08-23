package observatory

import "time"

type Point struct{ X, Y float64 }
type Layer struct {
	ID, Name  string
	Points    []Point
	Tags      []string
	UpdatedAt time.Time
}
type Workspace struct {
	ID, Name string
	Layers   map[string]Layer
	Revision int
}
type Change struct {
	WorkspaceID, LayerID, Actor string
	Kind                        string
	At                          time.Time
}

func NewLayer(id, name string, points []Point, tags []string) Layer {
	return Layer{ID: id, Name: name, Points: clonePoints(points), Tags: cloneStrings(tags), UpdatedAt: time.Now().UTC()}
}
func (l Layer) Valid() bool { return l.ID != "" && l.Name != "" && len(l.Points) > 1 }
func (l Layer) Bounds() (Point, Point) {
	if len(l.Points) == 0 {
		return Point{}, Point{}
	}
	min, max := l.Points[0], l.Points[0]
	for _, p := range l.Points[1:] {
		if p.X < min.X {
			min.X = p.X
		}
		if p.Y < min.Y {
			min.Y = p.Y
		}
		if p.X > max.X {
			max.X = p.X
		}
		if p.Y > max.Y {
			max.Y = p.Y
		}
	}
	return min, max
}
func clonePoints(in []Point) []Point    { out := make([]Point, len(in)); copy(out, in); return out }
func cloneStrings(in []string) []string { out := make([]string, len(in)); copy(out, in); return out }
