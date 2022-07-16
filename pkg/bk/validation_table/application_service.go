package validation_table

import (
	"fmt"

	"bjungle-consenso/internal/logger"
	"bjungle-consenso/internal/models"
	"github.com/asaskevich/govalidator"
)

type PortsServerValidationTable interface {
	CreateValidationTable(id string, lotteryId string, walletId string, participantsId string, vote bool) (*ValidationTable, int, error)
	UpdateValidationTable(id string, lotteryId string, walletId string, participantsId string, vote bool) (*ValidationTable, int, error)
	DeleteValidationTable(id string) (int, error)
	GetValidationTableByID(id string) (*ValidationTable, int, error)
	GetAllValidationTable() ([]*ValidationTable, error)
}

type service struct {
	repository ServicesValidationTableRepository
	user       *models.User
	txID       string
}

func NewValidationTableService(repository ServicesValidationTableRepository, user *models.User, TxID string) PortsServerValidationTable {
	return &service{repository: repository, user: user, txID: TxID}
}

func (s *service) CreateValidationTable(id string, lotteryId string, walletId string, participantsId string, vote bool) (*ValidationTable, int, error) {
	m := NewValidationTable(id, lotteryId, walletId, participantsId, vote)
	if valid, err := m.valid(); !valid {
		logger.Error.Println(s.txID, " - don't meet validations:", err)
		return m, 15, err
	}

	if err := s.repository.create(m); err != nil {
		if err.Error() == "ecatch:108" {
			return m, 108, nil
		}
		logger.Error.Println(s.txID, " - couldn't create ValidationTable :", err)
		return m, 3, err
	}
	return m, 29, nil
}

func (s *service) UpdateValidationTable(id string, lotteryId string, walletId string, participantsId string, vote bool) (*ValidationTable, int, error) {
	m := NewValidationTable(id, lotteryId, walletId, participantsId, vote)
	if valid, err := m.valid(); !valid {
		logger.Error.Println(s.txID, " - don't meet validations:", err)
		return m, 15, err
	}
	if err := s.repository.update(m); err != nil {
		logger.Error.Println(s.txID, " - couldn't update ValidationTable :", err)
		return m, 18, err
	}
	return m, 29, nil
}

func (s *service) DeleteValidationTable(id string) (int, error) {
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

func (s *service) GetValidationTableByID(id string) (*ValidationTable, int, error) {
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

func (s *service) GetAllValidationTable() ([]*ValidationTable, error) {
	return s.repository.getAll()
}
