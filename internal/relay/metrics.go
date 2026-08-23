package observatory

import "sync/atomic"

type Metrics struct {
	imports  atomic.Int64
	renders  atomic.Int64
	failures atomic.Int64
}

func (m *Metrics) Imported() { m.imports.Add(1) }
func (m *Metrics) Rendered() { m.renders.Add(1) }
func (m *Metrics) Failed()   { m.failures.Add(1) }
func (m *Metrics) Snapshot() map[string]int64 {
	return map[string]int64{"imports": m.imports.Load(), "renders": m.renders.Load(), "failures": m.failures.Load()}
}
