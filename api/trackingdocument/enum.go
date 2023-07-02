package trackingdocument

import (
	"encoding/xml"

	"github.com/platx/go-nova-poshta/custom/types"
)

type TrackingDocumentStatus types.IntString

func (v *TrackingDocumentStatus) UnmarshalJSON(data []byte) error {
	var bv types.IntString

	if err := bv.UnmarshalJSON(data); err != nil {
		return err
	}

	*v = TrackingDocumentStatus(bv)

	return nil
}

func (v *TrackingDocumentStatus) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var bv types.IntString

	if err := bv.UnmarshalXML(d, start); err != nil {
		return err
	}

	*v = TrackingDocumentStatus(bv)

	return nil
}

func (v TrackingDocumentStatus) MarshalJSON() ([]byte, error) {
	bv := types.IntString(v)

	return bv.MarshalJSON()
}

func (v TrackingDocumentStatus) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	bv := types.IntString(v)

	return bv.MarshalXML(e, start)
}

const (
	TrackingDocumentStatusCreatedBySender                  TrackingDocumentStatus = 1
	TrackingDocumentStatusRemoved                          TrackingDocumentStatus = 2
	TrackingDocumentStatusNotFound                         TrackingDocumentStatus = 3
	TrackingDocumentStatusCrossAreaInSourceCity            TrackingDocumentStatus = 4
	TrackingDocumentStatusGoingToTargetCity                TrackingDocumentStatus = 5
	TrackingDocumentStatusInTargetCity                     TrackingDocumentStatus = 6
	TrackingDocumentStatusInTargetWarehouse                TrackingDocumentStatus = 7
	TrackingDocumentStatusInTargetPostomat                 TrackingDocumentStatus = 8
	TrackingDocumentStatusReceived                         TrackingDocumentStatus = 9
	TrackingDocumentStatusReceivedMoneySending             TrackingDocumentStatus = 10
	TrackingDocumentStatusReceivedMoneyReceived            TrackingDocumentStatus = 11
	TrackingDocumentStatusShipmentCompleting               TrackingDocumentStatus = 12
	TrackingDocumentStatusInnerCityInCity                  TrackingDocumentStatus = 41
	TrackingDocumentStatusGoingToReceiver                  TrackingDocumentStatus = 101
	TrackingDocumentStatusCancelledBySender                TrackingDocumentStatus = 102
	TrackingDocumentStatusCancelledByReceiver              TrackingDocumentStatus = 103
	TrackingDocumentStatusAddressChanged                   TrackingDocumentStatus = 104
	TrackingDocumentStatusStorageCancelled                 TrackingDocumentStatus = 105
	TrackingDocumentStatusReceivedAndCreatedreturnShipping TrackingDocumentStatus = 106
	TrackingDocumentStatusDeliveryFailedReceiverNotFound   TrackingDocumentStatus = 111
	TrackingDocumentStatusDateChangedByReceiver            TrackingDocumentStatus = 112
)

func (s TrackingDocumentStatus) String() string {
	switch s {
	case TrackingDocumentStatusCreatedBySender:
		return "Відправник самостійно створив цю накладну, але ще не надав до відправки"
	case TrackingDocumentStatusRemoved:
		return "Видалено"
	case TrackingDocumentStatusNotFound:
		return "Номер не знайдено"
	case TrackingDocumentStatusCrossAreaInSourceCity:
		return "Відправлення у місті ХХXХ. (Статус для межобластных отправлений)"
	case TrackingDocumentStatusGoingToTargetCity:
		return "Відправлення прямує до міста YYYY"
	case TrackingDocumentStatusInTargetCity:
		return "Відправлення у місті YYYY, орієнтовна доставка до ВІДДІЛЕННЯ-XXX dd-mm. Очікуйте додаткове повідомлення про прибуття"
	case TrackingDocumentStatusInTargetWarehouse:
		return "Прибув на відділення"
	case TrackingDocumentStatusInTargetPostomat:
		return "Прибув на відділення (завантажено в Поштомат)"
	case TrackingDocumentStatusReceived:
		return "Відправлення отримано"
	case TrackingDocumentStatusReceivedMoneySending:
		return "Відправлення отримано %DateReceived%. Протягом доби ви одержите SMS-повідомлення про надходження грошового переказу та зможете отримати його в касі відділення «Нова пошта»"
	case TrackingDocumentStatusReceivedMoneyReceived:
		return "Відправлення отримано %DateReceived%. Грошовий переказ видано одержувачу."
	case TrackingDocumentStatusShipmentCompleting:
		return "Нова Пошта комплектує ваше відправлення"
	case TrackingDocumentStatusInnerCityInCity:
		return "Відправлення у місті ХХXХ. (Статус для услуг локал стандарт и локал экспресс - доставка в пределах города)"
	case TrackingDocumentStatusGoingToReceiver:
		return "На шляху до одержувача"
	case TrackingDocumentStatusCancelledBySender:
		return "Відмова від отримання (Відправником створено замовлення на повернення)"
	case TrackingDocumentStatusCancelledByReceiver:
		return "Відмова одержувача (отримувач відмовився від відправлення)"
	case TrackingDocumentStatusAddressChanged:
		return "Змінено адресу"
	case TrackingDocumentStatusStorageCancelled:
		return "Припинено зберігання"
	case TrackingDocumentStatusReceivedAndCreatedreturnShipping:
		return "Одержано і створено ЄН зворотньої доставки"
	case TrackingDocumentStatusDeliveryFailedReceiverNotFound:
		return "Невдала спроба доставки через відсутність Одержувача на адресі або зв'язку з ним"
	case TrackingDocumentStatusDateChangedByReceiver:
		return "Дата доставки перенесена Одержувачем\n\n"
	}

	return "Unknown"
}
