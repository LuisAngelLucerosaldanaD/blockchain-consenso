package reward_table

import (
	"database/sql"
	"fmt"
	"time"

	"bjungle-consenso/internal/models"
	"github.com/jmoiron/sqlx"
)

// Orcl estructura de conexión a la BD de Oracle
type orcl struct {
	DB   *sqlx.DB
	user *models.User
	TxID string
}

func newRewardTableOrclRepository(db *sqlx.DB, user *models.User, txID string) *orcl {
	return &orcl{
		DB:   db,
		user: user,
		TxID: txID,
	}
}

// Create registra en la BD
func (s *orcl) create(m *RewardTable) error {
	date := time.Now()
	m.UpdatedAt = date
	m.CreatedAt = date
	const osqlInsert = `INSERT INTO bk.reward_table (id ,lottery_id, id_wallet, amount, block_id, created_at, updated_at)  VALUES (:id ,:lottery_id, :id_wallet, :amount, :block_id,:created_at, :updated_at) `
	rs, err := s.DB.NamedExec(osqlInsert, &m)
	if err != nil {
		return err
	}
	if i, _ := rs.RowsAffected(); i == 0 {
		return fmt.Errorf("ecatch:108")
	}
	return nil
}

// Update actualiza un registro en la BD
func (s *orcl) update(m *RewardTable) error {
	date := time.Now()
	m.UpdatedAt = date
	const osqlUpdate = `UPDATE bk.reward_table SET lottery_id = :lottery_id, id_wallet = :id_wallet, amount = :amount, block_id = :block_id, updated_at = :updated_at WHERE id = :id  `
	rs, err := s.DB.NamedExec(osqlUpdate, &m)
	if err != nil {
		return err
	}
	if i, _ := rs.RowsAffected(); i == 0 {
		return fmt.Errorf("ecatch:108")
	}
	return nil
}

// Delete elimina un registro de la BD
func (s *orcl) delete(id string) error {
	const osqlDelete = `DELETE FROM bk.reward_table WHERE id = :id `
	m := RewardTable{ID: id}
	rs, err := s.DB.NamedExec(osqlDelete, &m)
	if err != nil {
		return err
	}
	if i, _ := rs.RowsAffected(); i == 0 {
		return fmt.Errorf("ecatch:108")
	}
	return nil
}

// GetByID consulta un registro por su ID
func (s *orcl) getByID(id string) (*RewardTable, error) {
	const osqlGetByID = `SELECT id , lottery_id, id_wallet, amount, block_id, created_at, updated_at FROM bk.reward_table WHERE id = :1 `
	mdl := RewardTable{}
	err := s.DB.Get(&mdl, osqlGetByID, id)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return &mdl, err
	}
	return &mdl, nil
}

// GetAll consulta todos los registros de la BD
func (s *orcl) getAll() ([]*RewardTable, error) {
	var ms []*RewardTable
	const osqlGetAll = ` SELECT id , lottery_id, id_wallet, amount, block_id, created_at, updated_at FROM bk.reward_table `

	err := s.DB.Select(&ms, osqlGetAll)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return ms, err
	}
	return ms, nil
}
