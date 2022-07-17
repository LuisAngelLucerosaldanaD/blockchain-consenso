package participants_table

import (
	"github.com/jmoiron/sqlx"

	"bjungle-consenso/internal/logger"
	"bjungle-consenso/internal/models"
)

const (
	Postgresql = "postgres"
)

type ServicesParticipantsTableRepository interface {
	create(m *ParticipantsTable) error
	update(m *ParticipantsTable) error
	delete(id string) error
	getByID(id string) (*ParticipantsTable, error)
	getAll() ([]*ParticipantsTable, error)
	getByWalletID(walletId string) (*ParticipantsTable, error)
}

func FactoryStorage(db *sqlx.DB, user *models.User, txID string) ServicesParticipantsTableRepository {
	var s ServicesParticipantsTableRepository
	engine := db.DriverName()
	switch engine {
	case Postgresql:
		return newParticipantsTablePsqlRepository(db, user, txID)
	default:
		logger.Error.Println("el motor de base de datos no está implementado.", engine)
	}
	return s
}
