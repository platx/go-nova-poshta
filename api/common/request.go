package common

import "github.com/google/uuid"

type GetTimeIntervalsFilter struct {
	RecipientCityRef uuid.UUID `json:"RecipientCityRef" xml:"RecipientCityRef"`
	DateTime         string    `json:"DateTime" xml:"DateTime"` // format: dd.mm.yy TODO: create custom type
}
