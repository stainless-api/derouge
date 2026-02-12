# derouge

> **Experimental** — this project is under active development and not ready for production use.

JWE-based credential proxy. Decrypts encrypted tokens and injects API credentials at the last hop, keeping raw secrets out of sandboxed environments.

## Build

```
scripts/build
```

## Run

```
./derouge -config config.json
```

See `config.example.json` for configuration options.

## Test

```
go test ./...
```

## Lint

```
scripts/lint
```

## Production Hardening TODO

- [ ] **Persist deny list** — currently in-memory, lost on restart, no cross-replica sync
- [ ] **Auth on `/revoke`** — endpoint is unauthenticated (unlike `/mint` which has `MintAuthMiddleware`)
- [ ] **Rate limiting and request body size limits** — unbounded JSON bodies on revoke/mint can OOM the process
- [ ] **TLS** — server listens plain HTTP; fine behind a terminating LB, not otherwise
- [ ] **Metrics / observability** — `RequestTrace` is logged but not exported to Prometheus or similar
- [ ] **Circuit breaker / retry budget for upstream** — slow/down upstream blocks goroutines up to `ResponseHeaderTimeout`
- [ ] **Meaningful health check** — `/health` always returns OK regardless of keystore or deny list state
- [ ] **Test coverage gaps** — keystore, httputil, and secretutil have no tests
- [ ] **Key rotation lifecycle** — keys are loaded from disk on startup with no rotation mechanism
- [ ] **Forwarded headers** — proxy doesn't set `X-Forwarded-For` / `X-Forwarded-Proto`

## License

[Elastic License 2.0 (ELv2)](LICENSE)
