package go_nova_poshta

import (
	"github.com/platx/go-nova-poshta/adapter"
	"github.com/platx/go-nova-poshta/api/additionalservice"
	"github.com/platx/go-nova-poshta/api/address"
	"github.com/platx/go-nova-poshta/api/common"
	"github.com/platx/go-nova-poshta/api/contactperson"
	"github.com/platx/go-nova-poshta/api/counterparty"
	"github.com/platx/go-nova-poshta/api/internetdocument"
	"github.com/platx/go-nova-poshta/api/scansheet"
	"github.com/platx/go-nova-poshta/api/trackingdocument"
)

type Client interface {
	Address() address.Model
	Counterparty() counterparty.Model
	ContactPerson() contactperson.Model
	ScanSheet() scansheet.Model
	Common() common.Model
	AdditionalService() additionalservice.Model
	InternetDocument() internetdocument.Model
	TrackingDocument() trackingdocument.Model
}

type client struct {
	address           address.Model
	counterparty      counterparty.Model
	contactPerson     contactperson.Model
	scanSheet         scansheet.Model
	common            common.Model
	additionalService additionalservice.Model
	trackingDocument  trackingdocument.Model
	internetDocument  internetdocument.Model
}

func NewClient(cfg adapter.Config) Client {
	adp := adapter.NewAdapter(cfg)

	return &client{
		address:           address.NewModel(adp),
		counterparty:      counterparty.NewModel(adp),
		contactPerson:     contactperson.NewModel(adp),
		scanSheet:         scansheet.NewModel(adp),
		common:            common.NewModel(adp),
		additionalService: additionalservice.NewModel(adp),
		internetDocument:  internetdocument.NewModel(adp),
		trackingDocument:  trackingdocument.NewModel(adp),
	}
}

func (c *client) Address() address.Model {
	return c.address
}

func (c *client) Counterparty() counterparty.Model {
	return c.counterparty
}

func (c *client) ContactPerson() contactperson.Model {
	return c.contactPerson
}

func (c *client) ScanSheet() scansheet.Model {
	return c.scanSheet
}

func (c *client) Common() common.Model {
	return c.common
}

func (c *client) AdditionalService() additionalservice.Model {
	return c.additionalService
}

func (c *client) InternetDocument() internetdocument.Model {
	return c.internetDocument
}

func (c *client) TrackingDocument() trackingdocument.Model {
	return c.trackingDocument
}
