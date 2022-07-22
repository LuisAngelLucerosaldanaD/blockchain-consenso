package validator

import (
	"bjungle-consenso/internal/middleware"
	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
)

func RouterValidators(app *fiber.App, db *sqlx.DB, txID string) {
	h := handlerValidators{DB: db, TxID: txID}
	api := app.Group("/api")
	v1 := api.Group("/v1")
	partRouter := v1.Group("/validators")
	partRouter.Post("/register", middleware.JWTProtected(), h.RegisterVoteValidator)
}
