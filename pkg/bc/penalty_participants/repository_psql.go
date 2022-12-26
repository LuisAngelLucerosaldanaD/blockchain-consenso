package penalty_participants

import (
	"bjungle-consenso/internal/models"
	"database/sql"
	"fmt"
	"github.com/jmoiron/sqlx"
	"time"
)

// psql estructura de conexión a la BD de postgresql
type psql struct {
	DB   *sqlx.DB
	user *models.User
	TxID string
}

func newPenaltyParticipantsPsqlRepository(db *sqlx.DB, user *models.User, txID string) *psql {
	return &psql{
		DB:   db,
		user: user,
		TxID: txID,
	}
}

// Create registra en la BD
func (s *psql) create(m *PenaltyParticipants) error {
	date := time.Now()
	m.UpdatedAt = date
	m.CreatedAt = date
	const psqlInsert = `INSERT INTO bc.penalty_participants (id ,lottery_id, participants_id, amount, penalty_percentage, created_at, updated_at) VALUES (:id ,:lottery_id, :participants_id, :amount, :penalty_percentage,:created_at, :updated_at) `
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
func (s *psql) update(m *PenaltyParticipants) error {
	date := time.Now()
	m.UpdatedAt = date
	const psqlUpdate = `UPDATE bc.penalty_participants SET lottery_id = :lottery_id, participants_id = :participants_id, amount = :amount, penalty_percentage = :penalty_percentage, updated_at = :updated_at WHERE id = :id `
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
	const psqlDelete = `DELETE FROM bc.penalty_participants WHERE id = :id `
	m := PenaltyParticipants{ID: id}
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
func (s *psql) getByID(id string) (*PenaltyParticipants, error) {
	const psqlGetByID = `SELECT id , lottery_id, participants_id, amount, penalty_percentage, created_at, updated_at FROM bc.penalty_participants WHERE id = $1 `
	mdl := PenaltyParticipants{}
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
func (s *psql) getByWalletID(walletId string) ([]*PenaltyParticipants, error) {
	const psqlGetByWalletID = `SELECT pp.id , pp.lottery_id, pp.participants_id, pp.amount, pp.penalty_percentage, pp.created_at, pp.updated_at FROM bc.penalty_participants pp
join bc.participants p on(p.id = pp.participants_id) WHERE p.wallet_id = '%s'`
	var ms []*PenaltyParticipants
	err := s.DB.Select(&ms, fmt.Sprintf(psqlGetByWalletID, walletId))
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return ms, err
	}
	return ms, nil
}

// GetAll consulta todos los registros de la BD
func (s *psql) getAll() ([]*PenaltyParticipants, error) {
	var ms []*PenaltyParticipants
	const psqlGetAll = ` SELECT id , lottery_id, participants_id, amount, penalty_percentage, created_at, updated_at FROM bc.penalty_participants `

	err := s.DB.Select(&ms, psqlGetAll)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return ms, err
	}
	return ms, nil
}
