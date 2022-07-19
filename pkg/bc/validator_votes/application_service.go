package validator_votes

import (
	"fmt"

	"bjungle-consenso/internal/logger"
	"bjungle-consenso/internal/models"
	"github.com/asaskevich/govalidator"
)

type PortsServerValidatorVotes interface {
	CreateValidatorVotes(id string, lotteryId string, participantsId string, hash string, vote bool) (*ValidatorVotes, int, error)
	UpdateValidatorVotes(id string, lotteryId string, participantsId string, hash string, vote bool) (*ValidatorVotes, int, error)
	DeleteValidatorVotes(id string) (int, error)
	GetValidatorVotesByID(id string) (*ValidatorVotes, int, error)
	GetAllValidatorVotes() ([]*ValidatorVotes, error)
	GetAllValidatorVotesByLotteryID(lotteryID string) ([]*ValidatorVotes, error)
}

type service struct {
	repository ServicesValidatorVotesRepository
	user       *models.User
	txID       string
}

func NewValidatorVotesService(repository ServicesValidatorVotesRepository, user *models.User, TxID string) PortsServerValidatorVotes {
	return &service{repository: repository, user: user, txID: TxID}
}

func (s *service) CreateValidatorVotes(id string, lotteryId string, participantsId string, hash string, vote bool) (*ValidatorVotes, int, error) {
	m := NewValidatorVotes(id, lotteryId, participantsId, hash, vote)
	if valid, err := m.valid(); !valid {
		logger.Error.Println(s.txID, " - don't meet validations:", err)
		return m, 15, err
	}

	if err := s.repository.create(m); err != nil {
		if err.Error() == "ecatch:108" {
			return m, 108, nil
		}
		logger.Error.Println(s.txID, " - couldn't create ValidatorVotes :", err)
		return m, 3, err
	}
	return m, 29, nil
}

func (s *service) UpdateValidatorVotes(id string, lotteryId string, participantsId string, hash string, vote bool) (*ValidatorVotes, int, error) {
	m := NewValidatorVotes(id, lotteryId, participantsId, hash, vote)
	if valid, err := m.valid(); !valid {
		logger.Error.Println(s.txID, " - don't meet validations:", err)
		return m, 15, err
	}
	if err := s.repository.update(m); err != nil {
		logger.Error.Println(s.txID, " - couldn't update ValidatorVotes :", err)
		return m, 18, err
	}
	return m, 29, nil
}

func (s *service) DeleteValidatorVotes(id string) (int, error) {
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

func (s *service) GetValidatorVotesByID(id string) (*ValidatorVotes, int, error) {
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

func (s *service) GetAllValidatorVotes() ([]*ValidatorVotes, error) {
	return s.repository.getAll()
}

func (s *service) GetAllValidatorVotesByLotteryID(lotteryID string) ([]*ValidatorVotes, error) {
	return s.repository.getByLotteryID(lotteryID)
}
