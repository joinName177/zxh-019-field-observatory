package observatory

import "context"

type Repository interface {
	Load(context.Context, string) (Workspace, error)
	Save(context.Context, Workspace) error
	Delete(context.Context, string) error
}
type LayerRepository interface {
	FindLayer(context.Context, string, string) (Layer, error)
	PutLayer(context.Context, string, Layer) error
}
