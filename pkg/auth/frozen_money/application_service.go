package frozen_money

import (
	"fmt"

	"bjungle-consenso/internal/logger"
	"bjungle-consenso/internal/models"
	"github.com/asaskevich/govalidator"
)

type PortsServerFrozenMoney interface {
	CreateFrozenMoney(id string, walletId string, amount int64, lotteryId string) (*FrozenMoney, int, error)
	UpdateFrozenMoney(id string, walletId string, amount int64, lotteryId string) (*FrozenMoney, int, error)
	DeleteFrozenMoney(id string) (int, error)
	GetFrozenMoneyByID(id string) (*FrozenMoney, int, error)
	GetAllFrozenMoney() ([]*FrozenMoney, error)
}

type service struct {
	repository ServicesFrozenMoneyRepository
	user       *models.User
	txID       string
}

func NewFrozenMoneyService(repository ServicesFrozenMoneyRepository, user *models.User, TxID string) PortsServerFrozenMoney {
	return &service{repository: repository, user: user, txID: TxID}
}

func (s *service) CreateFrozenMoney(id string, walletId string, amount int64, lotteryId string) (*FrozenMoney, int, error) {
	m := NewFrozenMoney(id, walletId, amount, lotteryId)
	if valid, err := m.valid(); !valid {
		logger.Error.Println(s.txID, " - don't meet validations:", err)
		return m, 15, err
	}

	if err := s.repository.create(m); err != nil {
		if err.Error() == "ecatch:108" {
			return m, 108, nil
		}
		logger.Error.Println(s.txID, " - couldn't create FrozenMoney :", err)
		return m, 3, err
	}
	return m, 29, nil
}

func (s *service) UpdateFrozenMoney(id string, walletId string, amount int64, lotteryId string) (*FrozenMoney, int, error) {
	m := NewFrozenMoney(id, walletId, amount, lotteryId)
	if valid, err := m.valid(); !valid {
		logger.Error.Println(s.txID, " - don't meet validations:", err)
		return m, 15, err
	}
	if err := s.repository.update(m); err != nil {
		logger.Error.Println(s.txID, " - couldn't update FrozenMoney :", err)
		return m, 18, err
	}
	return m, 29, nil
}

func (s *service) DeleteFrozenMoney(id string) (int, error) {
	if !govalidator.IsUUID(id) {
		logger.Error.Println(s.txID, " - don't meet validations:", fmt.Errorf("id isn't uuid"))
		return 15, fmt.Errorf("id isn't uuid")
	}

	if err := s.repository.delete(id); err != nil {
		if err.Error() == "ecatch:108" {
			return 108, nil
		}
		logger.Error.Println(s.txID, " - couldn't update row:", err)
		return 20, err
	}
	return 28, nil
}

func (s *service) GetFrozenMoneyByID(id string) (*FrozenMoney, int, error) {
	if !govalidator.IsUUID(id) {
		logger.Error.Println(s.txID, " - don't meet validations:", fmt.Errorf("id isn't uuid"))
		return nil, 15, fmt.Errorf("id isn't uuid")
	}
	m, err := s.repository.getByID(id)
	if err != nil {
		logger.Error.Println(s.txID, " - couldn`t getByID row:", err)
		return nil, 22, err
	}
	return m, 29, nil
}

func (s *service) GetAllFrozenMoney() ([]*FrozenMoney, error) {
	return s.repository.getAll()
}
