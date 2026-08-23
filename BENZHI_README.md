# Field Observatory

Field Observatory is a local-first desktop geospatial review workspace. It stores field tracks and observation layers in JSON, supports validated imports, tag and time-window queries, viewport rendering, review exports, audit events, and durable reopen flows.

Run locally with `go run ./cmd/canvasrelay ./workspace.json`. The Docker delivery script accepts a Buildx platform such as `linux/arm64` or `linux/amd64` and produces a self-contained CLI image.
