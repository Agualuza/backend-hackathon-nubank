package main

import (
	"bank/conf"
	"bank/endpoint"
	"github.com/tkanos/gonfig"

	//"bank/database"
	"github.com/labstack/echo"
)

var configuration = conf.Configuration{}

func init() {
	err := gonfig.GetConf("./conf/conf.json", &configuration)

	if err != nil {
		panic(err)
	}
}

//main contains all API endpoints
func main() {
	e := echo.New()

	//Login
	e.GET("/login", endpoint.Login)

	//Register
	e.GET("/register", endpoint.Register)

	//SafeBuy
	e.GET("/safebuy", endpoint.SafeBuy)

	//MakeTransaction
	e.GET("/maketransaction", endpoint.MakeTransaction)

	e.Logger.Fatal(e.Start(configuration.Server.Port))
}
