package delivery

import (
	"fmt"

	"github.com/ozonmp/omp-bot/internal/model/logistic"
)

type Service interface {
	Describe(DeliveryID uint64) (*logistic.Delivery, error)
	List(cursor uint64, limit uint64) ([]logistic.Delivery, error)
	Create(logistic.Delivery) (uint64, error)
	Update(DeliveryID uint64, Delivery logistic.Delivery) error
	Remove(DeliveryID uint64) (bool, error)
}

type DummyDeliveryService struct {
	nextId     uint64
	deliveries []logistic.Delivery
}

func (s *DummyDeliveryService) getNextId() uint64 {
	id := s.nextId
	s.nextId++
	return id
}

func NewDummyDeliveryService() *DummyDeliveryService {
	return &DummyDeliveryService{
		nextId: 3,
		deliveries: []logistic.Delivery{
			{
				ID:          1,
				AddressFrom: "Краснодар",
				AddressTo:   "Томск",
				Status:      logistic.DeliveryStatusCreated,
			},
			{
				ID:          2,
				AddressFrom: "Ростов",
				AddressTo:   "Нижний Новгород",
				Status:      logistic.DeliveryStatusInTransit,
			},
		},
	}
}

func (s *DummyDeliveryService) Describe(DeliveryID uint64) (*logistic.Delivery, error) {
	for index := range s.deliveries {
		if s.deliveries[index].ID == DeliveryID {
			return &s.deliveries[index], nil
		}
	}

	return nil, fmt.Errorf("delivery with id: %d was not found", DeliveryID)
}

func (s *DummyDeliveryService) List(cursor uint64, limit uint64) ([]logistic.Delivery, error) {
	total := uint64(len(s.deliveries))
	if limit == 0 || cursor >= total {
		return []logistic.Delivery{}, nil
	}
	remaining := total - cursor
	if limit > remaining {
		limit = remaining
	}

	return s.deliveries[cursor : cursor+limit], nil

}

func (s *DummyDeliveryService) Create(delivery logistic.Delivery) (uint64, error) {
	delivery.ID = s.getNextId()
	s.deliveries = append(
		s.deliveries,
		delivery,
	)
	return delivery.ID, nil
}

func (s *DummyDeliveryService) Update(DeliveryID uint64, Delivery logistic.Delivery) error {
	for index := range s.deliveries {
		if s.deliveries[index].ID == DeliveryID {
			Delivery.ID = DeliveryID
			s.deliveries[index] = Delivery
			return nil
		}
	}
	return fmt.Errorf("delivery with id %d was not found", DeliveryID)
}

func (s *DummyDeliveryService) Remove(DeliveryID uint64) (bool, error) {
	for index := range s.deliveries {
		if s.deliveries[index].ID == DeliveryID {
			s.deliveries = append(
				s.deliveries[:index],
				s.deliveries[index+1:]...)
			return true, nil

		}
	}
	return false,
		fmt.Errorf("delivered with id:%d was not found", DeliveryID)
}
