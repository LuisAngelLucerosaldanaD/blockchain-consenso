package transactions

import (
	"bjungle-consenso/internal/env"
	"bjungle-consenso/internal/grpc/transactions_proto"
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

type handlerTransactions struct {
	DB   *sqlx.DB
	TxID string
}

// TODO implements all methods
func (h *handlerTransactions) createTransaction(c *fiber.Ctx) error {
	return nil
}

// getAllTransactions godoc
// @Summary Método para obtener todas las transacciones de la blockchain
// @Description Método para obtener todas las transacciones de la blockchain
// @tags Transacción
// @Produce json
// @Param limit path string true "Número de transacciones por petición"
// @Param offset path string true "Salto de transacciones por petición"
// @Success 200 {object} ResTransactions
// @Router /api/v1/transactions/all/{limit}/{offset} [get]
func (h *handlerTransactions) getAllTransactions(c *fiber.Ctx) error {
	e := env.NewConfiguration()
	res := ResTransactions{Error: true}
	connTrx, err := grpc.Dial(e.TransactionsService.Port, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		logger.Error.Printf("error conectando con el servicio transaction de blockchain: %s", err)
		res.Code, res.Type, res.Msg = msg.GetByCode(22, h.DB, h.TxID)
		return c.Status(http.StatusAccepted).JSON(res)
	}
	defer connTrx.Close()

	limitStr := c.Params("limit")
	limit, err := strconv.ParseInt(limitStr, 10, 64)
	if err != nil {
		logger.Error.Printf("no se pudo obtener el parametro limit")
		res.Code, res.Type, res.Msg = msg.GetByCode(1, h.DB, h.TxID)
		return c.Status(http.StatusAccepted).JSON(res)
	}
	offsetStr := c.Params("offset")
	offset, err := strconv.ParseInt(offsetStr, 10, 64)
	if err != nil {
		logger.Error.Printf("no se pudo obtener el parametro offset")
		res.Code, res.Type, res.Msg = msg.GetByCode(1, h.DB, h.TxID)
		return c.Status(http.StatusAccepted).JSON(res)
	}

	client := transactions_proto.NewTransactionsServicesClient(connTrx)

	resGrpcTrx, err := client.GetAllTransactions(context.Background(), &transactions_proto.GetAllTransactionsRequest{
		Limit:   limit,
		Offset:  offset,
		BlockId: 0,
	})
	if err != nil {
		logger.Error.Printf("no se pudo obtener todas las transacciones, error: %s", err)
		res.Code, res.Type, res.Msg = msg.GetByCode(22, h.DB, h.TxID)
		return c.Status(http.StatusAccepted).JSON(res)
	}

	if resGrpcTrx == nil {
		logger.Error.Printf("no se pudo obtener todas las transacciones")
		res.Code, res.Type, res.Msg = msg.GetByCode(22, h.DB, h.TxID)
		return c.Status(http.StatusAccepted).JSON(res)
	}

	if resGrpcTrx.Error {
		logger.Error.Printf(resGrpcTrx.Msg)
		res.Code, res.Type, res.Msg = msg.GetByCode(int(resGrpcTrx.Code), h.DB, h.TxID)
		return c.Status(http.StatusAccepted).JSON(res)
	}

	var transactions []*Transaction
	for _, trx := range resGrpcTrx.Data {
		transactions = append(transactions, &Transaction{
			Id:        trx.Id,
			From:      trx.From,
			To:        trx.Data,
			Amount:    trx.Amount,
			TypeId:    trx.TypeId,
			Data:      trx.Data,
			Block:     trx.Block,
			Files:     trx.Files,
			CreatedAt: trx.CreatedAt,
			UpdatedAt: trx.UpdatedAt,
		})
	}

	res.Data = transactions
	res.Error = false
	res.Code, res.Type, res.Msg = msg.GetByCode(29, h.DB, h.TxID)
	return c.Status(http.StatusOK).JSON(res)
}

// getTransactionById godoc
// @Summary Método para obtener una transacción de la blockchain
// @Description Método para obtener una transacción de la blockchain por su id y el id del bloque
// @tags Transacción
// @Produce json
// @Param trx path string true "Id de la transacción"
// @Param block path string true "Id del bloque"
// @Success 200 {object} ResTransaction
// @Router /api/v1/transactions/{trx}/{block} [get]
func (h *handlerTransactions) getTransactionById(c *fiber.Ctx) error {
	e := env.NewConfiguration()
	res := ResTransaction{Error: true}
	connTrx, err := grpc.Dial(e.TransactionsService.Port, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		logger.Error.Printf("error conectando con el servicio transaction de blockchain: %s", err)
		res.Code, res.Type, res.Msg = msg.GetByCode(22, h.DB, h.TxID)
		return c.Status(http.StatusAccepted).JSON(res)
	}
	defer connTrx.Close()

	trx := c.Params("trx")
	blockStr := c.Params("block")
	block, err := strconv.ParseInt(blockStr, 10, 64)
	if err != nil {
		logger.Error.Printf("no se pudo obtener el parametro offset")
		res.Code, res.Type, res.Msg = msg.GetByCode(1, h.DB, h.TxID)
		return c.Status(http.StatusAccepted).JSON(res)
	}

	client := transactions_proto.NewTransactionsServicesClient(connTrx)

	resGrpcTrx, err := client.GetTransactionByID(context.Background(), &transactions_proto.GetTransactionByIdRequest{
		Id:      trx,
		BlockId: block,
	})
	if err != nil {
		logger.Error.Printf("no se pudo obtener todas la transaccion, error: %s", err)
		res.Code, res.Type, res.Msg = msg.GetByCode(22, h.DB, h.TxID)
		return c.Status(http.StatusAccepted).JSON(res)
	}

	if resGrpcTrx == nil {
		logger.Error.Printf("no se pudo obtener la transaccion")
		res.Code, res.Type, res.Msg = msg.GetByCode(22, h.DB, h.TxID)
		return c.Status(http.StatusAccepted).JSON(res)
	}

	if resGrpcTrx.Error {
		logger.Error.Printf(resGrpcTrx.Msg)
		res.Code, res.Type, res.Msg = msg.GetByCode(int(resGrpcTrx.Code), h.DB, h.TxID)
		return c.Status(http.StatusAccepted).JSON(res)
	}

	data := resGrpcTrx.Data

	res.Data = &Transaction{
		Id:        data.Id,
		From:      data.From,
		To:        data.To,
		Amount:    data.Amount,
		TypeId:    data.TypeId,
		Data:      data.Data,
		Block:     data.Block,
		Files:     data.Files,
		CreatedAt: data.CreatedAt,
		UpdatedAt: data.UpdatedAt,
	}
	res.Error = false
	res.Code, res.Type, res.Msg = msg.GetByCode(29, h.DB, h.TxID)
	return c.Status(http.StatusOK).JSON(res)
}

// getTransactionsByBlockId godoc
// @Summary Método para obtener todas las transacciones por id del bloque
// @Description Método para obtener todas las transacciones por id del bloque de la blockchain
// @tags Transacción
// @Produce json
// @Param block path string true "Id del bloque"
// @Success 200 {object} ResTransaction
// @Router /api/v1/transactions/all/{block} [get]
func (h *handlerTransactions) getTransactionsByBlockId(c *fiber.Ctx) error {
	e := env.NewConfiguration()
	res := ResTransactions{Error: true}
	connTrx, err := grpc.Dial(e.TransactionsService.Port, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		logger.Error.Printf("error conectando con el servicio transaction de blockchain: %s", err)
		res.Code, res.Type, res.Msg = msg.GetByCode(22, h.DB, h.TxID)
		return c.Status(http.StatusAccepted).JSON(res)
	}
	defer connTrx.Close()

	blockStr := c.Params("block")
	block, err := strconv.ParseInt(blockStr, 10, 64)
	if err != nil {
		logger.Error.Printf("no se pudo obtener el parametro block, error: ", err.Error())
		res.Code, res.Type, res.Msg = msg.GetByCode(1, h.DB, h.TxID)
		return c.Status(http.StatusAccepted).JSON(res)
	}

	client := transactions_proto.NewTransactionsServicesClient(connTrx)

	resGrpcTrx, err := client.GetTransactionsByBlockId(context.Background(), &transactions_proto.RqGetTransactionByBlock{
		BlockId: block,
	})
	if err != nil {
		logger.Error.Printf("no se pudo obtener todas las transacciones, error: %s", err)
		res.Code, res.Type, res.Msg = msg.GetByCode(22, h.DB, h.TxID)
		return c.Status(http.StatusAccepted).JSON(res)
	}

	if resGrpcTrx == nil {
		logger.Error.Printf("no se pudo obtener todas las transacciones")
		res.Code, res.Type, res.Msg = msg.GetByCode(22, h.DB, h.TxID)
		return c.Status(http.StatusAccepted).JSON(res)
	}

	if resGrpcTrx.Error {
		logger.Error.Printf(resGrpcTrx.Msg)
		res.Code, res.Type, res.Msg = msg.GetByCode(int(resGrpcTrx.Code), h.DB, h.TxID)
		return c.Status(http.StatusAccepted).JSON(res)
	}

	var transactions []*Transaction
	for _, trx := range resGrpcTrx.Data {
		transactions = append(transactions, &Transaction{
			Id:        trx.Id,
			From:      trx.From,
			To:        trx.Data,
			Amount:    trx.Amount,
			TypeId:    trx.TypeId,
			Data:      trx.Data,
			Block:     trx.Block,
			Files:     trx.Files,
			CreatedAt: trx.CreatedAt,
			UpdatedAt: trx.UpdatedAt,
		})
	}

	res.Data = transactions
	res.Error = false
	res.Code, res.Type, res.Msg = msg.GetByCode(29, h.DB, h.TxID)
	return c.Status(http.StatusOK).JSON(res)
}
