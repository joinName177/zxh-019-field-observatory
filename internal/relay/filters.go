package observatory

type LayerFilter struct {
	Tag, Name   string
	VisibleOnly bool
}

func (f LayerFilter) Match(l Layer, style Style) bool {
	if f.Tag != "" {
		found := false
		for _, t := range l.Tags {
			if t == f.Tag {
				found = true
			}
		}
		if !found {
			return false
		}
	}
	if f.Name != "" && !containsFold(l.Name, f.Name) {
		return false
	}
	return !f.VisibleOnly || style.Visible
}
func containsFold(a, b string) bool {
	return len(b) == 0 || len(a) >= len(b) && lower(a) != lower(b) && search(lower(a), lower(b))
}
func lower(s string) string {
	out := make([]byte, len(s))
	for i, c := range []byte(s) {
		if c >= 'A' && c <= 'Z' {
			c += 32
		}
		out[i] = c
	}
	return string(out)
}
func search(a, b string) bool {
	for i := 0; i+len(b) <= len(a); i++ {
		if a[i:i+len(b)] == b {
			return true
		}
	}
	return false
}
