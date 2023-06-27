package main

import (
	"bjungle-consenso/api"
	"bjungle-consenso/internal/env"
)

// @title BLion
// @version 1.1
// @description Documentación del API que conecta con el core de BLion
// @termsOfService https://www.bjungle.net/terms/
// @contact.name API Support
// @contact.email info@bjungle.net
// @contact.url https://www.bjungle.net
// @license.name Software Owner
// @license.url https://www.bjungle.net/terms/licenses
// @host https://sandbox.btigersystem.net/blion/dev/consenso
// @tag.name Block
// @tag.description Métodos para la gestión de los bloques
// @tag.name Miner
// @tag.description Métodos para la gestión de los mineros
// @tag.name Transacción
// @tag.description Métodos para la gestión de las transacciones
// @tag.Name Participants
// @tag.description Métodos para la gestión de participantes en la lotería
// @tag.Name User
// @tag.description Métodos para la gestión de usuarios
// @tag.Name Validator
// @tag.description Métodos para la gestión de validadores
// @tag.Name Sign
// @tag.description Métodos para la gestión de la firma de la transacción
// @BasePath /
func main() {
	c := env.NewConfiguration()
	api.Start(c.App.Port, c.App.ServiceName, c.App.LoggerHttp, c.App.AllowedDomains)
}
