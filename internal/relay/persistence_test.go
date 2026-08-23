package observatory

import (
	"context"
	"path/filepath"
	"testing"
)

func TestWorkspacePersistsAcrossReopen(t *testing.T) {
	path := filepath.Join(t.TempDir(), "workspace.json")
	repo, err := OpenFileRepository(path)
	if err != nil {
		t.Fatal(err)
	}
	ctx := context.Background()
	w := Workspace{ID: "ridge", Name: "Ridge Survey", Layers: map[string]Layer{}}
	if err := repo.Save(ctx, w); err != nil {
		t.Fatal(err)
	}
	if err := PutLayer(ctx, repo, "ridge", NewLayer("route", "Ridge", []Point{{1, 2}, {3, 4}}, []string{"survey"})); err != nil {
		t.Fatal(err)
	}
	reopened, err := OpenFileRepository(path)
	if err != nil {
		t.Fatal(err)
	}
	got, err := reopened.Load(ctx, "ridge")
	if err != nil {
		t.Fatal(err)
	}
	if got.Layers["route"].Name != "Ridge" || len(got.Layers["route"].Points) != 2 {
		t.Fatalf("unexpected reopen: %#v", got)
	}
}
