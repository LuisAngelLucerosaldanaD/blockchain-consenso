package lottery_table

import (
	"fmt"

	"bjungle-consenso/internal/logger"
	"bjungle-consenso/internal/models"
	"github.com/asaskevich/govalidator"
	"time"
)

type PortsServerLotteryTable interface {
	CreateLotteryTable(id string, blockId int64, registrationStartDate time.Time, registrationEndDate time.Time, lotteryStartDate time.Time, lotteryEndDate time.Time, processEndDate time.Time, processStatus int) (*LotteryTable, int, error)
	UpdateLotteryTable(id string, blockId int64, registrationStartDate time.Time, registrationEndDate time.Time, lotteryStartDate time.Time, lotteryEndDate time.Time, processEndDate time.Time, processStatus int) (*LotteryTable, int, error)
	DeleteLotteryTable(id string) (int, error)
	GetLotteryTableByID(id string) (*LotteryTable, int, error)
	GetAllLotteryTable() ([]*LotteryTable, error)
	GetLotteryActive() (*LotteryTable, int, error)
}

type service struct {
	repository ServicesLotteryTableRepository
	user       *models.User
	txID       string
}

func NewLotteryTableService(repository ServicesLotteryTableRepository, user *models.User, TxID string) PortsServerLotteryTable {
	return &service{repository: repository, user: user, txID: TxID}
}

func (s *service) CreateLotteryTable(id string, blockId int64, registrationStartDate time.Time, registrationEndDate time.Time, lotteryStartDate time.Time, lotteryEndDate time.Time, processEndDate time.Time, processStatus int) (*LotteryTable, int, error) {
	m := NewLotteryTable(id, blockId, registrationStartDate, registrationEndDate, lotteryStartDate, lotteryEndDate, processEndDate, processStatus)
	if valid, err := m.valid(); !valid {
		logger.Error.Println(s.txID, " - don't meet validations:", err)
		return m, 15, err
	}

	if err := s.repository.create(m); err != nil {
		if err.Error() == "ecatch:108" {
			return m, 108, nil
		}
		logger.Error.Println(s.txID, " - couldn't create LotteryTable :", err)
		return m, 3, err
	}
	return m, 29, nil
}

func (s *service) UpdateLotteryTable(id string, blockId int64, registrationStartDate time.Time, registrationEndDate time.Time, lotteryStartDate time.Time, lotteryEndDate time.Time, processEndDate time.Time, processStatus int) (*LotteryTable, int, error) {
	m := NewLotteryTable(id, blockId, registrationStartDate, registrationEndDate, lotteryStartDate, lotteryEndDate, processEndDate, processStatus)
	if valid, err := m.valid(); !valid {
		logger.Error.Println(s.txID, " - don't meet validations:", err)
		return m, 15, err
	}
	if err := s.repository.update(m); err != nil {
		logger.Error.Println(s.txID, " - couldn't update LotteryTable :", err)
		return m, 18, err
	}
	return m, 29, nil
}

func (s *service) DeleteLotteryTable(id string) (int, error) {
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

func (s *service) GetLotteryTableByID(id string) (*LotteryTable, int, error) {
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

func (s *service) GetAllLotteryTable() ([]*LotteryTable, error) {
	return s.repository.getAll()
}

func (s *service) GetLotteryActive() (*LotteryTable, int, error) {
	m, err := s.repository.getActive()
	if err != nil {
		logger.Error.Println(s.txID, " - couldn`t getActive, error: ", err)
		return nil, 22, err
	}
	return m, 29, nil
}
