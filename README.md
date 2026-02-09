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

## License

[Elastic License 2.0 (ELv2)](LICENSE)
