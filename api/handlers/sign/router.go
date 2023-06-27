package sign

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
)

func RouterSign(app *fiber.App, db *sqlx.DB, txID string) {
	h := handlerSign{DB: db, TxID: txID}
	api := app.Group("/api")
	v1 := api.Group("/v1")
	signRouter := v1.Group("/sign")
	signRouter.Post("/create", h.createSign)
	signRouter.Post("/export", h.exportSign)
}
