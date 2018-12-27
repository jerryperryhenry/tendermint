package privval

import (
	usercryto "github.com/chain-dev/bschain/crypto"
	"github.com/tendermint/go-amino"
	cryptoAmino "github.com/tendermint/tendermint/crypto/encoding/amino"
)

var cdc = amino.NewCodec()

func init() {
	usercryto.RegisterCodec(cdc)
	cryptoAmino.RegisterAmino(cdc)
	RegisterRemoteSignerMsg(cdc)
}
