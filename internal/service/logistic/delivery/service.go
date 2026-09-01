package delivery

import (
	"fmt"

	"github.com/ozonmp/omp-bot/internal/model/logistic"
)

type Service interface {
	Describe(DeliveryID uint64) (*logistic.Delivery, error)
	List(cursor uint64, limit uint64) ([]logistic.Delivery, error)
	Create(delivery *logistic.Delivery) (uint64, error)
	Update(DeliveryID uint64, Delivery logistic.Delivery) error
	Remove(DeliveryID uint64) (bool, error)
	Get(DeliveryId uint64) (*logistic.Delivery, error)
}

type DummyDeliveryService struct {
	nextId     uint64
	deliveries []logistic.Delivery
}

func (s *DummyDeliveryService) Get(DeliveryId uint64) (*logistic.Delivery, error) {
	for _, delivery := range s.deliveries {
		if delivery.ID == DeliveryId {
			return &delivery, nil
		}
	}
	return nil, fmt.Errorf(
		"delivery with id %d was not found",
		DeliveryId)
}

func (s *DummyDeliveryService) GetNextId() uint64 {
	id := s.nextId
	s.nextId++
	return id
}

func NewDummyDeliveryService() *DummyDeliveryService {
	return &DummyDeliveryService{
		nextId: 11,
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
			{
				ID:          3,
				AddressFrom: "Москва",
				AddressTo:   "Казань",
				Status:      logistic.DeliveryStatusCreated,
			},
			{
				ID:          4,
				AddressFrom: "Санкт-Петербург",
				AddressTo:   "Екатеринбург",
				Status:      logistic.DeliveryStatusInTransit,
			},
			{
				ID:          5,
				AddressFrom: "Новосибирск",
				AddressTo:   "Омск",
				Status:      logistic.DeliveryStatusCreated,
			},
			{
				ID:          6,
				AddressFrom: "Самара",
				AddressTo:   "Уфа",
				Status:      logistic.DeliveryStatusInTransit,
			},
			{
				ID:          7,
				AddressFrom: "Воронеж",
				AddressTo:   "Пермь",
				Status:      logistic.DeliveryStatusCreated,
			},
			{
				ID:          8,
				AddressFrom: "Волгоград",
				AddressTo:   "Саратов",
				Status:      logistic.DeliveryStatusInTransit,
			},
			{
				ID:          9,
				AddressFrom: "Тюмень",
				AddressTo:   "Челябинск",
				Status:      logistic.DeliveryStatusCreated,
			},
			{
				ID:          10,
				AddressFrom: "Иркутск",
				AddressTo:   "Красноярск",
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

func (s *DummyDeliveryService) Create(delivery *logistic.Delivery) (uint64, error) {
	delivery.ID = s.GetNextId()
	s.deliveries = append(
		s.deliveries,
		*delivery,
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
