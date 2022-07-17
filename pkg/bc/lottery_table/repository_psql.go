package lottery_table

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

func newLotteryTablePsqlRepository(db *sqlx.DB, user *models.User, txID string) *psql {
	return &psql{
		DB:   db,
		user: user,
		TxID: txID,
	}
}

// Create registra en la BD
func (s *psql) create(m *LotteryTable) error {
	date := time.Now()
	m.UpdatedAt = date
	m.CreatedAt = date
	const psqlInsert = `INSERT INTO bc.lottery_table (id ,block_id, registration_start_date, registration_end_date, lottery_start_date, lottery_end_date, process_end_date, process_status, created_at, updated_at) VALUES (:id ,:block_id, :registration_start_date, :registration_end_date, :lottery_start_date, :lottery_end_date, :process_end_date, :process_status,:created_at, :updated_at) `
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
func (s *psql) update(m *LotteryTable) error {
	date := time.Now()
	m.UpdatedAt = date
	const psqlUpdate = `UPDATE bc.lottery_table SET block_id = :block_id, registration_start_date = :registration_start_date, registration_end_date = :registration_end_date, lottery_start_date = :lottery_start_date, lottery_end_date = :lottery_end_date, process_end_date = :process_end_date, process_status = :process_status, updated_at = :updated_at WHERE id = :id `
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
	const psqlDelete = `DELETE FROM bc.lottery_table WHERE id = :id `
	m := LotteryTable{ID: id}
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
func (s *psql) getByID(id string) (*LotteryTable, error) {
	const psqlGetByID = `SELECT id , block_id, registration_start_date, registration_end_date, lottery_start_date, lottery_end_date, process_end_date, process_status, created_at, updated_at FROM bc.lottery_table WHERE id = $1 `
	mdl := LotteryTable{}
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
func (s *psql) getAll() ([]*LotteryTable, error) {
	var ms []*LotteryTable
	const psqlGetAll = ` SELECT id , block_id, registration_start_date, registration_end_date, lottery_start_date, lottery_end_date, process_end_date, process_status, created_at, updated_at FROM bc.lottery_table `

	err := s.DB.Select(&ms, psqlGetAll)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return ms, err
	}
	return ms, nil
}

func (s *psql) getActive() (*LotteryTable, error) {
	const psqlGetByID = `SELECT id , block_id, registration_start_date, registration_end_date, lottery_start_date, lottery_end_date, process_end_date, process_status, created_at, updated_at FROM bc.lottery_table WHERE registration_end_date = NULL AND registration_start_date <> NULL AND process_status = 25 limit 1`
	mdl := LotteryTable{}
	err := s.DB.Get(&mdl, psqlGetByID)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("no hay sorteos activos actualmente")
		}
		return &mdl, err
	}
	return &mdl, nil
}
