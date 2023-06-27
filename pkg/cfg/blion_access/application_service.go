package blion_access

import (
	"fmt"

	"bjungle-consenso/internal/logger"
	"bjungle-consenso/internal/models"
	"github.com/asaskevich/govalidator"
)

type PortsServerBLionAccess interface {
	CreateBLionAccess(id string, key string, status string) (*BLionAccess, int, error)
	UpdateBLionAccess(id string, key string, status string) (*BLionAccess, int, error)
	DeleteBLionAccess(id string) (int, error)
	GetBLionAccessByID(id string) (*BLionAccess, int, error)
	GetAllBLionAccess() ([]*BLionAccess, error)
	GetBLionAccessByIds(ids []string) ([]*BLionAccess, error)
}

type service struct {
	repository ServicesBLionAccessRepository
	user       *models.User
	txID       string
}

func NewBLionAccessService(repository ServicesBLionAccessRepository, user *models.User, TxID string) PortsServerBLionAccess {
	return &service{repository: repository, user: user, txID: TxID}
}

func (s *service) CreateBLionAccess(id string, key string, status string) (*BLionAccess, int, error) {
	m := NewBLionAccess(id, key, status)
	if valid, err := m.valid(); !valid {
		logger.Error.Println(s.txID, " - don't meet validations:", err)
		return m, 15, err
	}

	if err := s.repository.create(m); err != nil {
		if err.Error() == "ecatch:108" {
			return m, 108, nil
		}
		logger.Error.Println(s.txID, " - couldn't create BLionAccess :", err)
		return m, 3, err
	}
	return m, 29, nil
}

func (s *service) UpdateBLionAccess(id string, key string, status string) (*BLionAccess, int, error) {
	m := NewBLionAccess(id, key, status)
	if valid, err := m.valid(); !valid {
		logger.Error.Println(s.txID, " - don't meet validations:", err)
		return m, 15, err
	}
	if err := s.repository.update(m); err != nil {
		logger.Error.Println(s.txID, " - couldn't update BLionAccess :", err)
		return m, 18, err
	}
	return m, 29, nil
}

func (s *service) DeleteBLionAccess(id string) (int, error) {
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

func (s *service) GetBLionAccessByID(id string) (*BLionAccess, int, error) {
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

func (s *service) GetAllBLionAccess() ([]*BLionAccess, error) {
	return s.repository.getAll()
}

func (s *service) GetBLionAccessByIds(ids []string) ([]*BLionAccess, error) {
	return s.repository.getByIDs(ids)
}
