package main

import (
	"ms-template-golang/internal/app"
	"ms-template-golang/internal/config"
	"ms-template-golang/internal/logger"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	config.SetupConfig()
	var (
		cfg = config.Cfg()
	)
	logger.SetupLogger(cfg.ENV)
	var (
		application *app.App
		log         = logger.Logger()
	)
	application = app.New()
	go func() {
		application.Run()
	}()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGTERM, syscall.SIGINT)

	<-stop

	application.Stop()
	log.Info("Gracefully stopped")
}
