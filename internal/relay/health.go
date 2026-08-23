package observatory

import "context"

type Health struct {
	Repository bool
	Workspace  string
	Revision   int
}

func CheckHealth(ctx context.Context, r Repository, id string) Health {
	w, err := r.Load(ctx, id)
	if err != nil {
		return Health{}
	}
	return Health{Repository: true, Workspace: w.ID, Revision: w.Revision}
}
func (h Health) Ready() bool { return h.Repository && h.Workspace != "" }
