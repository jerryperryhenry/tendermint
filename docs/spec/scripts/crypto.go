package main

import (
	"fmt"
	"os"

	usercryto "github.com/chain-dev/bschain/crypto"
	amino "github.com/tendermint/go-amino"
	cryptoAmino "github.com/tendermint/tendermint/crypto/encoding/amino"
)

func main() {
	cdc := amino.NewCodec()
	cryptoAmino.RegisterAmino(cdc)
	usercryto.RegisterCodec(cdc)
	cdc.PrintTypes(os.Stdout)
	fmt.Println("")
}
