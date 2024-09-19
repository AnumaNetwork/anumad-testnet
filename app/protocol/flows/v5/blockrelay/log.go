package blockrelay

import (
	"github.com/AnumaNetwork/anumad-testnet/infrastructure/logger"
	"github.com/AnumaNetwork/anumad-testnet/util/panics"
)

var log = logger.RegisterSubSystem("PROT")
var spawn = panics.GoroutineWrapperFunc(log)
