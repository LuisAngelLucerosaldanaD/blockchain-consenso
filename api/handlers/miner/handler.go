package miner

import (
	"bjungle-consenso/internal/env"
	"bjungle-consenso/internal/grpc/mine_proto"
	"bjungle-consenso/internal/logger"
	"bjungle-consenso/internal/msg"
	"bjungle-consenso/pkg/bc"
	"context"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	grpcMetadata "google.golang.org/grpc/metadata"
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

func (h *handlerMiner) GetBlockToMine(c *fiber.Ctx) error {
	res := responseGetBlock{Error: true}
	e := env.NewConfiguration()
	connBk, err := grpc.Dial(e.BlockService.Port, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		logger.Error.Printf("error conectando con el servicio auth de blockchain: %s", err)
		res.Code, res.Type, res.Msg = msg.GetByCode(22, h.DB, h.TxID)
		return c.Status(http.StatusAccepted).JSON(res)
	}
	defer connBk.Close()

	clientMine := mine_proto.NewMineBlockServicesBlocksClient(connBk)
	token := c.Get("Authorization")[7:]
	ctx := grpcMetadata.AppendToOutgoingContext(context.Background(), "authorization", token)

	resBkMine, err := clientMine.GetBlockToMine(ctx, &mine_proto.GetBlockToMineRequest{})
	if err != nil {
		logger.Error.Printf("No se pudo obtener el bloque a minar, error: %s", err)
		res.Code, res.Type, res.Msg = msg.GetByCode(22, h.DB, h.TxID)
		return c.Status(http.StatusAccepted).JSON(res)
	}

	if resBkMine == nil {
		logger.Error.Printf("No se pudo obtener el bloque a minar")
		res.Code, res.Type, res.Msg = msg.GetByCode(22, h.DB, h.TxID)
		return c.Status(http.StatusAccepted).JSON(res)
	}

	if resBkMine.Error {
		logger.Error.Printf(resBkMine.Msg)
		res.Code, res.Type, res.Msg = msg.GetByCode(int(resBkMine.Code), h.DB, h.TxID)
		return c.Status(http.StatusAccepted).JSON(res)
	}

	block := resBkMine.Data
	res.Data = &DataBlockToMine{
		ID:         block.Id,
		PrevHash:   block.PrevHash,
		Difficulty: e.App.Difficulty,
		Cuota:      float64(e.App.MinimumFee),
	}
	res.Error = false
	res.Code, res.Type, res.Msg = msg.GetByCode(29, h.DB, h.TxID)
	return c.Status(http.StatusOK).JSON(res)
}
