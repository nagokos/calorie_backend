package main

import (
	"log"

	"github.com/nagokos/calorie_backend/config"
	"github.com/nagokos/calorie_backend/server"
	"github.com/nagokos/calorie_backend/utils"
)

func main() {
	utils.LogginsSettings(config.Config.Log)
	if err := server.StartWebServer(); err != nil {
		log.Fatal(err)
	}
}
