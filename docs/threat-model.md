# Threat Model

Threats include malicious workflow graphs, oversized artifact metadata, forged provenance records, replayed run submissions and worker output floods. Controls include cycle validation, bounded request bodies, idempotency keys, HMAC/nonce helpers, digest validation, immutable provenance hashes, output limits, tenant quotas and explicit worker adapters that never execute arbitrary scripts.
