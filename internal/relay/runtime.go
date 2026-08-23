package observatory

import "context"

type Runtime struct {
	Repo     Repository
	Catalog  *Catalog
	Query    *QueryService
	Importer *Importer
	Renderer *Renderer
	Exporter *Exporter
	Events   *EventLog
	Metrics  *Metrics
}

func NewRuntime(repo Repository) *Runtime {
	q := NewQueryService(repo)
	e := NewEventLog()
	m := &Metrics{}
	return &Runtime{Repo: repo, Catalog: NewCatalog(repo), Query: q, Importer: NewImporter(repo), Renderer: NewRenderer(q), Exporter: NewExporter(q), Events: e, Metrics: m}
}
func (r *Runtime) Open(ctx context.Context, id, name string) (Workspace, error) {
	return r.Catalog.Create(ctx, id, name)
}
func (r *Runtime) Close(ctx context.Context) error { return nil }
func (r *Runtime) Ready(ctx context.Context, id string) bool {
	return CheckHealth(ctx, r.Repo, id).Ready()
}
