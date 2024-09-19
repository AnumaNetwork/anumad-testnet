package standalone

import (
	"github.com/AnumaNetwork/anumad-testnet/infrastructure/logger"
	"github.com/AnumaNetwork/anumad-testnet/util/panics"
)

var log = logger.RegisterSubSystem("NTAR")
var spawn = panics.GoroutineWrapperFunc(log)
