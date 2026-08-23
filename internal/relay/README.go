package observatory

// Package relay owns route workspace domain behavior, storage boundaries, and presentation queries.
// Its APIs deliberately return cloned mutable values so desktop tools can edit snapshots safely.
func PackageName() string { return "canvas-relay" }
func PackageFeatures() []string {
	return []string{"durable-workspaces", "geometry-import", "render-frames", "safe-snapshots", "event-log"}
}
