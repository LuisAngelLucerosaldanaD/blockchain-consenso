package miner_response

import (
	"github.com/jmoiron/sqlx"

	"bjungle-consenso/internal/logger"
	"bjungle-consenso/internal/models"
)

const (
	Postgresql = "postgres"
)

type ServicesMinerResponseRepository interface {
	create(m *MinerResponse) error
	update(m *MinerResponse) error
	delete(id string) error
	getByID(id string) (*MinerResponse, error)
	getAll() ([]*MinerResponse, error)
}

func FactoryStorage(db *sqlx.DB, user *models.User, txID string) ServicesMinerResponseRepository {
	var s ServicesMinerResponseRepository
	engine := db.DriverName()
	switch engine {
	case Postgresql:
		return newMinerResponsePsqlRepository(db, user, txID)
	default:
		logger.Error.Println("el motor de base de datos no está implementado.", engine)
	}
	return s
}
