package blion_access

import (
	"github.com/jmoiron/sqlx"

	"bjungle-consenso/internal/logger"
	"bjungle-consenso/internal/models"
)

const (
	Postgresql = "postgres"
)

type ServicesBLionAccessRepository interface {
	create(m *BLionAccess) error
	update(m *BLionAccess) error
	delete(id string) error
	getByID(id string) (*BLionAccess, error)
	getAll() ([]*BLionAccess, error)
	getByIDs(ids []string) ([]*BLionAccess, error)
}

func FactoryStorage(db *sqlx.DB, user *models.User, txID string) ServicesBLionAccessRepository {
	var s ServicesBLionAccessRepository
	engine := db.DriverName()
	switch engine {
	case Postgresql:
		return newBLionAccessPsqlRepository(db, user, txID)
	default:
		logger.Error.Println("el motor de base de datos no está implementado.", engine)
	}
	return s
}
