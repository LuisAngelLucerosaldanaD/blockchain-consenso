package msg

import (
	"bjungle-consenso/internal/env"
	"bjungle-consenso/internal/logger"
	"bjungle-consenso/pkg/cfg"
	"strconv"

	"github.com/jmoiron/sqlx"
)

func GetByCode(code int, db *sqlx.DB, txID string) (int, int, string) {
	codRes := 0
	msg := ""
	c := env.NewConfiguration()
	srvCFG := cfg.NewServerCfg(db, nil, txID)
	m, codErr, err := srvCFG.SrvMessage.GetMessagesByID(code)
	if err != nil {
		return codRes, 0, strconv.Itoa(codErr)
	}

	switch c.App.Language {
	case "sp":
		msg = m.Spa
	case "en":
		msg = m.Eng
	default:
		logger.Error.Println("el sistema no tiene implementado el idioma: ", c.App.Language)
	}
	codRes = m.ID
	return codRes, m.TypeMessage, msg
}
