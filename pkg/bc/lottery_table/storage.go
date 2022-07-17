package lottery_table

import (
	"github.com/jmoiron/sqlx"

	"bjungle-consenso/internal/logger"
	"bjungle-consenso/internal/models"
)

const (
	Postgresql = "postgres"
)

type ServicesLotteryTableRepository interface {
	create(m *LotteryTable) error
	update(m *LotteryTable) error
	delete(id string) error
	getByID(id string) (*LotteryTable, error)
	getAll() ([]*LotteryTable, error)
	getActive() (*LotteryTable, error)
}

func FactoryStorage(db *sqlx.DB, user *models.User, txID string) ServicesLotteryTableRepository {
	var s ServicesLotteryTableRepository
	engine := db.DriverName()
	switch engine {
	case Postgresql:
		return newLotteryTablePsqlRepository(db, user, txID)
	default:
		logger.Error.Println("el motor de base de datos no está implementado.", engine)
	}
	return s
}
