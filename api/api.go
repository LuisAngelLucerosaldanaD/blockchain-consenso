package api

import "bjungle-consenso/internal/dbx"

func Start(port int, app string, loggerHttp bool, allowedOrigins string) {
	db := dbx.GetConnection()

	defer func() {
		err := db.Close()
		if err != nil {
			return
		}
	}()

	r := routes(db, loggerHttp, allowedOrigins)
	server := newServer(port, app, r)
	server.Start()
}
