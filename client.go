package go_nova_poshta

import (
	"github.com/platx/go-nova-poshta/adapter"
	"github.com/platx/go-nova-poshta/api/address"
	"github.com/platx/go-nova-poshta/api/common"
	"github.com/platx/go-nova-poshta/api/trackingdocument"
)

type Client interface {
	Address() address.Api
	Common() common.Api
	TrackingDocument() trackingdocument.Api
}

type RequestAdapter interface {
	Req(modelName string, calledMethod string, props any, res any) error
}

type client struct {
	addressApi          address.Api
	commonApi           common.Api
	trackingDocumentApi trackingdocument.Api
}

func NewClient(cfg adapter.Config) Client {
	adp := adapter.NewAdapter(cfg)

	return &client{
		addressApi:          address.NewApi(adp),
		commonApi:           common.NewApi(adp),
		trackingDocumentApi: trackingdocument.NewApi(adp),
	}
}

func NewClientJSON(httpClient adapter.HTTPClient, apiKey string) Client {
	return NewClient(adapter.CreateConfig(
		apiKey,
		adapter.WithHTTPClient(httpClient),
		adapter.WithFormat(adapter.FormatJSON),
	))
}

func NewDefaultClientJSON(apiKey string) Client {
	return NewClient(adapter.CreateConfig(
		apiKey,
		adapter.WithFormat(adapter.FormatJSON),
	))
}

func NewClientXML(httpClient adapter.HTTPClient, apiKey string) Client {
	return NewClient(adapter.CreateConfig(
		apiKey,
		adapter.WithHTTPClient(httpClient),
		adapter.WithFormat(adapter.FormatXML),
	))
}

func NewDefaultClientXML(apiKey string) Client {
	return NewClient(adapter.CreateConfig(
		apiKey,
		adapter.WithFormat(adapter.FormatXML),
	))
}

func (c *client) Address() address.Api {
	return c.addressApi
}

func (c *client) Common() common.Api {
	return c.commonApi
}

func (c *client) TrackingDocument() trackingdocument.Api {
	return c.trackingDocumentApi
}
