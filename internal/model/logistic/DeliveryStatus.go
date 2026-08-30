package logistic

type DeliveryStatus int

const (
	DeliveryStatusUnknown DeliveryStatus = iota
	DeliveryStatusCreated
	DeliveryStatusAccepted
	DeliveryStatusInTransit
	DeliveryStatusDelivered
	DeliveryStatusCanceled
)

func (d DeliveryStatus) String() string {
	switch d {
	case DeliveryStatusUnknown:
		return "unknown"
	case DeliveryStatusCreated:
		return "created"
	case DeliveryStatusAccepted:

		return "accepted"
	case DeliveryStatusInTransit:
		return "in transit"
	case DeliveryStatusDelivered:
		return "delivered"
	case DeliveryStatusCanceled:
		return "canceled"
	default:
		return "invalid"

	}
}
