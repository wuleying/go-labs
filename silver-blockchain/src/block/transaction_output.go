package block

import (
	"bytes"
	"go-labs/silver-blockchain/src/utils"
)

type TOutput struct {
	Value         int
	PublicKeyHash []byte
}

func (out *TOutput) Lock(address []byte) {
	publicHash := utils.Base58Decode(address)
	publicHash = publicHash[1 : len(publicHash)-4]
}

func (out *TOutput) IsLockedWithKey(publicKeyHash []byte) bool {
	return bytes.Compare(out.PublicKeyHash, publicKeyHash) == 0
}

func NewTOutput(value int, address string) *TOutput {
	out := &TOutput{value, nil}
	out.Lock([]byte(address))

	return out
}
