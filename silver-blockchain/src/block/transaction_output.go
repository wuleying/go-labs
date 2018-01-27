package block

import (
	"bytes"
	"encoding/gob"
	"github.com/go-clog/clog"
	"go-labs/silver-blockchain/src/util"
)

type TOutput struct {
	Value         int
	PublicKeyHash []byte
}

type TOutputs struct {
	Outputs []TOutput
}

func (out *TOutput) Lock(address []byte) {
	publicKeyHash := util.Base58Decode(address)
	publicKeyHash = publicKeyHash[1 : len(publicKeyHash)-util.ADDRESS_CHECKSUM_LEN]
	out.PublicKeyHash = publicKeyHash
}

func (out *TOutput) IsLockedWithKey(publicKeyHash []byte) bool {
	return bytes.Compare(out.PublicKeyHash, publicKeyHash) == 0
}

func (outs TOutputs) Serialize() []byte {
	var buff bytes.Buffer

	encoder := gob.NewEncoder(&buff)
	err := encoder.Encode(outs)
	if err != nil {
		clog.Fatal(2, err.Error())
	}

	return buff.Bytes()
}

func NewTOutput(value int, address string) *TOutput {
	out := &TOutput{value, nil}
	out.Lock([]byte(address))

	return out
}

func DeserializeOutputs(data []byte) TOutputs {
	var outputs TOutputs

	decoder := gob.NewDecoder(bytes.NewBuffer(data))
	err := decoder.Decode(&outputs)
	if err != nil {
		clog.Fatal(2, err.Error())
	}

	return outputs
}
