package main

import (
	"github.com/AnumaNetwork/anumad-testnet/domain/consensus/model/externalapi"
	"github.com/AnumaNetwork/anumad-testnet/domain/consensus/utils/consensushashing"
	"github.com/AnumaNetwork/anumad-testnet/stability-tests/common/rpc"
	"github.com/pkg/errors"
)

func checkTopBlockIsTip(rpcClient *rpc.Client, topBlock *externalapi.DomainBlock) error {
	selectedTipHashResponse, err := rpcClient.GetSelectedTipHash()
	if err != nil {
		return err
	}

	topBlockHash := consensushashing.BlockHash(topBlock)
	if selectedTipHashResponse.SelectedTipHash != topBlockHash.String() {
		return errors.Errorf("selectedTipHash is '%s' while expected to be topBlock's hash `%s`",
			selectedTipHashResponse.SelectedTipHash, topBlockHash)
	}

	return nil
}
