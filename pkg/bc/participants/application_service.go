package participants

import (
	"fmt"

	"bjungle-consenso/internal/logger"
	"bjungle-consenso/internal/models"
	"github.com/asaskevich/govalidator"
)

type PortsServerParticipants interface {
	CreateParticipants(id string, lotteryId string, walletId string, amount int64, accepted bool, typeCharge int, returned bool) (*Participants, int, error)
	UpdateParticipants(id string, lotteryId string, walletId string, amount int64, accepted bool, typeCharge int, returned bool) (*Participants, int, error)
	DeleteParticipants(id string) (int, error)
	GetParticipantsByID(id string) (*Participants, int, error)
	GetAllParticipants() ([]*Participants, error)
	GetParticipantsByWalletID(walletId string) (*Participants, int, error)
	GetParticipantsByLotteryID(lotteryID string) ([]*Participants, int, error)
	GetParticipantsByWalletIDAndLotteryID(walletId string, lotteryID string) (*Participants, int, error)
}

type service struct {
	repository ServicesParticipantsRepository
	user       *models.User
	txID       string
}

func NewParticipantsService(repository ServicesParticipantsRepository, user *models.User, TxID string) PortsServerParticipants {
	return &service{repository: repository, user: user, txID: TxID}
}

func (s *service) CreateParticipants(id string, lotteryId string, walletId string, amount int64, accepted bool, typeCharge int, returned bool) (*Participants, int, error) {
	m := NewParticipants(id, lotteryId, walletId, amount, accepted, typeCharge, returned)
	if valid, err := m.valid(); !valid {
		logger.Error.Println(s.txID, " - don't meet validations:", err)
		return m, 15, err
	}

	if err := s.repository.create(m); err != nil {
		if err.Error() == "ecatch:108" {
			return m, 108, nil
		}
		logger.Error.Println(s.txID, " - couldn't create Participants :", err)
		return m, 3, err
	}
	return m, 29, nil
}

func (s *service) UpdateParticipants(id string, lotteryId string, walletId string, amount int64, accepted bool, typeCharge int, returned bool) (*Participants, int, error) {
	m := NewParticipants(id, lotteryId, walletId, amount, accepted, typeCharge, returned)
	if valid, err := m.valid(); !valid {
		logger.Error.Println(s.txID, " - don't meet validations:", err)
		return m, 15, err
	}
	if err := s.repository.update(m); err != nil {
		logger.Error.Println(s.txID, " - couldn't update Participants :", err)
		return m, 18, err
	}
	return m, 29, nil
}

func (s *service) DeleteParticipants(id string) (int, error) {
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

func (s *service) GetParticipantsByID(id string) (*Participants, int, error) {
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

func (s *service) GetParticipantsByWalletID(walletId string) (*Participants, int, error) {
	if !govalidator.IsUUID(walletId) {
		logger.Error.Println(s.txID, " - don't meet validations:", fmt.Errorf("wallet id isn't uuid"))
		return nil, 15, fmt.Errorf("wallet id isn't uuid")
	}
	m, err := s.repository.getByWalletID(walletId)
	if err != nil {
		logger.Error.Println(s.txID, " - couldn`t getByWalletID row:", err)
		return nil, 22, err
	}
	return m, 29, nil
}

func (s *service) GetParticipantsByWalletIDAndLotteryID(walletId string, lotteryID string) (*Participants, int, error) {
	m, err := s.repository.getByWalletIDAndLotteryID(walletId, lotteryID)
	if err != nil {
		logger.Error.Println(s.txID, " - couldn`t getByWalletID row:", err)
		return nil, 22, err
	}
	return m, 29, nil
}

func (s *service) GetParticipantsByLotteryID(lotteryID string) ([]*Participants, int, error) {
	if !govalidator.IsUUID(lotteryID) {
		logger.Error.Println(s.txID, " - don't meet validations:", fmt.Errorf("lottery id isn't uuid"))
		return nil, 15, fmt.Errorf("lottery id isn't uuid")
	}
	m, err := s.repository.getByLotteryID(lotteryID)
	if err != nil {
		logger.Error.Println(s.txID, " - couldn`t getByWalletID row:", err)
		return nil, 22, err
	}
	return m, 29, nil
}

func (s *service) GetAllParticipants() ([]*Participants, error) {
	return s.repository.getAll()
}
