package httpserver

import (
	"github.com/kaverhovsky/tinyUrl/pkg/common"
	"go.uber.org/zap"
)

type Server struct {
	config common.Config
	zap.Logger
}
