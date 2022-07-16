package auth

import (
	"bjungle-consenso/internal/models"
	"bjungle-consenso/pkg/auth/frozen_money"
	"github.com/jmoiron/sqlx"
)

type Server struct {
	SrvFrozenMoney frozen_money.PortsServerFrozenMoney
}

func NewServerAuth(db *sqlx.DB, user *models.User, txID string) *Server {

	repoFrozen := frozen_money.FactoryStorage(db, user, txID)
	srvFrozen := frozen_money.NewFrozenMoneyService(repoFrozen, user, txID)

	return &Server{
		SrvFrozenMoney: srvFrozen,
	}
}
