package miner_response

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

func newMinerResponsePsqlRepository(db *sqlx.DB, user *models.User, txID string) *psql {
	return &psql{
		DB:   db,
		user: user,
		TxID: txID,
	}
}

// Create registra en la BD
func (s *psql) create(m *MinerResponse) error {
	date := time.Now()
	m.UpdatedAt = date
	m.CreatedAt = date
	const psqlInsert = `INSERT INTO bc.miner_response (id ,lottery_id, participants_id, hash, status, nonce, difficulty, created_at, updated_at) VALUES (:id ,:lottery_id, :participants_id, :hash, :status, :nonce, :difficulty,:created_at, :updated_at) `
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
func (s *psql) update(m *MinerResponse) error {
	date := time.Now()
	m.UpdatedAt = date
	const psqlUpdate = `UPDATE bc.miner_response SET lottery_id = :lottery_id, participants_id = :participants_id, hash = :hash, status = :status, nonce = :nonce, difficulty = :difficulty, updated_at = :updated_at WHERE id = :id `
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
	const psqlDelete = `DELETE FROM bc.miner_response WHERE id = :id `
	m := MinerResponse{ID: id}
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
func (s *psql) getByID(id string) (*MinerResponse, error) {
	const psqlGetByID = `SELECT id , lottery_id, participants_id, hash, status, nonce, difficulty, created_at, updated_at FROM bc.miner_response WHERE id = $1 `
	mdl := MinerResponse{}
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
func (s *psql) getAll() ([]*MinerResponse, error) {
	var ms []*MinerResponse
	const psqlGetAll = ` SELECT id , lottery_id, participants_id, hash, status, nonce, difficulty, created_at, updated_at FROM bc.miner_response `

	err := s.DB.Select(&ms, psqlGetAll)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return ms, err
	}
	return ms, nil
}

// GetByID consulta un registro por su ID
func (s *psql) getByLotteryID(id string) (*MinerResponse, error) {
	const psqlGetByID = `SELECT id , lottery_id, participants_id, hash, status, nonce, difficulty, created_at, updated_at FROM bc.miner_response WHERE lottery_id = $1 limit 1`
	mdl := MinerResponse{}
	err := s.DB.Get(&mdl, psqlGetByID, id)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return &mdl, err
	}
	return &mdl, nil
}

// GetByID consulta un registro por su ID
func (s *psql) getRegister(lotteryId string) (*MinerResponse, error) {
	const psqlRegister = `SELECT id , lottery_id, participants_id, hash, status, nonce, difficulty, created_at, updated_at FROM bc.miner_response WHERE status = 29 and lottery_id = $1 limit 1`
	mdl := MinerResponse{}
	err := s.DB.Get(&mdl, psqlRegister, lotteryId)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return &mdl, err
	}
	return &mdl, nil
}
