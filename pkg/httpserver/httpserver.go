package httpserver

import (
	"github.com/fasthttp/router"
	"github.com/kaverhovsky/tinyUrl/pkg/common"
	"github.com/valyala/fasthttp"
	"go.uber.org/zap"
	"net"
)

type Server struct {
	config *common.Config
	logger *zap.Logger
	server *fasthttp.Server
}

func NewServer(config *common.Config, logger *zap.Logger, router *router.Router) *Server {
	logger = logger.Named("httpserver")
	server := &fasthttp.Server{
		Name:               "url-shortener-httpserver",
		Handler:            router.Handler,
		MaxConnsPerIP:      config.MaxConnsPerIP,
		MaxRequestsPerConn: config.MaxRequestsPerConn,
		MaxRequestBodySize: config.MaxRequestBodySize * 1024 * 1024,
		WriteTimeout:       config.WriteTimeout,
		ReadTimeout:        config.ReadTimeout,
		TCPKeepalive:       config.TCPKeepalive,
		IdleTimeout:        config.IdleTimeout,
	}
	return &Server{
		config: config,
		logger: logger,
		server: server,
	}
}

func (srv *Server) Run(ln net.Listener) {
	srv.logger.Info("Running HTTP server...")
	if err := srv.server.Serve(ln); err != nil {
		srv.logger.Fatal(err.Error())
	}
}
func (srv *Server) Shutdown() {
	srv.logger.Info("Shutting down HTTP server...")
	if err := srv.server.Shutdown(); err != nil {
		srv.logger.Fatal(err.Error())
	}
}
