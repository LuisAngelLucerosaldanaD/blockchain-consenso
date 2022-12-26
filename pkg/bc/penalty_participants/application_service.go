package penalty_participants

import (
	"bjungle-consenso/internal/logger"
	"bjungle-consenso/internal/models"
	"fmt"
	"github.com/asaskevich/govalidator"
)

type PortsServerPenaltyParticipants interface {
	CreatePenaltyParticipants(id string, lotteryId string, participantsId string, amount float64, penaltyPercentage float64) (*PenaltyParticipants, int, error)
	UpdatePenaltyParticipants(id string, lotteryId string, participantsId string, amount float64, penaltyPercentage float64) (*PenaltyParticipants, int, error)
	DeletePenaltyParticipants(id string) (int, error)
	GetPenaltyParticipantsByID(id string) (*PenaltyParticipants, int, error)
	GetAllPenaltyParticipants() ([]*PenaltyParticipants, error)
	GetAllPenaltyParticipantsByWalletID(walletID string) ([]*PenaltyParticipants, error)
}

type service struct {
	repository ServicesPenaltyParticipantsRepository
	user       *models.User
	txID       string
}

func NewPenaltyParticipantsService(repository ServicesPenaltyParticipantsRepository, user *models.User, TxID string) PortsServerPenaltyParticipants {
	return &service{repository: repository, user: user, txID: TxID}
}

func (s *service) CreatePenaltyParticipants(id string, lotteryId string, participantsId string, amount float64, penaltyPercentage float64) (*PenaltyParticipants, int, error) {
	m := NewPenaltyParticipants(id, lotteryId, participantsId, amount, penaltyPercentage)
	if valid, err := m.valid(); !valid {
		logger.Error.Println(s.txID, " - don't meet validations:", err)
		return m, 15, err
	}

	if err := s.repository.create(m); err != nil {
		if err.Error() == "ecatch:108" {
			return m, 108, nil
		}
		logger.Error.Println(s.txID, " - couldn't create PenaltyParticipants :", err)
		return m, 3, err
	}
	return m, 29, nil
}

func (s *service) UpdatePenaltyParticipants(id string, lotteryId string, participantsId string, amount float64, penaltyPercentage float64) (*PenaltyParticipants, int, error) {
	m := NewPenaltyParticipants(id, lotteryId, participantsId, amount, penaltyPercentage)
	if valid, err := m.valid(); !valid {
		logger.Error.Println(s.txID, " - don't meet validations:", err)
		return m, 15, err
	}
	if err := s.repository.update(m); err != nil {
		logger.Error.Println(s.txID, " - couldn't update PenaltyParticipants :", err)
		return m, 18, err
	}
	return m, 29, nil
}

func (s *service) DeletePenaltyParticipants(id string) (int, error) {
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

func (s *service) GetPenaltyParticipantsByID(id string) (*PenaltyParticipants, int, error) {
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

func (s *service) GetAllPenaltyParticipants() ([]*PenaltyParticipants, error) {
	return s.repository.getAll()
}

func (s *service) GetAllPenaltyParticipantsByWalletID(walletID string) ([]*PenaltyParticipants, error) {
	return s.repository.getByWalletID(walletID)
}
