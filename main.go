package main

import (
	"altastore/config"
	"altastore/routes"
)

func main() {
	config.InitDB()
	e := routes.Start()
	e.Logger.Fatal(e.Start(":8000"))
}
