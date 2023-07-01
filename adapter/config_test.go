package adapter

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestConfig(t *testing.T) {
	t.Run("Default", func(t *testing.T) {
		apiKey := "test-token"

		cfg := NewConfig(apiKey)

		assert.Same(t, http.DefaultClient, cfg.HTTPClient)
		assert.Equal(t, FormatJSON, cfg.Format)
		assert.Equal(t, apiKey, cfg.ApiKey)
		assert.False(t, cfg.Debug)
	})
	t.Run("Custom", func(t *testing.T) {
		apiKey := "test-token"
		httpClient := &http.Client{}
		format := FormatXML

		cfg := NewConfig(apiKey, WithHTTPClient(httpClient), WithFormat(format), WithDebug(true))

		assert.Same(t, httpClient, cfg.HTTPClient)
		assert.Equal(t, format, cfg.Format)
		assert.Equal(t, apiKey, cfg.ApiKey)
		assert.True(t, cfg.Debug)
	})
	t.Run("EmptyToken", func(t *testing.T) {
		require.Panics(t, func() {
			NewConfig("")
		})
	})
}
