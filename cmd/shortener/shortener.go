package main

import (
	"flag"
	"github.com/fasthttp/router"
	"github.com/kaverhovsky/tinyUrl/pkg/common"
	"github.com/kaverhovsky/tinyUrl/pkg/httpserver"
	"net"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	configPath := flag.String("config", "", "Config path")
	flag.Parse()
	logger := common.NewLogger("development", "info")
	mainLogger := logger.Named("main")
	cfg, err := common.ReadConfig(*configPath, logger)
	if err != nil {
		mainLogger.Fatal(err.Error())
	}
	r := router.New()
	srv := httpserver.NewServer(cfg, logger, r)
	ln, err := net.Listen("tcp4", cfg.Listen)
	if err != nil {
		mainLogger.Fatal(err.Error())
	}
	go srv.Run(ln)
	sigChan := make(chan os.Signal)
	signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM, syscall.SIGINT)
	<-sigChan
	srv.Shutdown()
	mainLogger.Info("Server closed")
}
