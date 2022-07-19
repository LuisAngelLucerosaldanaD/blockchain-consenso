package lotteries

import (
	"github.com/jmoiron/sqlx"

	"bjungle-consenso/internal/logger"
	"bjungle-consenso/internal/models"
)

const (
	Postgresql = "postgres"
)

type ServicesLotteryRepository interface {
	create(m *Lottery) error
	update(m *Lottery) error
	delete(id string) error
	getByID(id string) (*Lottery, error)
	getAll() ([]*Lottery, error)
	getActive() (*Lottery, error)
	getActiveForMined() (*Lottery, error)
}

func FactoryStorage(db *sqlx.DB, user *models.User, txID string) ServicesLotteryRepository {
	var s ServicesLotteryRepository
	engine := db.DriverName()
	switch engine {
	case Postgresql:
		return newLotteryPsqlRepository(db, user, txID)
	default:
		logger.Error.Println("el motor de base de datos no está implementado.", engine)
	}
	return s
}
