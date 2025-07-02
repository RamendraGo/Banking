package main

import (
	"github.com/RamendraGo/Banking/app"
	"github.com/RamendraGo/Banking/logger"
)

func main() {

	logger.Info("Starting the application..")
	app.Start()

}
