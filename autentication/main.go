package main

import (
	"autentication/db"
	"autentication/routes"
)

func main() {
	db.Init()

	e := routes.Init()

	e.Logger.Fatal(e.Start(":1234"))
}
