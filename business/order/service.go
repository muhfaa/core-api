package order

import (
	"core-data/business"
	"core-data/business/kerusakan"
	"core-data/business/teknisi"
	"errors"
)

type Service interface {
	InsertOrder(orderSpec OrderSpecRequest) error
	UpdateOrderStatus(orderSpec UpdateOrderStatusSpec) (*UpdateOrderStatus, error)
}

type service struct {
	orderRepository     Repository
	teknisiRepository   teknisi.Repository
	kerusakanRepository kerusakan.Repository
}

func NewService(orderRepository Repository, teknisiRepository teknisi.Repository, kerusakanRepository kerusakan.Repository) Service {

	return &service{
		orderRepository,
		teknisiRepository,
		kerusakanRepository,
	}
}

func (s *service) InsertOrder(orderSpec OrderSpecRequest) error {
	var teknisiVersion int

	allTeknisi, err := s.teknisiRepository.GetListTeknisi()
	if err != nil || allTeknisi == nil {
		return err
	}

	kerusakan, err := s.kerusakanRepository.GetKerusakan(orderSpec.KerusakanID)
	if err != nil || kerusakan == nil {
		return err
	}

	// get data teknisi id
	var min int
	for _, data := range allTeknisi {
		if data.Platform == orderSpec.JenisPlatform && min == 0 {
			min = data.JumlahAntrian
		}
		if data.Platform == orderSpec.JenisPlatform && data.JumlahAntrian < min {
			min = data.JumlahAntrian
		}
	}

	for _, teknisi := range allTeknisi {
		if teknisi.Specialist == orderSpec.JenisHP && teknisi.JumlahAntrian < 3 {
			orderSpec.TeknisiID = teknisi.ID
			teknisiVersion = teknisi.Version
		} else if teknisi.Platform == orderSpec.JenisPlatform && teknisi.JumlahAntrian < 3 {
			orderSpec.TeknisiID = teknisi.ID
			teknisiVersion = teknisi.Version
		} else if teknisi.Platform == orderSpec.JenisPlatform && teknisi.JumlahAntrian >= 3 && teknisi.JumlahAntrian == min {
			orderSpec.TeknisiID = teknisi.ID
			teknisiVersion = teknisi.Version
		}

	}

	order := NewOrder(
		orderSpec.TeknisiID,
		orderSpec.KerusakanID,
		orderSpec.JenisHP,
		orderSpec.JenisPlatform,
		StateAntrian,
		kerusakan.LamaPengerjaan,
		orderSpec.Version,
		teknisiVersion,
	)

	err = s.orderRepository.InsertOrder(order)
	if err != nil && err != business.ErrNotFound {
		return err
	}

	updateAntrian := teknisi.NewUpdateJumlahAntrian(
		order.TeknisiID,
		teknisiVersion,
	)

	isAdd, err := s.teknisiRepository.AddAntrian(updateAntrian)
	if err != nil || !isAdd {
		return err
	}

	return nil
}

func (s *service) UpdateOrderStatus(orderSpec UpdateOrderStatusSpec) (*UpdateOrderStatus, error) {

	order, err := s.orderRepository.FindOrderByID(orderSpec.ID)
	if err != nil || order == nil {
		return nil, err
	}

	if order.StatusService == StateSelesai {
		return nil, errors.New("Order was finished")
	}

	updateSpec := NewUpdateStatus(
		orderSpec.ID,
		order.TeknisiID,
		order.StatusService,
		orderSpec.Version,
	)

	isUpdated, err := s.orderRepository.UpdateOrderStatus(updateSpec)
	if err != nil || !isUpdated {
		return nil, err
	}

	updateAntrian := teknisi.NewUpdateJumlahAntrian(
		order.TeknisiID,
		order.VersionTeknisi,
	)

	if updateSpec.StatusService == StateSelesai {
		isErase, err := s.teknisiRepository.EraseAntrian(updateAntrian)
		if err != nil || !isErase {
			return nil, err
		}
	}

	return &updateSpec, nil
}
