package block

import (
	"bytes"
	"crypto/sha256"
	"encoding/gob"
	"encoding/hex"
	"fmt"
	"go-labs/silver-blockchain/src/wallet"
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

func (t Transaction) Serialize() []byte {
	var encoded bytes.Buffer

	encoder := gob.NewEncoder(&encoded)
	err := encoder.Encode(t)
	if err != nil {
		log.Panic(err)
	}

	return encoded.Bytes()
}

func (t *Transaction) Hash() []byte {
	var hash [32]byte

	tCopy := *t
	tCopy.Id = []byte{}

	hash = sha256.Sum256(tCopy.Serialize())

	return hash[:]
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

func NewCoinBase(to string, data string) *Transaction {
	if data == "" {
		data = fmt.Sprintf("Reward to '%s'", to)
	}

	tIn := TInput{[]byte{}, -1, nil, []byte(data)}
	tOut := NewTOutput(subsidy, to)

	t := Transaction{nil, []TInput{tIn}, []TOutput{*tOut}}
	t.Id = t.Hash()

	return &t
}

func NewUTXOTransaction(from string, to string, amount int, bc *BlockChain) *Transaction {
	var inputs []TInput
	var outputs []TOutput

	wallets, err := wallet.NewWallets()
	if err != nil {
		log.Panic(err)
	}
	walletFrom := wallets.GetWallet(from)
	publicKeyHash := wallet.HashPublicKey(walletFrom.PublicKey)
	account, vaildOutputs := bc.FindSpendableOutputs(publicKeyHash, amount)

	if account < amount {
		log.Panic("Error: Not enough funds")
	}

	for id, outs := range vaildOutputs {
		tId, err := hex.DecodeString(id)
		if err != nil {
			log.Panic(err)
		}

		for _, out := range outs {
			inputs = append(inputs, TInput{tId, out, nil, walletFrom.PublicKey})
		}
	}

	outputs = append(outputs, *NewTOutput(amount, to))

	if account > amount {
		outputs = append(outputs, *NewTOutput(account-amount, from))
	}

	t := Transaction{nil, inputs, outputs}
	t.SetId()

	return &t
}
