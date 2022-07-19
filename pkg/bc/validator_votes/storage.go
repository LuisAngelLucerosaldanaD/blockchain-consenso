package validator_votes

import (
	"github.com/jmoiron/sqlx"

	"bjungle-consenso/internal/logger"
	"bjungle-consenso/internal/models"
)

const (
	Postgresql = "postgres"
)

type ServicesValidatorVotesRepository interface {
	create(m *ValidatorVotes) error
	update(m *ValidatorVotes) error
	delete(id string) error
	getByID(id string) (*ValidatorVotes, error)
	getAll() ([]*ValidatorVotes, error)
	getByLotteryID(lotteryID string) ([]*ValidatorVotes, error)
}

func FactoryStorage(db *sqlx.DB, user *models.User, txID string) ServicesValidatorVotesRepository {
	var s ServicesValidatorVotesRepository
	engine := db.DriverName()
	switch engine {
	case Postgresql:
		return newValidatorVotesPsqlRepository(db, user, txID)
	default:
		logger.Error.Println("el motor de base de datos no está implementado.", engine)
	}
	return s
}
