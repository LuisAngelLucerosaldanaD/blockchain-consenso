package miner

import (
	"bjungle-consenso/internal/middleware"
	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
)

func RouterMiner(app *fiber.App, db *sqlx.DB, txID string) {
	h := handlerMiner{DB: db, TxID: txID}
	api := app.Group("/api")
	v1 := api.Group("/v1")
	partRouter := v1.Group("/miner")
	partRouter.Post("/register-mined", middleware.JWTProtected(), h.RegisterHashMined)
	partRouter.Get("/block-to-mine", middleware.JWTProtected(), h.GetBlockToMine)
	partRouter.Get("/hash-mined", h.GetHashMined)
	partRouter.Get("/:id", h.GetMiner)
}
