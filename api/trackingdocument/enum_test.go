package trackingdocument

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestTrackingDocumentStatus(t *testing.T) {
	type testStruct struct {
		Status TrackingDocumentStatus `json:"Status" xml:"Status"`
	}

	for _, tc := range []struct {
		encoded  string
		expected TrackingDocumentStatus
	}{
		{"1", TrackingDocumentStatusCreatedBySender},
		{"2", TrackingDocumentStatusRemoved},
		{"3", TrackingDocumentStatusNotFound},
		{"4", TrackingDocumentStatusCrossAreaInSourceCity},
		{"5", TrackingDocumentStatusGoingToTargetCity},
		{"6", TrackingDocumentStatusInTargetCity},
		{"7", TrackingDocumentStatusInTargetWarehouse},
		{"8", TrackingDocumentStatusInTargetPostomat},
		{"9", TrackingDocumentStatusReceived},
		{"10", TrackingDocumentStatusReceivedMoneySending},
		{"11", TrackingDocumentStatusReceivedMoneyReceived},
		{"12", TrackingDocumentStatusShipmentCompleting},
		{"41", TrackingDocumentStatusInnerCityInCity},
		{"101", TrackingDocumentStatusGoingToReceiver},
		{"102", TrackingDocumentStatusCancelledBySender},
		{"103", TrackingDocumentStatusCancelledByReceiver},
		{"104", TrackingDocumentStatusAddressChanged},
		{"105", TrackingDocumentStatusStorageCancelled},
		{"106", TrackingDocumentStatusReceivedAndCreatedreturnShipping},
		{"111", TrackingDocumentStatusDeliveryFailedReceiverNotFound},
		{"112", TrackingDocumentStatusDateChangedByReceiver},
	} {
		t.Run(fmt.Sprintf("JSON/%s", tc.encoded), func(t *testing.T) {
			var decoded testStruct

			require.NoError(t, json.Unmarshal([]byte(fmt.Sprintf(`{"Status":"%s"}`, tc.encoded)), &decoded))
			assert.Equal(t, tc.expected, decoded.Status)
			assert.NotEqual(t, "unknown", decoded.Status.String())

			encoded, err := json.Marshal(decoded)

			require.NoError(t, err)

			assert.Equal(t, fmt.Sprintf(`{"Status":"%s"}`, tc.encoded), string(encoded))
		})
		t.Run(fmt.Sprintf("XML/%s", tc.encoded), func(t *testing.T) {
			var decoded testStruct

			require.NoError(t, xml.Unmarshal([]byte(
				fmt.Sprintf("<testStruct><Status>%s</Status></testStruct>", tc.encoded),
			), &decoded))
			assert.Equal(t, tc.expected, decoded.Status)
			assert.NotEqual(t, "unknown", decoded.Status.String())

			encoded, err := xml.Marshal(decoded)

			require.NoError(t, err)

			assert.Equal(
				t,
				fmt.Sprintf("<testStruct><Status>%s</Status></testStruct>", tc.encoded),
				string(encoded),
			)
		})
	}

	t.Run("Unknown", func(t *testing.T) {
		var decoded TrackingDocumentStatus

		require.NoError(t, json.Unmarshal([]byte(`"9999"`), &decoded))
		assert.Equal(t, "Unknown", decoded.String())
	})
}
