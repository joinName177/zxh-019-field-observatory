# Field Observatory

Field Observatory is a local desktop tool for curators who review field survey tracks and observation layers before publishing a daily map packet. A workspace stores named tracks, coordinate geometry, tags, and revision history in a durable JSON file. The product flow is real and local-first: create or reopen a workspace, validate and enqueue an import, process the import, query isolated snapshots, render a viewport, filter layers by tags or time window, export a review packet, and record an audit event.

The application is split into domain models and validation, repository persistence, import queue and transaction services, query and rendering services, and operational reporting. Cancellation is part of the import contract: a field operator can abandon a queued import while another desktop action holds the file lock, and an abandoned import must not appear after reopening.
