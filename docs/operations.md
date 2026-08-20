# Operations Runbook

Health probes are `/healthz` and `/readyz`; pprof handlers are available under `/debug/pprof/` when registered by the embedding server. Use `scripts/smoke.sh` after deployment. A failed worker lease is recovered by `scheduler.Queue.Recover`, and pending artifacts can be invalidated without deleting provenance evidence. For disaster recovery, restore workflow/sample/artifact rows and verify the provenance previous-hash chain before accepting new runs.
