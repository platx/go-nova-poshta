package enum

type PayerType string

const (
	PayerTypeRecipient   PayerType = "Recipient"
	PayerTypeSender      PayerType = "Sender"
	PayerTypeThirdPerson PayerType = "ThirdPerson"
)
