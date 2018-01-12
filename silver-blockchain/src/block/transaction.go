package block

import (
	"bytes"
	"crypto/sha256"
	"encoding/gob"
	"encoding/hex"
	"fmt"
	"log"
)

// 挖出新块的奖励金
const subsidy = 10

type Transaction struct {
	Id  []byte
	In  []TInput
	Out []TOutput
}

func (t Transaction) IsCoinBase() bool {
	return len(t.In) == 1 && len(t.In[0].Id) == 0 && t.In[0].Out == -1
}

func (t Transaction) SetId() {
	var encoded bytes.Buffer
	var hash [32]byte

	encoder := gob.NewEncoder(&encoded)

	err := encoder.Encode(t)

	if err != nil {
		log.Panic(err)
	}

	hash = sha256.Sum256(encoded.Bytes())
	t.Id = hash[:]
}

/*
func (in TInput) CanUnlockOutputWith(data string) bool {
	return in.ScriptSig == data
}

func (out TOutput) CanBeUnlockWith(data string) bool {
	return out.ScriptPubKey == data
}
*/

func NewCoinBase(to string, data string) *Transaction {
	if data == "" {
		data = fmt.Sprintf("Reward to '%s'", to)
	}

	tIn := TInput{[]byte{}, -1, data}
	tOut := TOutput{subsidy, to}

	t := Transaction{nil, []TInput{tIn}, []TOutput{tOut}}
	t.SetId()

	return &t
}

func NewUTXOTransaction(from string, to string, amount int, bc *BlockChain) *Transaction {
	var inputs []TInput
	var outputs []TOutput

	account, vaildOutputs := bc.FindSpendableOutputs(from, amount)

	if account < amount {
		log.Panic("Error: Not enough funds")
	}

	for id, outs := range vaildOutputs {
		tId, err := hex.DecodeString(id)
		if err != nil {
			log.Panic(err)
		}

		for _, out := range outs {
			inputs = append(inputs, TInput{tId, out, from})
		}
	}

	outputs = append(outputs, TOutput{amount, to})

	if account > amount {
		outputs = append(outputs, TOutput{account - amount, from})
	}

	t := Transaction{nil, inputs, outputs}
	t.SetId()

	return &t
}
