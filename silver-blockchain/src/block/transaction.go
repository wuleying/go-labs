package block

import (
	"bytes"
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/sha256"
	"encoding/gob"
	"encoding/hex"
	"fmt"
	"github.com/go-clog/clog"
	"go-labs/silver-blockchain/src/wallet"
	"math/big"
)

// 挖出新块的奖励金
const subsidy = 10

type Transaction struct {
	Id  []byte
	In  []TInput
	Out []TOutput
}

func NewCoinBase(to string, data string) *Transaction {
	if data == "" {
		randData := make([]byte, 20)
		_, err := rand.Read(randData)
		if err != nil {
			clog.Fatal(2, err.Error())
		}

		data = fmt.Sprintf("%x", randData)
	}

	tIn := TInput{[]byte{}, -1, nil, []byte(data)}
	tOut := NewTOutput(subsidy, to)

	t := Transaction{nil, []TInput{tIn}, []TOutput{*tOut}}
	t.Id = t.Hash()

	return &t
}

func NewUTXOTransaction(from string, to string, amount int, UTXOSet *UTXOSet) *Transaction {
	var inputs []TInput
	var outputs []TOutput

	wallets, err := wallet.NewWallets()
	if err != nil {
		clog.Fatal(2, err.Error())
	}

	walletFrom := wallets.GetWallet(from)
	publicKeyHash := wallet.HashPublicKey(walletFrom.PublicKey)

	accumulated, validOutputs := UTXOSet.FindSpendableOutputs(publicKeyHash, amount)

	if accumulated < amount {
		clog.Fatal(2, "Not enough funds")
	}

	for id, outs := range validOutputs {
		tId, err := hex.DecodeString(id)
		if err != nil {
			clog.Fatal(2, err.Error())
		}

		for _, out := range outs {
			inputs = append(inputs, TInput{tId, out, nil, walletFrom.PublicKey})
		}
	}

	outputs = append(outputs, *NewTOutput(amount, to))

	if accumulated > amount {
		outputs = append(outputs, *NewTOutput(accumulated-amount, from))
	}

	t := Transaction{nil, inputs, outputs}
	t.Id = t.Hash()
	UTXOSet.BlockChain.SignTransaction(&t, walletFrom.PrivateKey)

	return &t
}

func (t Transaction) IsCoinBase() bool {
	return len(t.In) == 1 && len(t.In[0].Id) == 0 && t.In[0].Out == -1
}

func (t Transaction) Serialize() []byte {
	var encoded bytes.Buffer

	encoder := gob.NewEncoder(&encoded)
	err := encoder.Encode(t)
	if err != nil {
		clog.Fatal(2, err.Error())
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

func (t *Transaction) Sign(privateKey ecdsa.PrivateKey, prevTs map[string]Transaction) {
	if t.IsCoinBase() {
		return
	}

	for _, in := range t.In {
		if prevTs[hex.EncodeToString(in.Id)].Id == nil {
			clog.Fatal(2, "Previous transaction is not correct")
		}
	}

	tCopy := t.TrimmedCopy()

	for inIdx, in := range tCopy.In {
		prevT := prevTs[hex.EncodeToString(in.Id)]
		tCopy.In[inIdx].Signature = nil
		tCopy.In[inIdx].PublicKey = prevT.Out[in.Out].PublicKeyHash
		tCopy.Id = tCopy.Hash()
		tCopy.In[inIdx].PublicKey = nil

		r, s, err := ecdsa.Sign(rand.Reader, &privateKey, tCopy.Id)
		if err != nil {
			clog.Fatal(2, err.Error())
		}
		signature := append(r.Bytes(), s.Bytes()...)

		t.In[inIdx].Signature = signature
	}
}

func (t *Transaction) Verify(prevTs map[string]Transaction) bool {
	if t.IsCoinBase() {
		return true
	}

	for _, in := range t.In {
		if prevTs[hex.EncodeToString(in.Id)].Id == nil {
			clog.Fatal(2, "Previous transaction is not correct.")
		}
	}

	tCopy := t.TrimmedCopy()
	curve := elliptic.P256()

	for inIdx, in := range t.In {
		prevT := prevTs[hex.EncodeToString(in.Id)]
		tCopy.In[inIdx].Signature = nil
		tCopy.In[inIdx].PublicKey = prevT.Out[in.Out].PublicKeyHash
		tCopy.Id = tCopy.Hash()
		tCopy.In[inIdx].PublicKey = nil

		r := big.Int{}
		s := big.Int{}
		signLen := len(in.Signature)
		r.SetBytes(in.Signature[:(signLen / 2)])
		s.SetBytes(in.Signature[(signLen / 2):])

		x := big.Int{}
		y := big.Int{}
		keyLen := len(in.PublicKey)
		x.SetBytes(in.PublicKey[:(keyLen / 2)])
		y.SetBytes(in.PublicKey[(keyLen / 2):])

		rawPublicKey := ecdsa.PublicKey{curve, &x, &y}
		if ecdsa.Verify(&rawPublicKey, tCopy.Id, &r, &s) == false {
			return false
		}
	}

	return true
}

func (t *Transaction) TrimmedCopy() Transaction {
	var inputs []TInput
	var outputs []TOutput

	for _, in := range t.In {
		inputs = append(inputs, TInput{in.Id, in.Out, nil, nil})
	}

	for _, out := range t.Out {
		outputs = append(outputs, TOutput{out.Value, out.PublicKeyHash})
	}

	tCopy := Transaction{t.Id, inputs, outputs}

	return tCopy
}

func (t Transaction) SetId() {
	var encoded bytes.Buffer
	var hash [32]byte

	encoder := gob.NewEncoder(&encoded)

	err := encoder.Encode(t)
	if err != nil {
		clog.Fatal(2, err.Error())
	}

	hash = sha256.Sum256(encoded.Bytes())
	t.Id = hash[:]
}
