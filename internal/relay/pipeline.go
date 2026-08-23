package observatory

import "context"

type Pipeline struct {
	session  *Session
	renderer *Renderer
	exporter *Exporter
	metrics  *Metrics
}

func NewPipeline(s *Session, r *Renderer, e *Exporter, m *Metrics) *Pipeline {
	return &Pipeline{session: s, renderer: r, exporter: e, metrics: m}
}
func (p *Pipeline) ImportAndFrame(ctx context.Context, wid string, l Layer) (RenderFrame, error) {
	if _, err := p.session.Add(ctx, wid, l); err != nil {
		p.metrics.Failed()
		return RenderFrame{}, err
	}
	p.metrics.Imported()
	f, err := p.renderer.Frame(ctx, wid)
	if err != nil {
		p.metrics.Failed()
		return RenderFrame{}, err
	}
	p.metrics.Rendered()
	return f, nil
}
func (p *Pipeline) Export(ctx context.Context, wid string) ([]byte, error) {
	return p.exporter.JSON(ctx, wid)
}
