package database

import (
	"bank/conf"
	"database/sql"
	_ "database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/tkanos/gonfig"
	"os"
)

func ConnectDB() *sql.DB {
	configuration := conf.Configuration{}
	err := gonfig.GetConf("./conf/conf.json", &configuration)

	if err != nil {
		panic(err)
	}

	db, err := sql.Open(configuration.DB.Driver, os.Getenv("DBConnQuery"))

	return db

}