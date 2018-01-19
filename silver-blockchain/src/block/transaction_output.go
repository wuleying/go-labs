package block

import (
	"bytes"
	"encoding/gob"
	"github.com/go-clog/clog"
	"go-labs/silver-blockchain/src/utils"
)

type TOutput struct {
	Value         int
	PublicKeyHash []byte
}

type TOutputs struct {
	Outputs []TOutput
}

func (out *TOutput) Lock(address []byte) {
	publicHash := utils.Base58Decode(address)
	publicHash = publicHash[1 : len(publicHash)-4]
	out.PublicKeyHash = publicHash
}

func (out *TOutput) IsLockedWithKey(publicKeyHash []byte) bool {
	return bytes.Compare(out.PublicKeyHash, publicKeyHash) == 0
}

func (outs TOutputs) Serialize() []byte {
	var buff bytes.Buffer

	encoder := gob.NewEncoder(&buff)
	err := encoder.Encode(outs)
	if err != nil {
		clog.Fatal(0, err.Error())
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
		clog.Fatal(0, err.Error())
	}

	return outputs
}
