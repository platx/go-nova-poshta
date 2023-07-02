package go_nova_poshta

import (
	"testing"

	"github.com/stretchr/testify/assert"

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

func TestClient(t *testing.T) {
	cfg := adapter.NewConfig("fake-token")

	c := NewClient(cfg)

	assert.IsType(t, &client{}, c)

	adp := adapter.NewAdapter(cfg)

	assert.IsType(t, address.NewModel(adp), c.Address())
	assert.IsType(t, counterparty.NewModel(adp), c.Counterparty())
	assert.IsType(t, contactperson.NewModel(adp), c.ContactPerson())
	assert.IsType(t, scansheet.NewModel(adp), c.ScanSheet())
	assert.IsType(t, common.NewModel(adp), c.Common())
	assert.IsType(t, additionalservice.NewModel(adp), c.AdditionalService())
	assert.IsType(t, internetdocument.NewModel(adp), c.InternetDocument())
	assert.IsType(t, trackingdocument.NewModel(adp), c.TrackingDocument())
	assert.IsType(t, adp, c.Custom())
}
