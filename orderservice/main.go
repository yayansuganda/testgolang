package main

import (
	"orderservice/config"
	"orderservice/routes"
)

func main() {
	config.ConnectDB()

	e := routes.Init()

	e.Logger.Fatal(e.Start(":1235"))
}
