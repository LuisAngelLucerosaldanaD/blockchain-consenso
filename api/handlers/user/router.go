package user

import (
	"bjungle-consenso/internal/middleware"
	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
)

func RouterUser(app *fiber.App, db *sqlx.DB, txID string) {
	h := handlerUser{DB: db, TxID: txID}
	api := app.Group("/api")
	v1 := api.Group("/v1")
	user := v1.Group("/user")
	user.Get("/accounting/:wallet", middleware.JWTProtected(), h.GetAccountByWalletID)
	user.Get("/freeze-money/:wallet", middleware.JWTProtected(), h.GetFreezeMoney)
}
