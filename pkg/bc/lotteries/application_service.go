package lotteries

import (
	"fmt"

	"bjungle-consenso/internal/logger"
	"bjungle-consenso/internal/models"
	"github.com/asaskevich/govalidator"
	"time"
)

type PortsServerLottery interface {
	CreateLottery(id string, blockId int64, registrationStartDate time.Time, registrationEndDate *time.Time, lotteryStartDate *time.Time, lotteryEndDate *time.Time, processEndDate *time.Time, processStatus int) (*Lottery, int, error)
	UpdateLottery(id string, blockId int64, registrationStartDate time.Time, registrationEndDate *time.Time, lotteryStartDate *time.Time, lotteryEndDate *time.Time, processEndDate *time.Time, processStatus int) (*Lottery, int, error)
	DeleteLottery(id string) (int, error)
	GetLotteryByID(id string) (*Lottery, int, error)
	GetAllLottery() ([]*Lottery, error)
	GetLotteryActive() (*Lottery, int, error)
	GetLotteryActiveForMined() (*Lottery, int, error)
}

type service struct {
	repository ServicesLotteryRepository
	user       *models.User
	txID       string
}

func NewLotteryService(repository ServicesLotteryRepository, user *models.User, TxID string) PortsServerLottery {
	return &service{repository: repository, user: user, txID: TxID}
}

func (s *service) CreateLottery(id string, blockId int64, registrationStartDate time.Time, registrationEndDate *time.Time, lotteryStartDate *time.Time, lotteryEndDate *time.Time, processEndDate *time.Time, processStatus int) (*Lottery, int, error) {
	m := NewLottery(id, blockId, registrationStartDate, registrationEndDate, lotteryStartDate, lotteryEndDate, processEndDate, processStatus)
	if valid, err := m.valid(); !valid {
		logger.Error.Println(s.txID, " - don't meet validations:", err)
		return m, 15, err
	}

	if err := s.repository.create(m); err != nil {
		if err.Error() == "ecatch:108" {
			return m, 108, nil
		}
		logger.Error.Println(s.txID, " - couldn't create Lottery :", err)
		return m, 3, err
	}
	return m, 29, nil
}

func (s *service) UpdateLottery(id string, blockId int64, registrationStartDate time.Time, registrationEndDate *time.Time, lotteryStartDate *time.Time, lotteryEndDate *time.Time, processEndDate *time.Time, processStatus int) (*Lottery, int, error) {
	m := NewLottery(id, blockId, registrationStartDate, registrationEndDate, lotteryStartDate, lotteryEndDate, processEndDate, processStatus)
	if valid, err := m.valid(); !valid {
		logger.Error.Println(s.txID, " - don't meet validations:", err)
		return m, 15, err
	}
	if err := s.repository.update(m); err != nil {
		logger.Error.Println(s.txID, " - couldn't update Lottery :", err)
		return m, 18, err
	}
	return m, 29, nil
}

func (s *service) DeleteLottery(id string) (int, error) {
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

func (s *service) GetLotteryByID(id string) (*Lottery, int, error) {
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

func (s *service) GetAllLottery() ([]*Lottery, error) {
	return s.repository.getAll()
}

func (s *service) GetLotteryActive() (*Lottery, int, error) {
	m, err := s.repository.getActive()
	if err != nil {
		logger.Error.Println(s.txID, " - couldn`t getActive, error: ", err)
		return nil, 22, err
	}
	return m, 29, nil
}

func (s *service) GetLotteryActiveForMined() (*Lottery, int, error) {
	m, err := s.repository.getActiveForMined()
	if err != nil {
		logger.Error.Println(s.txID, " - couldn`t getActive, error: ", err)
		return nil, 22, err
	}
	return m, 29, nil
}
