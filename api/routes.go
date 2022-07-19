package api

import (
	"bjungle-consenso/api/handlers/miner"
	"bjungle-consenso/api/handlers/participants"
	"github.com/ansrivas/fiberprometheus/v2"
	swagger "github.com/arsmn/fiber-swagger/v2"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

// @title OnlyOne Smart Contract.
// @version 1.0
// @description Documentation Smart Contract.
// @termsOfService https://www.bjungle.net/terms/
// @contact.name API Support.
// @contact.email info@bjungle.net
// @license.name Software Owner
// @license.url https://www.bjungle.net/terms/licenses
// @host localhost:50025
// @BasePath /
func routes(db *sqlx.DB, loggerHttp bool, allowedOrigins string) *fiber.App {
	app := fiber.New()

	prometheus := fiberprometheus.New("OnlyOne Smart Contract")
	prometheus.RegisterAt(app, "/metrics")

	app.Get("/swagger/*", swagger.New(swagger.Config{
		URL:         "/swagger/doc.json",
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
}
