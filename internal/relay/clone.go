package observatory

func CloneLayer(l Layer) Layer {
	l.Points = clonePoints(l.Points)
	l.Tags = cloneStrings(l.Tags)
	return l
}
func CloneWorkspace(w Workspace) Workspace {
	out := Workspace{ID: w.ID, Name: w.Name, Revision: w.Revision, Layers: map[string]Layer{}}
	for id, l := range w.Layers {
		out.Layers[id] = CloneLayer(l)
	}
	return out
}
func MergeTags(a, b []string) []string {
	out := cloneStrings(a)
	seen := map[string]bool{}
	for _, x := range out {
		seen[x] = true
	}
	for _, x := range b {
		if !seen[x] {
			out = append(out, x)
			seen[x] = true
		}
	}
	return out
}
func AddPoint(l Layer, p Point) Layer { l = CloneLayer(l); l.Points = append(l.Points, p); return l }
func RenameLayer(l Layer, name string) Layer {
	l = CloneLayer(l)
	l.Name = NormalizeName(name)
	return l
}
