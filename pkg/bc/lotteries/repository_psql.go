package lotteries

import (
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

func newLotteryPsqlRepository(db *sqlx.DB, user *models.User, txID string) *psql {
	return &psql{
		DB:   db,
		user: user,
		TxID: txID,
	}
}

// Create registra en la BD
func (s *psql) create(m *Lottery) error {
	date := time.Now()
	m.UpdatedAt = date
	m.CreatedAt = date
	const psqlInsert = `INSERT INTO bc.lotteries (id ,block_id, registration_start_date, registration_end_date, lottery_start_date, lottery_end_date, process_end_date, process_status, created_at, updated_at) VALUES (:id ,:block_id, :registration_start_date, :registration_end_date, :lottery_start_date, :lottery_end_date, :process_end_date, :process_status,:created_at, :updated_at) `
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
func (s *psql) update(m *Lottery) error {
	date := time.Now()
	m.UpdatedAt = date
	const psqlUpdate = `UPDATE bc.lotteries SET block_id = :block_id, registration_start_date = :registration_start_date, registration_end_date = :registration_end_date, lottery_start_date = :lottery_start_date, lottery_end_date = :lottery_end_date, process_end_date = :process_end_date, process_status = :process_status, updated_at = :updated_at WHERE id = :id `
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
	const psqlDelete = `DELETE FROM bc.lotteries WHERE id = :id `
	m := Lottery{ID: id}
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
func (s *psql) getByID(id string) (*Lottery, error) {
	const psqlGetByID = `SELECT id , block_id, registration_start_date, registration_end_date, lottery_start_date, lottery_end_date, process_end_date, process_status, created_at, updated_at FROM bc.lotteries WHERE id = $1 `
	mdl := Lottery{}
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
func (s *psql) getAll() ([]*Lottery, error) {
	var ms []*Lottery
	const psqlGetAll = ` SELECT id , block_id, registration_start_date, registration_end_date, lottery_start_date, lottery_end_date, process_end_date, process_status, created_at, updated_at FROM bc.lotteries `

	err := s.DB.Select(&ms, psqlGetAll)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return ms, err
	}
	return ms, nil
}

func (s *psql) getActive() (*Lottery, error) {
	const psqlGetByID = `SELECT id , block_id, registration_start_date, registration_end_date, lottery_start_date, lottery_end_date, process_end_date, process_status, created_at, updated_at FROM bc.lotteries WHERE process_status = 25 limit 1`
	mdl := Lottery{}
	err := s.DB.Get(&mdl, psqlGetByID)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return &mdl, err
	}
	return &mdl, nil
}

func (s *psql) getActiveForMined() (*Lottery, error) {
	const psqlGetByID = `SELECT id , block_id, registration_start_date, registration_end_date, lottery_start_date, lottery_end_date, process_end_date, process_status, created_at, updated_at FROM bc.lotteries WHERE process_status = 27 limit 1`
	mdl := Lottery{}
	err := s.DB.Get(&mdl, psqlGetByID)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return &mdl, err
	}
	return &mdl, nil
}
