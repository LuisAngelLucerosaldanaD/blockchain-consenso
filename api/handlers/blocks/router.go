package blocks

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
)

func RouterBlocks(app *fiber.App, db *sqlx.DB, txID string) {
	h := handlerBlocks{DB: db, TxID: txID}
	api := app.Group("/api")
	v1 := api.Group("/v1")
	partRouter := v1.Group("/block")
	partRouter.Get("/get-all/:limit/:offset", h.GetAllBlocks)
	partRouter.Get("/current-lottery", h.GetCurrentLottery)
	partRouter.Get("/:id", h.GetBlockById)
}
