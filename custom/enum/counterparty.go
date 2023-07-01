package enum

type CounterpartyProperty string

const (
	CounterpartyPropertySender      CounterpartyProperty = "Sender"
	CounterpartyPropertyRecipient   CounterpartyProperty = "Recipient"
	CounterpartyPropertyThirdPerson CounterpartyProperty = "ThirdPerson"
)

type CounterpartyType string

const (
	CounterpartyTypePrivatePerson CounterpartyType = "PrivatePerson"
	CounterpartyTypeOrganization  CounterpartyType = "Organization"
)
