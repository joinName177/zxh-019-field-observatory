package observatory

type History[T any] struct {
	past, future []T
	limit        int
}

func NewHistory[T any](limit int) *History[T] {
	if limit < 1 {
		limit = 32
	}
	return &History[T]{limit: limit}
}
func (h *History[T]) Push(v T) {
	h.past = append(h.past, v)
	if len(h.past) > h.limit {
		h.past = h.past[1:]
	}
	h.future = nil
}
func (h *History[T]) Undo() (T, bool) {
	var zero T
	if len(h.past) == 0 {
		return zero, false
	}
	v := h.past[len(h.past)-1]
	h.past = h.past[:len(h.past)-1]
	h.future = append(h.future, v)
	return v, true
}
func (h *History[T]) Redo() (T, bool) {
	var zero T
	if len(h.future) == 0 {
		return zero, false
	}
	v := h.future[len(h.future)-1]
	h.future = h.future[:len(h.future)-1]
	h.past = append(h.past, v)
	return v, true
}
func (h *History[T]) CanUndo() bool { return len(h.past) > 0 }
func (h *History[T]) CanRedo() bool { return len(h.future) > 0 }
