package adapter

import (
	"net/http"
)

type Config struct {
	HTTPClient httpClient
	Format     Format
	ApiKey     string
	Debug      bool
	Logger     logger
}

func NewConfig(apiKey string, opts ...Option) Config {
	if apiKey == "" {
		panic("api key is required")
	}

	c := Config{
		HTTPClient: http.DefaultClient,
		Format:     FormatJSON,
		ApiKey:     apiKey,
		Logger:     &noopLogger{},
	}

	for _, opt := range opts {
		opt(&c)
	}

	return c
}

type Option func(*Config)

func WithHTTPClient(v httpClient) Option {
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

		switch c.Logger.(type) {
		case *noopLogger:
			c.Logger = &defaultLogger{}
		}
	}
}

func WithLogger(v logger) Option {
	return func(c *Config) {
		c.Logger = v
	}
}
