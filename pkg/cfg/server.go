package cfg

import (
	"bjungle-consenso/internal/models"
	"bjungle-consenso/pkg/cfg/blion_access"
	"bjungle-consenso/pkg/cfg/messages"

	"github.com/jmoiron/sqlx"
)

type Server struct {
	SrvMessage     messages.PortsServerMessages
	SrvBLionAccess blion_access.PortsServerBLionAccess
}

func NewServerCfg(db *sqlx.DB, user *models.User, txID string) *Server {

	repoMessage := messages.FactoryStorage(db, user, txID)
	srvMessage := messages.NewMessagesService(repoMessage, user, txID)

	repoBLionAccess := blion_access.FactoryStorage(db, user, txID)
	srvBLionAccess := blion_access.NewBLionAccessService(repoBLionAccess, user, txID)

	return &Server{
		SrvMessage:     srvMessage,
		SrvBLionAccess: srvBLionAccess,
	}
}
