package appmessage

import (
	"github.com/AnumaNetwork/anumad-testnet/domain/consensus/model/externalapi"
)

// MsgIBDChainBlockLocator implements the Message interface and represents a anuma
// locator message. It is used to find the blockLocator of a peer that is
// syncing with you.
type MsgIBDChainBlockLocator struct {
	baseMessage
	BlockLocatorHashes []*externalapi.DomainHash
}

// Command returns the protocol command string for the message. This is part
// of the Message interface implementation.
func (msg *MsgIBDChainBlockLocator) Command() MessageCommand {
	return CmdIBDChainBlockLocator
}

// NewMsgIBDChainBlockLocator returns a new anuma locator message that conforms to
// the Message interface. See MsgBlockLocator for details.
func NewMsgIBDChainBlockLocator(locatorHashes []*externalapi.DomainHash) *MsgIBDChainBlockLocator {
	return &MsgIBDChainBlockLocator{
		BlockLocatorHashes: locatorHashes,
	}
}
