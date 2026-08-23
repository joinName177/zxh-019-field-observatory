package observatory

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"time"
)

type ReviewPacket struct {
	WorkspaceID string       `json:"workspace_id"`
	Revision    int          `json:"revision"`
	GeneratedAt time.Time    `json:"generated_at"`
	Layers      []Layer      `json:"layers"`
	Events      []AuditEntry `json:"events,omitempty"`
}

func BuildReviewPacket(ctx context.Context, repo Repository, audit *AuditLog, workspaceID string) (ReviewPacket, error) {
	if err := ctx.Err(); err != nil {
		return ReviewPacket{}, err
	}
	workspace, err := repo.Load(ctx, workspaceID)
	if err != nil {
		return ReviewPacket{}, err
	}
	layers := make([]Layer, 0, len(workspace.Layers))
	for _, layer := range workspace.Layers {
		layers = append(layers, CloneLayer(layer))
	}
	packet := ReviewPacket{WorkspaceID: workspace.ID, Revision: workspace.Revision, GeneratedAt: time.Now().UTC(), Layers: layers}
	if audit != nil {
		packet.Events = audit.Snapshot()
	}
	return packet, nil
}

func EncodeReviewPacket(ctx context.Context, packet ReviewPacket) ([]byte, error) {
	if err := ctx.Err(); err != nil {
		return nil, err
	}
	data, err := json.MarshalIndent(packet, "", "  ")
	if err != nil {
		return nil, fmt.Errorf("encode review packet: %w", err)
	}
	return bytes.TrimSpace(data), nil
}
