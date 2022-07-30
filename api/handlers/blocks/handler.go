package blocks

import (
	"bjungle-consenso/internal/env"
	"bjungle-consenso/internal/grpc/blocks_proto"
	"bjungle-consenso/internal/logger"
	"bjungle-consenso/internal/msg"
	"context"
	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"net/http"
	"strconv"
)

type handlerBlocks struct {
	DB   *sqlx.DB
	TxID string
}

func (h *handlerBlocks) GetAllBlocks(c *fiber.Ctx) error {
	res := resAllBlocks{Error: true}
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
		blocks = append(blocks, &Block{
			Id:                 block.Id,
			Data:               block.Data,
			Nonce:              block.Nonce,
			Difficulty:         block.Difficulty,
			MinedBy:            block.MinedAt,
			MinedAt:            block.MinedAt,
			Timestamp:          block.Timestamp,
			Hash:               block.Hash,
			PrevHash:           block.PrevHash,
			StatusId:           block.StatusId,
			IdUser:             block.IdUser,
			LastValidationDate: block.LastValidationDate,
			CreatedAt:          block.CreatedAt,
			UpdatedAt:          block.UpdatedAt,
		})
	}

	res.Data = blocks
	res.Code, res.Type, res.Msg = msg.GetByCode(29, h.DB, h.TxID)
	res.Error = false
	return c.Status(http.StatusOK).JSON(res)
}
