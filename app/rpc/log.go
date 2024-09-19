package rpc

import (
	"github.com/AnumaNetwork/anumad-testnet/infrastructure/logger"
	"github.com/AnumaNetwork/anumad-testnet/util/panics"
)

var log = logger.RegisterSubSystem("RPCS")
var spawn = panics.GoroutineWrapperFunc(log)
