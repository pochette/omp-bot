package logistic

type DeliveryStatus string

const (
	DeliveryStatusUnknown   DeliveryStatus = "Unknown"
	DeliveryStatusCreated   DeliveryStatus = "Created"
	DeliveryStatusAccepted  DeliveryStatus = "Accepted"
	DeliveryStatusInTransit DeliveryStatus = "In transit"
	DeliveryStatusDelivered DeliveryStatus = "Delivered"
	DeliveryStatusCanceled  DeliveryStatus = "Canceled"
	DeliveryStatusUpdated   DeliveryStatus = "Updated"
)

//type DeliveryStatus int
//
//const (
//	DeliveryStatusUnknown DeliveryStatus = iota
//	DeliveryStatusCreated
//	DeliveryStatusAccepted
//	DeliveryStatusInTransit
//	DeliveryStatusDelivered
//	DeliveryStatusCanceled
//	DeliveryStatusUpdated
//)
//
//func (d DeliveryStatus) String() string {
//	switch d {
//	case DeliveryStatusUnknown:
//		return "unknown"
//	case DeliveryStatusCreated:
//		return "created"
//	case DeliveryStatusAccepted:
//
//		return "accepted"
//	case DeliveryStatusInTransit:
//		return "in transit"
//	case DeliveryStatusDelivered:
//		return "delivered"
//	case DeliveryStatusCanceled:
//		return "canceled"
//	case DeliveryStatusUpdated:
//		return "updated"
//
//	default:
//		return "invalid"
//
//	}
//}
//
