package observatory

import (
	"context"
	"sort"
)

type WorkspaceReport struct {
	WorkspaceID string
	Revision    int
	LayerCount  int
	TagCounts   map[string]int
	Latest      Layer
}

func BuildWorkspaceReport(ctx context.Context, repo Repository, workspaceID string) (WorkspaceReport, error) {
	workspace, err := repo.Load(ctx, workspaceID)
	if err != nil {
		return WorkspaceReport{}, err
	}
	report := WorkspaceReport{WorkspaceID: workspace.ID, Revision: workspace.Revision, LayerCount: len(workspace.Layers), TagCounts: map[string]int{}}
	layers := make([]Layer, 0, len(workspace.Layers))
	for _, layer := range workspace.Layers {
		layers = append(layers, CloneLayer(layer))
		for _, tag := range layer.Tags {
			report.TagCounts[tag]++
		}
	}
	sort.Slice(layers, func(i, j int) bool { return layers[i].ID < layers[j].ID })
	if latest, ok := LatestLayer(layers); ok {
		report.Latest = latest
	}
	return report, nil
}
