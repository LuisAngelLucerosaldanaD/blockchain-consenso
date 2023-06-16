package blocks

import (
	"bjungle-consenso/internal/env"
	"bjungle-consenso/internal/grpc/blocks_proto"
	"bjungle-consenso/internal/logger"
	"bjungle-consenso/internal/msg"
	"bjungle-consenso/pkg/bc"
	"context"
	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"net/http"
	"strconv"
	"time"
)

type handlerBlocks struct {
	DB   *sqlx.DB
	TxID string
}

// GetAllBlocks godoc
// @Summary Método para obtener todos los bloques de la Blockchain
// @Description Método para obtener todos los bloques de la Blockchain de BLion
// @tags Block
// @Accept json
// @Produce json
// @Param limit path string true "Número de bloques por petición"
// @Param offset path string true "Salto de bloques por petición"
// @Success 200 {object} ResAllBlocks
// @Router /api/v1/block/get-all/{limit}/{offset} [get]
func (h *handlerBlocks) GetAllBlocks(c *fiber.Ctx) error {
	res := ResAllBlocks{Error: true}
	e := env.NewConfiguration()

	limitStr := c.Params("limit")
	limit, _ := strconv.ParseInt(limitStr, 10, 16)
	offsetStr := c.Params("offset")
	offset, _ := strconv.ParseInt(offsetStr, 10, 16)
	connBk, err := grpc.Dial(e.BlockService.Port, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		logger.Error.Printf("error conectando con el servicio auth de blockchain: %s", err)
		res.Code, res.Type, res.Msg = msg.GetByCode(22, h.DB, h.TxID)
		return c.Status(http.StatusAccepted).JSON(res)
	}
	defer connBk.Close()

	clientBlock := blocks_proto.NewBlockServicesBlocksClient(connBk)

	resBlocks, err := clientBlock.GetBlock(context.Background(), &blocks_proto.GetAllBlockRequest{Limit: limit, Offset: offset})
	if err != nil {
		logger.Error.Printf("No se pudo conectar con el servicio block engine, Error: %v", err)
		res.Code, res.Type, res.Msg = msg.GetByCode(22, h.DB, h.TxID)
		return c.Status(http.StatusAccepted).JSON(res)
	}
	if resBlocks == nil {
		logger.Error.Printf("No se pudo conectar con el servicio block engine")
		res.Code, res.Type, res.Msg = msg.GetByCode(22, h.DB, h.TxID)
		return c.Status(http.StatusAccepted).JSON(res)
	}
	if resBlocks.Error {
		logger.Error.Printf(resBlocks.Msg)
		res.Code, res.Type, res.Msg = int(resBlocks.Code), int(resBlocks.Type), resBlocks.Msg
		return c.Status(http.StatusAccepted).JSON(res)
	}

	var blocks []*Block
	for _, block := range resBlocks.Data {

		layout := "2006-01-02 15:04:05.999999999 -0700 MST"

		timestamp, _ := time.Parse(layout, block.Timestamp)
		createdAt, _ := time.Parse(layout, block.CreatedAt)
		updatedAt, _ := time.Parse(layout, block.UpdatedAt)
		minedAt, _ := time.Parse(layout, block.MinedAt)
		lastValidationDate, _ := time.Parse(layout, block.LastValidationDate)

		blocks = append(blocks, &Block{
			Id:                 block.Id,
			Data:               block.Data,
			Nonce:              block.Nonce,
			Difficulty:         block.Difficulty,
			MinedBy:            block.MinedBy,
			MinedAt:            minedAt,
			Timestamp:          timestamp,
			Hash:               block.Hash,
			PrevHash:           block.PrevHash,
			StatusId:           block.StatusId,
			IdUser:             block.IdUser,
			LastValidationDate: lastValidationDate,
			CreatedAt:          createdAt,
			UpdatedAt:          updatedAt,
		})
	}

	res.Data = blocks
	res.Code, res.Type, res.Msg = msg.GetByCode(29, h.DB, h.TxID)
	res.Error = false
	return c.Status(http.StatusOK).JSON(res)
}

// GetCurrentLottery godoc
// @Summary Método para obtener la lotería actual
// @Description Método para obtener la lotería actual
// @tags Block
// @Accept json
// @Produce json
// @Success 200 {object} resCurrentLottery
// @Router /api/v1/block/current-lottery [get]
func (h *handlerBlocks) GetCurrentLottery(c *fiber.Ctx) error {

	res := resCurrentLottery{Error: true}
	srv := bc.NewServerBk(h.DB, nil, h.TxID)
	lotteryActive, code, err := srv.SrvLottery.GetLotteryActive()
	if err != nil {
		logger.Error.Printf("error al traer la loteria activa: %s", err.Error())
		res.Code, res.Type, res.Msg = msg.GetByCode(code, h.DB, h.TxID)
		return c.Status(http.StatusAccepted).JSON(res)
	}

	if lotteryActive != nil {
		res.Error = false
		res.Data = Lottery(*lotteryActive)
		res.Code, res.Type, res.Msg = msg.GetByCode(29, h.DB, h.TxID)
		return c.Status(http.StatusAccepted).JSON(res)
	}

	lotteryMined, code, err := srv.SrvLottery.GetLotteryActiveForMined()
	if err != nil {
		logger.Error.Printf("error al traer la loteria activa: %s", err.Error())
		res.Code, res.Type, res.Msg = msg.GetByCode(code, h.DB, h.TxID)
		return c.Status(http.StatusAccepted).JSON(res)
	}

	if lotteryMined == nil {
		logger.Error.Printf("No hay una loteria lista para mineria")
		res.Code, res.Type, res.Msg = msg.GetByCode(22, h.DB, h.TxID)
		return c.Status(http.StatusAccepted).JSON(res)
	}

	res.Error = false
	res.Data = Lottery(*lotteryMined)
	res.Code, res.Type, res.Msg = msg.GetByCode(29, h.DB, h.TxID)
	return c.Status(http.StatusOK).JSON(res)
}
