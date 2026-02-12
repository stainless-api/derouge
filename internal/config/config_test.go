package config

import (
	"encoding/json"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestParseConfigValueEnvEmpty(t *testing.T) {
	const envVar = "DEROUGE_TEST_EMPTY_ENV"
	t.Setenv(envVar, "")

	raw := json.RawMessage(`{"$env":"` + envVar + `"}`)
	val, err := parseConfigValue(raw)
	require.NoError(t, err, "env var set to empty string should succeed")
	assert.Equal(t, "", val)
}

func TestParseConfigValueEnvUnset(t *testing.T) {
	const envVar = "DEROUGE_TEST_UNSET_ENV"
	os.Unsetenv(envVar)

	raw := json.RawMessage(`{"$env":"` + envVar + `"}`)
	_, err := parseConfigValue(raw)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "not set")
}

func TestParseConfigValueLiteralString(t *testing.T) {
	raw := json.RawMessage(`"hello"`)
	val, err := parseConfigValue(raw)
	require.NoError(t, err)
	assert.Equal(t, "hello", val)
}

func TestParseConfigValueEnvNonEmpty(t *testing.T) {
	const envVar = "DEROUGE_TEST_SET_ENV"
	t.Setenv(envVar, "my-secret")

	raw := json.RawMessage(`{"$env":"` + envVar + `"}`)
	val, err := parseConfigValue(raw)
	require.NoError(t, err)
	assert.Equal(t, "my-secret", val)
}

func TestLoadConfig(t *testing.T) {
	dir := t.TempDir()
	cfgPath := filepath.Join(dir, "config.json")
	err := os.WriteFile(cfgPath, []byte(`{
		"addr": ":9090",
		"keyDir": "/tmp/keys",
		"mintSecret": "literal-secret"
	}`), 0644)
	require.NoError(t, err)

	cfg, err := Load(cfgPath)
	require.NoError(t, err)
	assert.Equal(t, ":9090", cfg.Addr)
	assert.Equal(t, Secret("literal-secret"), cfg.MintSecret)
}
