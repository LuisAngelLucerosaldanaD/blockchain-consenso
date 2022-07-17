package validation_table

import (
	"github.com/jmoiron/sqlx"

	"bjungle-consenso/internal/logger"
	"bjungle-consenso/internal/models"
)

const (
	Postgresql = "postgres"
)

type ServicesValidationTableRepository interface {
	create(m *ValidationTable) error
	update(m *ValidationTable) error
	delete(id string) error
	getByID(id string) (*ValidationTable, error)
	getAll() ([]*ValidationTable, error)
}

func FactoryStorage(db *sqlx.DB, user *models.User, txID string) ServicesValidationTableRepository {
	var s ServicesValidationTableRepository
	engine := db.DriverName()
	switch engine {
	case Postgresql:
		return newValidationTablePsqlRepository(db, user, txID)
	default:
		logger.Error.Println("el motor de base de datos no está implementado.", engine)
	}
	return s
}
