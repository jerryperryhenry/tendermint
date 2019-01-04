package conn

import (
	usercryto "github.com/chain-dev/bschain/crypto"
	"github.com/tendermint/go-amino"
	cryptoAmino "github.com/tendermint/tendermint/crypto/encoding/amino"
)

var cdc *amino.Codec = amino.NewCodec()

func init() {
	cryptoAmino.RegisterAmino(cdc)
	usercryto.RegisterCodec(cdc)
	RegisterPacket(cdc)
}
