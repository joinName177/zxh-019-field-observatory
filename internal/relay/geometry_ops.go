package observatory

func Reverse(l Layer) Layer {
	out := CloneLayer(l)
	for i, j := 0, len(out.Points)-1; i < j; i, j = i+1, j-1 {
		out.Points[i], out.Points[j] = out.Points[j], out.Points[i]
	}
	return out
}
func TranslateLayer(l Layer, x, y float64) Layer { return Translate(x, y).ApplyLayer(l) }
func ScaleLayer(l Layer, v float64) Layer        { return ScaleBy(v).ApplyLayer(l) }
func RotateLayer(l Layer, v float64) Layer       { return RotateBy(v).ApplyLayer(l) }
func Closed(l Layer) bool                        { return len(l.Points) > 2 && l.Points[0] == l.Points[len(l.Points)-1] }
func Close(l Layer) Layer {
	out := CloneLayer(l)
	if len(out.Points) > 0 && !Closed(out) {
		out.Points = append(out.Points, out.Points[0])
	}
	return out
}
func RemoveDuplicatePoints(l Layer) Layer {
	out := CloneLayer(l)
	if len(out.Points) < 2 {
		return out
	}
	dst := out.Points[:1]
	for _, p := range out.Points[1:] {
		if p != dst[len(dst)-1] {
			dst = append(dst, p)
		}
	}
	out.Points = dst
	return out
}
