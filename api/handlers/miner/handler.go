package miner

import (
	"bjungle-consenso/internal/env"
	"bjungle-consenso/internal/logger"
	"bjungle-consenso/internal/msg"
	"bjungle-consenso/pkg/bc"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"net/http"
)

type handlerMiner struct {
	DB   *sqlx.DB
	TxID string
}

func (h *handlerMiner) RegisterHashMined(c *fiber.Ctx) error {
	res := responseRegisterMined{Error: true}
	request := rqRegisterMined{}
	e := env.NewConfiguration()
	err := c.BodyParser(&request)
	if err != nil {
		logger.Error.Printf("couldn't bind model rqRegisterMined: %v", err)
		res.Code, res.Type, res.Msg = msg.GetByCode(1, h.DB, h.TxID)
		return c.Status(http.StatusAccepted).JSON(res)
	}

	srvBc := bc.NewServerBk(h.DB, nil, h.TxID)

	lottery, code, err := srvBc.SrvLottery.GetLotteryActiveForMined()
	if err != nil {
		logger.Error.Printf("couldn't get lottery for mined, error: %v", err)
		res.Code, res.Type, res.Msg = msg.GetByCode(code, h.DB, h.TxID)
		return c.Status(http.StatusAccepted).JSON(res)
	}

	hash, code, err := srvBc.SrvMinerResponse.GetMinerResponseRegister(lottery.ID)
	if err != nil {
		logger.Error.Printf("couldn't get hash register: %v", err)
		res.Code, res.Type, res.Msg = msg.GetByCode(code, h.DB, h.TxID)
		return c.Status(http.StatusAccepted).JSON(res)
	}

	if hash != nil {
		logger.Error.Printf("El hash para este bloque ya ha sido registrado")
		res.Code, res.Type, res.Msg = 22, 1, "El hash para este bloque ya ha sido registrado"
		return c.Status(http.StatusAccepted).JSON(res)
	}

	participant, code, err := srvBc.SrvParticipants.GetParticipantsByWalletIDAndLotteryID(request.WalletID, lottery.ID)
	if err != nil {
		logger.Error.Printf("couldn't get participant by wallet id, error: %v", err)
		res.Code, res.Type, res.Msg = msg.GetByCode(code, h.DB, h.TxID)
		return c.Status(http.StatusAccepted).JSON(res)
	}

	if !participant.Accepted {
		logger.Error.Printf("Usted no ha sido selecionado para minar")
		res.Code, res.Type, res.Msg = msg.GetByCode(22, h.DB, h.TxID)
		return c.Status(http.StatusAccepted).JSON(res)
	}

	if participant.TypeCharge != 23 {
		logger.Error.Printf("solo los mineros pueden registrar el hash del bloque minado como respuesta a la loteria")
		res.Code, res.Type, res.Msg = msg.GetByCode(22, h.DB, h.TxID)
		return c.Status(http.StatusAccepted).JSON(res)
	}

	_, code, err = srvBc.SrvMinerResponse.CreateMinerResponse(uuid.New().String(), lottery.ID, participant.ID, request.Hash, 29, request.Nonce, e.App.Difficulty)
	if err != nil {
		logger.Error.Printf("couldn't register hash, error: %v", err)
		res.Code, res.Type, res.Msg = msg.GetByCode(code, h.DB, h.TxID)
		return c.Status(http.StatusAccepted).JSON(res)
	}

	res.Data = "hash registrado correctamente"
	res.Code, res.Type, res.Msg = msg.GetByCode(29, h.DB, h.TxID)
	res.Error = false
	return c.Status(http.StatusOK).JSON(res)
}
