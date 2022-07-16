package api

import (
	"fmt"
	"github.com/fatih/color"
	"github.com/gofiber/fiber/v2"
	"log"
)

const (
	version     = "0.0.1"
	website     = "https://www.bjungle.net"
	banner      = `BLion api consenso`
	description = `BLion api consenso - %s - Port: %s
by BJungle 
Version: %s
%s`
)

type server struct {
	listening string
	app       string
	fb        *fiber.App
}

func newServer(listening int, app string, fb *fiber.App) *server {
	return &server{fmt.Sprintf(":%d", listening), app, fb}
}

func (srv *server) Start() {
	color.Blue(banner)
	color.Cyan(fmt.Sprintf(description, srv.app, srv.listening, version, website))
	log.Fatal(srv.fb.Listen(srv.listening))
}
