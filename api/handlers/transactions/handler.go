package transactions

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
)

type handlerTransactions struct {
	DB   *sqlx.DB
	TxID string
}

// TODO implements all methods
func (h *handlerTransactions) createTransaction(c *fiber.Ctx) error {
	return nil
}

func (h *handlerTransactions) getAllTransactions(c *fiber.Ctx) error {
	return nil
}

func (h *handlerTransactions) getTransactionById(c *fiber.Ctx) error {
	return nil
}

func (h *handlerTransactions) getTransactionsByBlockId(c *fiber.Ctx) error {
	return nil
}
