package validation_table

import (
	"database/sql"
	"fmt"

	"bjungle-consenso/internal/models"
	"github.com/jmoiron/sqlx"
)

// psql estructura de conexión a la BD de postgresql
type psql struct {
	DB   *sqlx.DB
	user *models.User
	TxID string
}

func newValidationTablePsqlRepository(db *sqlx.DB, user *models.User, txID string) *psql {
	return &psql{
		DB:   db,
		user: user,
		TxID: txID,
	}
}

// Create registra en la BD
func (s *psql) create(m *ValidationTable) error {
	date := time.Now()
	m.UpdatedAt = date
	m.CreatedAt = date
	const psqlInsert = `INSERT INTO bk.validation_table (id ,lottery_id, wallet_id, participants_id, vote, created_at, updated_at) VALUES (:id ,:lottery_id, :wallet_id, :participants_id, :vote,:created_at, :updated_at) `
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
func (s *psql) update(m *ValidationTable) error {
	date := time.Now()
	m.UpdatedAt = date
	const psqlUpdate = `UPDATE bk.validation_table SET lottery_id = :lottery_id, wallet_id = :wallet_id, participants_id = :participants_id, vote = :vote, updated_at = :updated_at WHERE id = :id `
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
	const psqlDelete = `DELETE FROM bk.validation_table WHERE id = :id `
	m := ValidationTable{ID: id}
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
func (s *psql) getByID(id string) (*ValidationTable, error) {
	const psqlGetByID = `SELECT id , lottery_id, wallet_id, participants_id, vote, created_at, updated_at FROM bk.validation_table WHERE id = $1 `
	mdl := ValidationTable{}
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
func (s *psql) getAll() ([]*ValidationTable, error) {
	var ms []*ValidationTable
	const psqlGetAll = ` SELECT id , lottery_id, wallet_id, participants_id, vote, created_at, updated_at FROM bk.validation_table `

	err := s.DB.Select(&ms, psqlGetAll)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return ms, err
	}
	return ms, nil
}
