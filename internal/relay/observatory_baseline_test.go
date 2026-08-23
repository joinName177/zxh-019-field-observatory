package observatory

import (
	"context"
	"testing"
	"time"
)

func TestFieldObservatoryReviewFlow(t *testing.T) {
	ctx := context.Background()
	repo := NewMemoryRepository()
	if _, err := NewCatalog(repo).Create(ctx, "field-day", "Field Day"); err != nil {
		t.Fatal(err)
	}
	layer := NewLayer("ridge", "Ridge survey", []Point{{1, 2}, {3, 4}}, []string{"north", "review"})
	if _, err := NewImporter(repo).Import(ctx, "field-day", layer); err != nil {
		t.Fatal(err)
	}
	audit := NewAuditLog()
	audit.Record(AuditEntry{Actor: "operator", Action: "import", Resource: layer.ID})
	packet, err := BuildReviewPacket(ctx, repo, audit, "field-day")
	if err != nil || len(packet.Layers) != 1 || packet.Events[0].Action != "import" {
		t.Fatalf("unexpected packet: %#v, %v", packet, err)
	}
	report, err := BuildWorkspaceReport(ctx, repo, "field-day")
	if err != nil || report.TagCounts["north"] != 1 {
		t.Fatalf("unexpected report: %#v, %v", report, err)
	}
	window := NewTimeWindow(time.Now().Add(-time.Minute), time.Now().Add(time.Minute))
	if len(FilterLayersByUpdated(packet.Layers, window)) != 1 {
		t.Fatal("recent layer was excluded from time window")
	}
}
