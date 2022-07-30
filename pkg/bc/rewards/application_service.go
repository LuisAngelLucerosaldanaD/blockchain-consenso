package rewards

import (
	"fmt"

	"bjungle-consenso/internal/logger"
	"bjungle-consenso/internal/models"
	"github.com/asaskevich/govalidator"
)

type PortsServerRewards interface {
	CreateReward(id string, lotteryId string, idWallet string, amount float64) (*Reward, int, error)
	UpdateReward(id string, lotteryId string, idWallet string, amount float64) (*Reward, int, error)
	DeleteReward(id string) (int, error)
	GetRewardByID(id string) (*Reward, int, error)
	GetAllReward() ([]*Reward, error)
}

type service struct {
	repository ServicesRewardsRepository
	user       *models.User
	txID       string
}

func NewRewardsService(repository ServicesRewardsRepository, user *models.User, TxID string) PortsServerRewards {
	return &service{repository: repository, user: user, txID: TxID}
}

func (s *service) CreateReward(id string, lotteryId string, idWallet string, amount float64) (*Reward, int, error) {
	m := NewReward(id, lotteryId, idWallet, amount)
	if valid, err := m.valid(); !valid {
		logger.Error.Println(s.txID, " - don't meet validations:", err)
		return m, 15, err
	}

	if err := s.repository.create(m); err != nil {
		if err.Error() == "ecatch:108" {
			return m, 108, nil
		}
		logger.Error.Println(s.txID, " - couldn't create Reward :", err)
		return m, 3, err
	}
	return m, 29, nil
}

func (s *service) UpdateReward(id string, lotteryId string, idWallet string, amount float64) (*Reward, int, error) {
	m := NewReward(id, lotteryId, idWallet, amount)
	if valid, err := m.valid(); !valid {
		logger.Error.Println(s.txID, " - don't meet validations:", err)
		return m, 15, err
	}
	if err := s.repository.update(m); err != nil {
		logger.Error.Println(s.txID, " - couldn't update Reward :", err)
		return m, 18, err
	}
	return m, 29, nil
}

func (s *service) DeleteReward(id string) (int, error) {
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

func (s *service) GetRewardByID(id string) (*Reward, int, error) {
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

func (s *service) GetAllReward() ([]*Reward, error) {
	return s.repository.getAll()
}
