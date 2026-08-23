package observatory

type Selection struct {
	LayerIDs map[string]bool
	Active   string
}

func NewSelection() *Selection { return &Selection{LayerIDs: map[string]bool{}} }
func (s *Selection) Toggle(id string) {
	s.LayerIDs[id] = !s.LayerIDs[id]
	if s.LayerIDs[id] {
		s.Active = id
	} else if s.Active == id {
		s.Active = ""
	}
}
func (s *Selection) Set(id string)           { s.LayerIDs = map[string]bool{id: true}; s.Active = id }
func (s *Selection) Clear()                  { s.LayerIDs = map[string]bool{}; s.Active = "" }
func (s *Selection) Contains(id string) bool { return s.LayerIDs[id] }
func (s *Selection) IDs() []string {
	out := []string{}
	for id, ok := range s.LayerIDs {
		if ok {
			out = append(out, id)
		}
	}
	return out
}
