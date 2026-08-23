package observatory

import "sync"

type Registry struct {
	mu     sync.RWMutex
	values map[string]string
}

func NewRegistry() *Registry        { return &Registry{values: map[string]string{}} }
func (r *Registry) Set(k, v string) { r.mu.Lock(); defer r.mu.Unlock(); r.values[k] = v }
func (r *Registry) Get(k string) (string, bool) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	v, ok := r.values[k]
	return v, ok
}
func (r *Registry) Delete(k string) { r.mu.Lock(); defer r.mu.Unlock(); delete(r.values, k) }
func (r *Registry) Snapshot() map[string]string {
	r.mu.RLock()
	defer r.mu.RUnlock()
	out := map[string]string{}
	for k, v := range r.values {
		out[k] = v
	}
	return out
}
