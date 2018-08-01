package rest

import (
	"github.com/ya-enot/go-crypto/keys"
	wire "github.com/ya-enot/go-wire"

	ctypes "github.com/ya-enot/tendermint/rpc/core/types"

	sdk "github.com/ya-enot/cosmos-sdk"
	"github.com/ya-enot/cosmos-sdk/client/commands"
	keycmd "github.com/ya-enot/cosmos-sdk/client/commands/keys"
)

// PostTx is same as a tx
func PostTx(tx sdk.Tx) (*ctypes.ResultBroadcastTxCommit, error) {
	packet := wire.BinaryBytes(tx)
	// post the bytes
	node := commands.GetNode()
	return node.BroadcastTxCommit(packet)
}

// SignTx will modify the tx in-place, adding a signature if possible
func SignTx(name, pass string, tx sdk.Tx) error {
	if sign, ok := tx.Unwrap().(keys.Signable); ok {
		manager := keycmd.GetKeyManager()
		return manager.Sign(name, pass, sign)
	}
	return nil
}
