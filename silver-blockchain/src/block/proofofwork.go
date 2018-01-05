package block

import (
	"bytes"
	"crypto/sha256"
	"fmt"
	"go-labs/silver-blockchain/src/utils"
	"math"
	"math/big"
	"time"
)

const targetBits = 18

// Number once最大值
var maxNonce = math.MaxInt64

// 工作量证明结构体
type ProofOfWork struct {
	block  *Block
	target *big.Int
}

func NewProofOfWork(b *Block) *ProofOfWork {
	target := big.NewInt(1)

	target.Lsh(target, uint(256-targetBits))

	pow := &ProofOfWork{b, target}

	return pow
}

func (pow *ProofOfWork) prepareData(nonce int) []byte {
	data := bytes.Join(
		[][]byte{
			pow.block.PrevBlockHash,
			pow.block.HashTransaction(),
			utils.IntToHex(int64(pow.block.Timestamp)),
			utils.IntToHex(int64(targetBits)),
			utils.IntToHex(int64(nonce)),
		}, []byte{})

	return data
}

func (pow *ProofOfWork) Run() (int, []byte) {
	var hashInt big.Int
	var hash [32]byte
	nonce := 0
	startTime := time.Now()

	fmt.Printf("Mining the block containing [#%d]\n", pow.block.Id)

	for nonce < maxNonce {
		data := pow.prepareData(nonce)

		hash = sha256.Sum256(data)
		fmt.Printf("\rhash: %x, nonce: %d", hash, nonce)
		hashInt.SetBytes(hash[:])

		if hashInt.Cmp(pow.target) == -1 {
			fmt.Println()
			fmt.Printf("Mining time: [#%d] %s", pow.block.Id, time.Since(startTime))
			break
		} else {
			nonce++
		}
	}

	fmt.Print("\n\n")

	return nonce, hash[:]
}

// 验证工作量证明
func (pow *ProofOfWork) Validate() bool {
	var hashInt big.Int

	data := pow.prepareData(pow.block.Nonce)
	hash := sha256.Sum256(data)
	hashInt.SetBytes(hash[:])

	isValid := hashInt.Cmp(pow.target) == -1

	return isValid
}
