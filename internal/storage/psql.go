package storage

import (
	//"database/sql"
	"fmt"

	"github.com/blackflagsoftware/prognos/config"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

var PsqlDB *sqlx.DB
var postgresConnection string

func PostgresInit() *sqlx.DB {
	if PsqlDB == nil {
		postgresConnection = GetPostgresConnection()
		var err error
		PsqlDB, err = sqlx.Connect("postgres", postgresConnection)
		if err != nil {
			panic(fmt.Sprintf("Could not connect to the DB host: %s*****; %s", string(config.DBHost[:6]), err))
		}
	}
	return PsqlDB
}

func GetPostgresConnection() string {
	postgresConnection = fmt.Sprintf("user=%s password=%s dbname=%s host=%s sslmode=disable", config.DBUser, config.DBPass, config.DBDB, config.DBHost)
	if config.DBPass == "" {
		postgresConnection = fmt.Sprintf("user=%s dbname=%s host=%s sslmode=disable", config.DBUser, config.DBDB, config.DBHost)
	}
	return postgresConnection
}
