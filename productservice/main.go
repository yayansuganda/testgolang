package main

import (
	"productservice/config"
	"productservice/routes"
)

func main() {
	config.ConnectDB()

	e := routes.Init()

	e.Logger.Fatal(e.Start(":1236"))
}
