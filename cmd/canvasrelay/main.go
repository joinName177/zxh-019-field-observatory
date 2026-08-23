package main

import (
	"context"
	"fmt"
	"os"

	relay "github.com/joinName177/zxh-019-field-observatory/internal/relay"
)

func main() {
	path := "workspace.json"
	if len(os.Args) > 1 {
		path = os.Args[1]
	}
	repo, err := relay.OpenFileRepository(path)
	if err != nil {
		panic(err)
	}
	ctx := context.Background()
	cat := relay.NewCatalog(repo)
	if _, err = cat.Create(ctx, "demo", "Field Observatory"); err != nil && err != relay.ErrWorkspaceNotFound {
		fmt.Println(err)
	}
	audit := relay.NewAuditLog()
	audit.Record(relay.AuditEntry{Actor: "desktop-cli", Action: "workspace-opened", Resource: "demo"})
	fmt.Println("field observatory ready", path, "events", audit.Len())
}
