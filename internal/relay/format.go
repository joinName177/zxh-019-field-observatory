package observatory

import "fmt"

func PointString(p Point) string { return fmt.Sprintf("%.3f,%.3f", p.X, p.Y) }
func LayerLabel(l Layer) string  { return fmt.Sprintf("%s (%d points)", l.Name, len(l.Points)) }
func WorkspaceLabel(w Workspace) string {
	return fmt.Sprintf("%s: %d layers, revision %d", w.Name, len(w.Layers), w.Revision)
}
func BoundsString(l Layer) string {
	a, b := l.Bounds()
	return fmt.Sprintf("[%s]-[%s]", PointString(a), PointString(b))
}
