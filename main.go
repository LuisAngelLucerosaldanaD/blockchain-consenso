package main

import (
	"bjungle-consenso/api"
	"bjungle-consenso/internal/env"
)

func main() {
	c := env.NewConfiguration()
	api.Start(c.App.Port, c.App.ServiceName, c.App.LoggerHttp, c.App.AllowedDomains)
}
