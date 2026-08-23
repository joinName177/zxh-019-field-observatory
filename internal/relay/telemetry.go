package observatory

import "time"

type Span struct {
	Name              string
	Started, Finished time.Time
	Attributes        map[string]string
}

func StartSpan(name string) Span {
	return Span{Name: name, Started: time.Now().UTC(), Attributes: map[string]string{}}
}
func (s *Span) End()            { s.Finished = time.Now().UTC() }
func (s *Span) Set(k, v string) { s.Attributes[k] = v }
func (s Span) Duration() time.Duration {
	if s.Finished.IsZero() {
		return time.Since(s.Started)
	}
	return s.Finished.Sub(s.Started)
}
func (s Span) Done() bool { return !s.Finished.IsZero() }
