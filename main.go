package main

import (
	"log"

	"productlist/app"
	"productlist/config"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	config := config.NewConfig()
	app.ConfigAndRunApp(config)

}
