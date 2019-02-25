package types

import (
	usercryto "github.com/chain-dev/bschain/crypto"
	amino "github.com/tendermint/go-amino"
	cryptoAmino "github.com/tendermint/tendermint/crypto/encoding/amino"
)

var cdc = amino.NewCodec()

func init() {
	RegisterBlockAmino(cdc)
}

func RegisterBlockAmino(cdc *amino.Codec) {
	cryptoAmino.RegisterAmino(cdc)
	usercryto.RegisterCodec(cdc)
	RegisterEvidences(cdc)
}

// GetCodec returns a codec used by the package. For testing purposes only.
func GetCodec() *amino.Codec {
	return cdc
}

// For testing purposes only
func RegisterMockEvidencesGlobal() {
	RegisterMockEvidences(cdc)
}
