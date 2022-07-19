package participants

import (
	"github.com/jmoiron/sqlx"

	"bjungle-consenso/internal/logger"
	"bjungle-consenso/internal/models"
)

const (
	Postgresql = "postgres"
)

type ServicesParticipantsRepository interface {
	create(m *Participants) error
	update(m *Participants) error
	delete(id string) error
	getByID(id string) (*Participants, error)
	getAll() ([]*Participants, error)
	getByWalletID(walletId string) (*Participants, error)
	getByLotteryID(lotteryId string) ([]*Participants, error)
	getByWalletIDAndLotteryID(walletId string, lotteryId string) (*Participants, error)
}

func FactoryStorage(db *sqlx.DB, user *models.User, txID string) ServicesParticipantsRepository {
	var s ServicesParticipantsRepository
	engine := db.DriverName()
	switch engine {
	case Postgresql:
		return newParticipantsPsqlRepository(db, user, txID)
	default:
		logger.Error.Println("el motor de base de datos no está implementado.", engine)
	}
	return s
}
