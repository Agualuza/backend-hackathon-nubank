package main

import (
	"bank/conf"
	"bank/endpoint"
	"os"

	"github.com/tkanos/gonfig"

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

	//MakeTransaction
	e.GET("/maketransaction", endpoint.MakeTransaction)

	e.GET("/persona", endpoint.Persona)

	//Balance
	e.GET("/balance", endpoint.Balance)

	//LoadPersonas
	e.GET("/loadpersonas", endpoint.LoadPersonas)

	//LoadCategories
	e.GET("/loadcategories", endpoint.LoadCategories)

	//SafeBuy
	e.GET("/safebuy", endpoint.SafeBuy)

	e.Logger.Fatal(e.Start(":" + os.Getenv("PORT")))
}
