package enum

type OrderType string

const (
	OrderTypeCargoReturn OrderType = "orderCargoReturn"
	OrderTypeRedirecting OrderType = "orderRedirecting"
	OrderTypeChangeEW    OrderType = "orderChangeEW"
)
