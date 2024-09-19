package main

import (
	"github.com/AnumaNetwork/anumad-testnet/infrastructure/logger"
	"github.com/AnumaNetwork/anumad-testnet/util/panics"
)

var (
	backendLog = logger.NewBackend()
	log        = backendLog.Logger("RORG")
	spawn      = panics.GoroutineWrapperFunc(log)
)
