package observatory

type Annotation struct {
	ID, LayerID, Text, Author string
	Position                  Point
}
type AnnotationBook struct{ items map[string]Annotation }

func NewAnnotationBook() *AnnotationBook { return &AnnotationBook{items: map[string]Annotation{}} }
func (b *AnnotationBook) Add(a Annotation) error {
	if a.ID == "" || a.LayerID == "" {
		return ErrInvalidLayer
	}
	b.items[a.ID] = a
	return nil
}
func (b *AnnotationBook) Get(id string) (Annotation, bool) { a, ok := b.items[id]; return a, ok }
func (b *AnnotationBook) Remove(id string)                 { delete(b.items, id) }
func (b *AnnotationBook) ForLayer(id string) []Annotation {
	out := []Annotation{}
	for _, a := range b.items {
		if a.LayerID == id {
			out = append(out, a)
		}
	}
	return out
}
func (b *AnnotationBook) Count() int { return len(b.items) }
