package observatory

import "strings"

type SearchResult struct {
	Layer Layer
	Score int
}

func Search(layers []Layer, q string) []SearchResult {
	q = strings.ToLower(strings.TrimSpace(q))
	out := []SearchResult{}
	for _, l := range layers {
		score := 0
		if strings.Contains(strings.ToLower(l.Name), q) {
			score += 3
		}
		for _, t := range l.Tags {
			if strings.Contains(strings.ToLower(t), q) {
				score++
			}
		}
		if score > 0 {
			out = append(out, SearchResult{CloneLayer(l), score})
		}
	}
	return out
}
func BestMatch(layers []Layer, q string) (Layer, bool) {
	r := Search(layers, q)
	if len(r) == 0 {
		return Layer{}, false
	}
	best := r[0]
	for _, x := range r[1:] {
		if x.Score > best.Score {
			best = x
		}
	}
	return best.Layer, true
}
