// Package secretutil wraps runtime/secret.Do (Go 1.26+) to zero sensitive
// memory after use, falling back to a plain function call without the experiment.
package secretutil
