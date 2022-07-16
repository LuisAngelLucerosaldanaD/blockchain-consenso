package reward_table

import (
	"fmt"

	"bjungle-consenso/internal/logger"
	"bjungle-consenso/internal/models"
	"github.com/asaskevich/govalidator"
)

type PortsServerRewardTable interface {
	CreateRewardTable(id string, lotteryId string, idWallet string, amount int64, blockId int64) (*RewardTable, int, error)
	UpdateRewardTable(id string, lotteryId string, idWallet string, amount int64, blockId int64) (*RewardTable, int, error)
	DeleteRewardTable(id string) (int, error)
	GetRewardTableByID(id string) (*RewardTable, int, error)
	GetAllRewardTable() ([]*RewardTable, error)
}

type service struct {
	repository ServicesRewardTableRepository
	user       *models.User
	txID       string
}

func NewRewardTableService(repository ServicesRewardTableRepository, user *models.User, TxID string) PortsServerRewardTable {
	return &service{repository: repository, user: user, txID: TxID}
}

func (s *service) CreateRewardTable(id string, lotteryId string, idWallet string, amount int64, blockId int64) (*RewardTable, int, error) {
	m := NewRewardTable(id, lotteryId, idWallet, amount, blockId)
	if valid, err := m.valid(); !valid {
		logger.Error.Println(s.txID, " - don't meet validations:", err)
		return m, 15, err
	}

	if err := s.repository.create(m); err != nil {
		if err.Error() == "ecatch:108" {
			return m, 108, nil
		}
		logger.Error.Println(s.txID, " - couldn't create RewardTable :", err)
		return m, 3, err
	}
	return m, 29, nil
}

func (s *service) UpdateRewardTable(id string, lotteryId string, idWallet string, amount int64, blockId int64) (*RewardTable, int, error) {
	m := NewRewardTable(id, lotteryId, idWallet, amount, blockId)
	if valid, err := m.valid(); !valid {
		logger.Error.Println(s.txID, " - don't meet validations:", err)
		return m, 15, err
	}
	if err := s.repository.update(m); err != nil {
		logger.Error.Println(s.txID, " - couldn't update RewardTable :", err)
		return m, 18, err
	}
	return m, 29, nil
}

func (s *service) DeleteRewardTable(id string) (int, error) {
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

func (s *service) GetRewardTableByID(id string) (*RewardTable, int, error) {
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

func (s *service) GetAllRewardTable() ([]*RewardTable, error) {
	return s.repository.getAll()
}
