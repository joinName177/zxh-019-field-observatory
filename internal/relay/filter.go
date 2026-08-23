package observatory

import "strings"

func ByTag(layers []Layer, tag string) []Layer {
	out := []Layer{}
	for _, l := range layers {
		for _, t := range l.Tags {
			if strings.EqualFold(t, tag) {
				out = append(out, CloneLayer(l))
				break
			}
		}
	}
	return out
}
func ByName(layers []Layer, needle string) []Layer {
	needle = strings.ToLower(needle)
	out := []Layer{}
	for _, l := range layers {
		if strings.Contains(strings.ToLower(l.Name), needle) {
			out = append(out, CloneLayer(l))
		}
	}
	return out
}
func SortByName(layers []Layer) []Layer {
	out := append([]Layer(nil), layers...)
	for i := 0; i < len(out); i++ {
		for j := i + 1; j < len(out); j++ {
			if out[j].Name < out[i].Name {
				out[i], out[j] = out[j], out[i]
			}
		}
	}
	return out
}
