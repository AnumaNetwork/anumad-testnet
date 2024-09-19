package rpchandlers

import (
	"github.com/AnumaNetwork/anumad-testnet/app/appmessage"
	"github.com/AnumaNetwork/anumad-testnet/app/rpc/rpccontext"
	"github.com/AnumaNetwork/anumad-testnet/infrastructure/network/netadapter/router"
)

// HandleGetHeaders handles the respectively named RPC command
func HandleGetHeaders(context *rpccontext.Context, _ *router.Router, request appmessage.Message) (appmessage.Message, error) {
	response := &appmessage.GetHeadersResponseMessage{}
	response.Error = appmessage.RPCErrorf("not implemented")
	return response, nil
}
