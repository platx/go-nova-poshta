package adapter

import (
	"net/http"
)

type Config struct {
	HTTPClient HTTPClient
	Format     Format
	ApiKey     string
	Debug      bool
}

func NewConfig(apiKey string, opts ...Option) Config {
	if apiKey == "" {
		panic("api key is required")
	}

	c := Config{
		HTTPClient: http.DefaultClient,
		Format:     FormatJSON,
		ApiKey:     apiKey,
	}

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

func WithDebug(v bool) Option {
	return func(c *Config) {
		c.Debug = v
	}
}
