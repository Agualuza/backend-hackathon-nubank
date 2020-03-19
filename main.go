package main

import (
	"bank/conf"
	"bank/endpoint"
	"os"

	"github.com/labstack/echo"
	"github.com/tkanos/gonfig"
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

	//Persona
	e.GET("/persona", endpoint.Persona)

	//Balance
	e.GET("/balance", endpoint.Balance)

	//LoadPersonas
	e.GET("/loadpersonas", endpoint.LoadPersonas)

	//LoadCategories
	e.GET("/loadcategories", endpoint.LoadCategories)

	//SafeBuy
	e.GET("/safebuy", endpoint.SafeBuyWithIA)
	
	//LoadQuestions
	e.GET("/loadquestions", endpoint.LoadQuestions)

	//LoadPosts
	e.GET("/loadposts", endpoint.LoadPosts)
	
	
	//LoadPersonaPosts
	e.GET("/loadpersonaposts", endpoint.LoadPersonaPosts)
	
	e.Logger.Fatal(e.Start(":" + os.Getenv("PORT")))
}
