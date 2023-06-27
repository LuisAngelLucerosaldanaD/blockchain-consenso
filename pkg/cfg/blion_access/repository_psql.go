package blion_access

import (
	"bjungle-consenso/internal/helpers"
	"database/sql"
	"fmt"
	"time"

	"bjungle-consenso/internal/models"
	"github.com/jmoiron/sqlx"
)

// psql estructura de conexión a la BD de postgresql
type psql struct {
	DB   *sqlx.DB
	user *models.User
	TxID string
}

func newBLionAccessPsqlRepository(db *sqlx.DB, user *models.User, txID string) *psql {
	return &psql{
		DB:   db,
		user: user,
		TxID: txID,
	}
}

// Create registra en la BD
func (s *psql) create(m *BLionAccess) error {
	date := time.Now()
	m.UpdatedAt = date
	m.CreatedAt = date
	const psqlInsert = `INSERT INTO cfg.blion_access (id ,key, status, created_at, updated_at) VALUES (:id ,:key, :status,:created_at, :updated_at) `
	rs, err := s.DB.NamedExec(psqlInsert, &m)
	if err != nil {
		return err
	}
	if i, _ := rs.RowsAffected(); i == 0 {
		return fmt.Errorf("ecatch:108")
	}
	return nil
}

// Update actualiza un registro en la BD
func (s *psql) update(m *BLionAccess) error {
	date := time.Now()
	m.UpdatedAt = date
	const psqlUpdate = `UPDATE cfg.blion_access SET key = :key, status = :status, updated_at = :updated_at WHERE id = :id `
	rs, err := s.DB.NamedExec(psqlUpdate, &m)
	if err != nil {
		return err
	}
	if i, _ := rs.RowsAffected(); i == 0 {
		return fmt.Errorf("ecatch:108")
	}
	return nil
}

// Delete elimina un registro de la BD
func (s *psql) delete(id string) error {
	const psqlDelete = `DELETE FROM cfg.blion_access WHERE id = :id `
	m := BLionAccess{ID: id}
	rs, err := s.DB.NamedExec(psqlDelete, &m)
	if err != nil {
		return err
	}
	if i, _ := rs.RowsAffected(); i == 0 {
		return fmt.Errorf("ecatch:108")
	}
	return nil
}

// GetByID consulta un registro por su ID
func (s *psql) getByID(id string) (*BLionAccess, error) {
	const psqlGetByID = `SELECT id , key, status, created_at, updated_at FROM cfg.blion_access WHERE id = $1 `
	mdl := BLionAccess{}
	err := s.DB.Get(&mdl, psqlGetByID, id)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return &mdl, err
	}
	return &mdl, nil
}

// GetAll consulta todos los registros de la BD
func (s *psql) getAll() ([]*BLionAccess, error) {
	var ms []*BLionAccess
	const psqlGetAll = ` SELECT id , key, status, created_at, updated_at FROM cfg.blion_access `

	err := s.DB.Select(&ms, psqlGetAll)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return ms, err
	}
	return ms, nil
}

func (s *psql) getByIDs(ids []string) ([]*BLionAccess, error) {
	var ms []*BLionAccess
	const psqlGetAll = ` SELECT id , key, status, created_at, updated_at FROM cfg.blion_access where id in(%s)`

	err := s.DB.Select(&ms, fmt.Sprintf(psqlGetAll, helpers.SliceToString(ids)))
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return ms, err
	}
	return ms, nil
}
