package reward_table

import (
	"database/sql"
	"fmt"

	"bjungle-consenso/internal/models"
	"github.com/jmoiron/sqlx"
)

// sqlServer estructura de conexión a la BD de mssql
type sqlserver struct {
	DB   *sqlx.DB
	user *models.User
	TxID string
}

func newRewardTableSqlServerRepository(db *sqlx.DB, user *models.User, txID string) *sqlserver {
	return &sqlserver{
		DB:   db,
		user: user,
		TxID: txID,
	}
}

// Create registra en la BD
func (s *sqlserver) create(m *RewardTable) error {
	date := time.Now()
	m.UpdatedAt = date
	m.CreatedAt = date
	const sqlInsert = `INSERT INTO bk.reward_table (id ,lottery_id, id_wallet, amount, block_id, created_at, updated_at) VALUES (:id ,:lottery_id, :id_wallet, :amount, :block_id:created_at, :updated_at) `
	rs, err := s.DB.NamedExec(sqlInsert, &m)
	if err != nil {
		return err
	}
	if i, _ := rs.RowsAffected(); i == 0 {
		return fmt.Errorf("ecatch:108")
	}
	return nil
}

// Update actualiza un registro en la BD
func (s *sqlserver) update(m *RewardTable) error {
	date := time.Now()
	m.UpdatedAt = date
	const sqlUpdate = `UPDATE bk.reward_table SET lottery_id = :lottery_id, id_wallet = :id_wallet, amount = :amount, block_id = :block_id, updated_at = :updated_at WHERE id = :id `
	rs, err := s.DB.NamedExec(sqlUpdate, &m)
	if err != nil {
		return err
	}
	if i, _ := rs.RowsAffected(); i == 0 {
		return fmt.Errorf("ecatch:108")
	}
	return nil
}

// Delete elimina un registro de la BD
func (s *sqlserver) delete(id string) error {
	const sqlDelete = `DELETE FROM bk.reward_table WHERE id = :id `
	m := RewardTable{ID: id}
	rs, err := s.DB.NamedExec(sqlDelete, &m)
	if err != nil {
		return err
	}
	if i, _ := rs.RowsAffected(); i == 0 {
		return fmt.Errorf("ecatch:108")
	}
	return nil
}

// GetByID consulta un registro por su ID
func (s *sqlserver) getByID(id string) (*RewardTable, error) {
	const sqlGetByID = `SELECT convert(nvarchar(50), id) id , lottery_id, id_wallet, amount, block_id, created_at, updated_at FROM bk.reward_table  WITH (NOLOCK)  WHERE id = @id `
	mdl := RewardTable{}
	err := s.DB.Get(&mdl, sqlGetByID, sql.Named("id", id))
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return &mdl, err
	}
	return &mdl, nil
}

// GetAll consulta todos los registros de la BD
func (s *sqlserver) getAll() ([]*RewardTable, error) {
	var ms []*RewardTable
	const sqlGetAll = `SELECT convert(nvarchar(50), id) id , lottery_id, id_wallet, amount, block_id, created_at, updated_at FROM bk.reward_table  WITH (NOLOCK) `

	err := s.DB.Select(&ms, sqlGetAll)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return ms, err
	}
	return ms, nil
}
