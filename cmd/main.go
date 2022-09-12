package main

import (
	"mesto/pkg/logger"
	"os"
	"os/signal"
	"syscall"

	_ "mesto/docs"

	"mesto/server"
)

// @title Документация API Mesto
// @version 1.0

// @contact.name - Поддержка API
// @contact.email kuzuru.dev@gmail.com
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// Параметр host нужно поменять на расшаренный IP
// @host localhost
// @BasePath /

// @securityDefinitions.apiKey ApiKeyAuth
// @in header
// @name Authorization
func main() {
	// Initializing logger
	logger.Init()

	// Starting server
	server.Run()

	// Waiting for quit signal on exit
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM, syscall.SIGHUP, syscall.SIGQUIT)

	<-quit
}
