package cfg

import (
	"bjungle-consenso/internal/models"
	"bjungle-consenso/pkg/cfg/messages"

	"github.com/jmoiron/sqlx"
)

type Server struct {
	SrvMessage messages.PortsServerMessages
}

func NewServerCfg(db *sqlx.DB, user *models.User, txID string) *Server {

	repoMessage := messages.FactoryStorage(db, user, txID)
	srvMessage := messages.NewMessagesService(repoMessage, user, txID)

	return &Server{
		SrvMessage: srvMessage,
	}
}
