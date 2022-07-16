package dbx

import (
	"bjungle-consenso/internal/env"
	"bjungle-consenso/internal/logger"
	"fmt"
	"strings"
	"sync"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

var (
	dbx  *sqlx.DB
	once sync.Once
)

func init() {
	once.Do(func() {
		setConnection()
	})
}

func setConnection() {
	var err error
	c := env.NewConfiguration()

	dbx, err = sqlx.Open(c.DB.Engine, connectionString())
	if err != nil {
		logger.Error.Printf("couldn't connect to database: %v", err)
		panic(err)
	}
	err = dbx.Ping()
	if err != nil {
		logger.Error.Printf("couldn't connect to database: %v", err)
		panic(err)
	}

}

//TODO encrypt-decrypt password
func connectionString() string {
	var host, database, username, password, instance string
	var port int
	c := env.NewConfiguration()
	host = c.DB.Server
	database = c.DB.Name
	username = c.DB.User
	password = c.DB.Password
	instance = c.DB.Instance
	port = c.DB.Port

	switch strings.ToLower(c.DB.Engine) {
	case "postgres":
		return fmt.Sprintf("dbname=%s user=%s password=%s host=%s port=%d sslmode=disable", database, username, password, host, port)
	case "sqlserver":
		return fmt.Sprintf(
			"server=%s\\%s;User id=%s;database=%s;password=%s;port=%d", host, instance, username, database, password, port)
	default:
		logger.Error.Printf("database engine is not configured")
	}
	return ""
}

func GetConnection() *sqlx.DB {
	return dbx
}
