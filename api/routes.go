package api

import (
	"bjungle-consenso/api/handlers/blocks"
	"bjungle-consenso/api/handlers/miner"
	"bjungle-consenso/api/handlers/participants"
	"bjungle-consenso/api/handlers/sign"
	"bjungle-consenso/api/handlers/transactions"
	"bjungle-consenso/api/handlers/user"
	"bjungle-consenso/api/handlers/validator"
	_ "bjungle-consenso/docs"
	"github.com/ansrivas/fiberprometheus/v2"
	swagger "github.com/arsmn/fiber-swagger/v2"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

func routes(db *sqlx.DB, loggerHttp bool, allowedOrigins string) *fiber.App {
	app := fiber.New()

	prometheus := fiberprometheus.New("BLion Consenso")
	prometheus.RegisterAt(app, "/metrics")

	app.Get("/swagger/*", swagger.New(swagger.Config{
		URL:         "/swagger/doc.json",
		Title:       "BLion Consenso",
		DeepLinking: false,
	}))

	app.Use(recover.New())
	app.Use(prometheus.Middleware)
	app.Use(cors.New(cors.Config{
		AllowOrigins: allowedOrigins,
		AllowHeaders: "Origin, X-Requested-With, Content-Type, Accept, Authorization",
		AllowMethods: "GET,POST",
	}))
	if loggerHttp {
		app.Use(logger.New())
	}
	TxID := uuid.New().String()

	loadRoutes(app, db, TxID)

	return app
}

func loadRoutes(app *fiber.App, db *sqlx.DB, TxID string) {
	participants.RouterParticipants(app, db, TxID)
	miner.RouterMiner(app, db, TxID)
	validator.RouterValidators(app, db, TxID)
	user.RouterUser(app, db, TxID)
	blocks.RouterBlocks(app, db, TxID)
	transactions.RouterTransaction(app, db, TxID)
	sign.RouterSign(app, db, TxID)
}
