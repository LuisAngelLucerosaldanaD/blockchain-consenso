package participants_table

import (
	"fmt"

	"bjungle-consenso/internal/logger"
	"bjungle-consenso/internal/models"
	"github.com/asaskevich/govalidator"
)

type PortsServerParticipantsTable interface {
	CreateParticipantsTable(id string, lotteryId string, walletId string, amount int64, accepted bool, typeCharge int, returned bool) (*ParticipantsTable, int, error)
	UpdateParticipantsTable(id string, lotteryId string, walletId string, amount int64, accepted bool, typeCharge int, returned bool) (*ParticipantsTable, int, error)
	DeleteParticipantsTable(id string) (int, error)
	GetParticipantsTableByID(id string) (*ParticipantsTable, int, error)
	GetAllParticipantsTable() ([]*ParticipantsTable, error)
	GetParticipantsTableByWalletID(walletId string) (*ParticipantsTable, int, error)
}

type service struct {
	repository ServicesParticipantsTableRepository
	user       *models.User
	txID       string
}

func NewParticipantsTableService(repository ServicesParticipantsTableRepository, user *models.User, TxID string) PortsServerParticipantsTable {
	return &service{repository: repository, user: user, txID: TxID}
}

func (s *service) CreateParticipantsTable(id string, lotteryId string, walletId string, amount int64, accepted bool, typeCharge int, returned bool) (*ParticipantsTable, int, error) {
	m := NewParticipantsTable(id, lotteryId, walletId, amount, accepted, typeCharge, returned)
	if valid, err := m.valid(); !valid {
		logger.Error.Println(s.txID, " - don't meet validations:", err)
		return m, 15, err
	}

	if err := s.repository.create(m); err != nil {
		if err.Error() == "ecatch:108" {
			return m, 108, nil
		}
		logger.Error.Println(s.txID, " - couldn't create ParticipantsTable :", err)
		return m, 3, err
	}
	return m, 29, nil
}

func (s *service) UpdateParticipantsTable(id string, lotteryId string, walletId string, amount int64, accepted bool, typeCharge int, returned bool) (*ParticipantsTable, int, error) {
	m := NewParticipantsTable(id, lotteryId, walletId, amount, accepted, typeCharge, returned)
	if valid, err := m.valid(); !valid {
		logger.Error.Println(s.txID, " - don't meet validations:", err)
		return m, 15, err
	}
	if err := s.repository.update(m); err != nil {
		logger.Error.Println(s.txID, " - couldn't update ParticipantsTable :", err)
		return m, 18, err
	}
	return m, 29, nil
}

func (s *service) DeleteParticipantsTable(id string) (int, error) {
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

func (s *service) GetParticipantsTableByID(id string) (*ParticipantsTable, int, error) {
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

func (s *service) GetParticipantsTableByWalletID(walletId string) (*ParticipantsTable, int, error) {
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

func (s *service) GetAllParticipantsTable() ([]*ParticipantsTable, error) {
	return s.repository.getAll()
}
