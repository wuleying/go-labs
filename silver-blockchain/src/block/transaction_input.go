package block

import (
	"bytes"
	"go-labs/silver-blockchain/src/wallet"
)

type TInput struct {
	Id        []byte
	Out       int
	Signature []byte
	PublicKey []byte
}

func (in *TInput) UsesKey(publicKeyHash []byte) bool {
	lockingHash := wallet.HashPublicKey(in.PublicKey)

	return bytes.Compare(lockingHash, publicKeyHash) == 0
}
