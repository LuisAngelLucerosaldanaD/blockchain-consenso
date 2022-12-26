package transactions

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
)

func RouterParticipants(app *fiber.App, db *sqlx.DB, txID string) {
	h := handlerTransactions{DB: db, TxID: txID}
	api := app.Group("/api")
	v1 := api.Group("/v1")
	trxRouter := v1.Group("/transactions")
	trxRouter.Post("/create", h.createTransaction)
	trxRouter.Get("/all", h.getAllTransactions)
	trxRouter.Get("/all/:block", h.getTransactionsByBlockId)
	trxRouter.Get("/:id", h.getAllTransactions)
}
