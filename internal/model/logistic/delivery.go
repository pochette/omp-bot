package logistic

import (
	"fmt"
)

type Delivery struct {
	ID          uint64
	AddressFrom string
	AddressTo   string
	Status      DeliveryStatus
}

func NewDelivery(addressFrom string, addressTo string, status DeliveryStatus) *Delivery {

	return &Delivery{
		AddressFrom: addressFrom,
		AddressTo:   addressTo,
		Status:      status}
}
func (d Delivery) String() string {
	return fmt.Sprintf(
		"Delivery #%d: %s -> %s, status: %s",
		d.ID,
		d.AddressFrom,
		d.AddressTo,
		d.Status.String())
}
