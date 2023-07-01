package adapter

import (
	"net/http"
)

type Config struct {
	HTTPClient HTTPClient
	Format     Format
	ApiKey     string
}

func CreateConfig(apiKey string, opts ...Option) Config {
	if apiKey == "" {
		panic("api key is required")
	}

	c := Config{
		HTTPClient: http.DefaultClient,
		Format:     FormatJSON,
	}

	WithApiKey(apiKey)(&c)

	for _, opt := range opts {
		opt(&c)
	}

	return c
}

type Option func(*Config)

func WithHTTPClient(v HTTPClient) Option {
	return func(c *Config) {
		c.HTTPClient = v
	}
}

func WithFormat(v Format) Option {
	return func(c *Config) {
		c.Format = v
	}
}

func WithApiKey(v string) Option {
	return func(c *Config) {
		c.ApiKey = v
	}
}
