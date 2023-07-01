package common

import (
	"github.com/google/uuid"

	"github.com/platx/go-nova-poshta/custom/types"
)

type GetTimeIntervalsReq struct {
	RecipientCityRef uuid.UUID        `json:"RecipientCityRef" xml:"RecipientCityRef"`
	DateTime         types.CustomDate `json:"DateTime" xml:"DateTime"`
}
