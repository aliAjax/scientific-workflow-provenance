#!/bin/sh
set -eu
base="${BASE_URL:-http://127.0.0.1:8088}"
curl -fsS "$base/healthz" >/dev/null
curl -fsS -X POST "$base/api/v1/workflows" -H 'Content-Type: application/json' -d '{"id":"smoke","name":"smoke","version":1,"nodes":[{"id":"source","name":"source","kind":"task","outputs":["raw"]},{"id":"analysis","name":"analysis","kind":"task","inputs":["raw"],"outputs":["result"]}],"edges":[{"from":"source","to":"analysis"}]}' >/dev/null
curl -fsS "$base/api/v1/workflows/smoke/validate" >/dev/null
curl -fsS -X POST "$base/api/v1/samples" -H 'Content-Type: application/json' -d '{"id":"sample-1","name":"sample","source":"lab","batch":"b1"}' >/dev/null
curl -fsS "$base/api/v1/samples/sample-1/lineage" >/dev/null
echo smoke-ok
