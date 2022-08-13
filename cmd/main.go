package main

import (
	"mesto/pkg/logger"
	"os"
	"os/signal"
	"syscall"

	"mesto/server"
)

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
