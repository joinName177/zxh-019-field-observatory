package observatory

import "context"

type Session struct {
	catalog  *Catalog
	importer *Importer
	events   *EventLog
}

func NewSession(c *Catalog, i *Importer, e *EventLog) *Session {
	return &Session{catalog: c, importer: i, events: e}
}
func (s *Session) Open(ctx context.Context, id, name string) (Workspace, error) {
	w, err := s.catalog.Create(ctx, id, name)
	if err == nil {
		s.events.Append(Change{WorkspaceID: id, Kind: "open"})
	}
	return w, err
}
func (s *Session) Add(ctx context.Context, wid string, l Layer) (Layer, error) {
	out, err := s.importer.Import(ctx, wid, l)
	if err == nil {
		s.events.Append(Change{WorkspaceID: wid, LayerID: l.ID, Kind: "import"})
	}
	return out, err
}
func (s *Session) Rename(ctx context.Context, id, name string) (Workspace, error) {
	w, err := s.catalog.Rename(ctx, id, name)
	if err == nil {
		s.events.Append(Change{WorkspaceID: id, Kind: "rename"})
	}
	return w, err
}
