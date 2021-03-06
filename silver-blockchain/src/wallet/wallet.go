package wallet

import (
	"bytes"
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/sha256"
	"github.com/go-clog/clog"
	"github.com/wuleying/go-labs/silver-blockchain/src/util"
	"golang.org/x/crypto/ripemd160"
)

type Wallet struct {
	PrivateKey ecdsa.PrivateKey
	PublicKey  []byte
}

func (w Wallet) GetAddress() []byte {
	publicKeyHash := HashPublicKey(w.PublicKey)

	versionedPayload := append([]byte{util.VERSION}, publicKeyHash...)
	checksum := checksum(versionedPayload)

	fullPayload := append(versionedPayload, checksum...)
	address := util.Base58Encode(fullPayload)

	return address
}

func NewWallet() *Wallet {
	private, public := newKeyPair()

	wallet := Wallet{private, public}

	return &wallet
}

func HashPublicKey(publicKey []byte) []byte {
	publicSha256 := sha256.Sum256(publicKey)

	ripemd160Hasher := ripemd160.New()

	_, err := ripemd160Hasher.Write(publicSha256[:])
	if err != nil {
		clog.Fatal(util.CLOG_SKIP_DISPLAY_INFO, err.Error())
	}

	publicRipemd160 := ripemd160Hasher.Sum(nil)

	return publicRipemd160
}

func ValidateAddress(address string) bool {
	publicKeyHash := util.Base58Decode([]byte(address))
	actualChecksum := publicKeyHash[len(publicKeyHash)-util.ADDRESS_CHECKSUM_LEN:]
	version := publicKeyHash[0]
	publicKeyHash = publicKeyHash[1 : len(publicKeyHash)-util.ADDRESS_CHECKSUM_LEN]
	targetChecksum := checksum(append([]byte{version}, publicKeyHash...))

	return bytes.Compare(actualChecksum, targetChecksum) == 0
}

func newKeyPair() (ecdsa.PrivateKey, []byte) {
	curve := elliptic.P256()
	private, err := ecdsa.GenerateKey(curve, rand.Reader)
	if err != nil {
		clog.Fatal(util.CLOG_SKIP_DISPLAY_INFO, err.Error())
	}

	public := append(private.PublicKey.X.Bytes(), private.PublicKey.Y.Bytes()...)

	return *private, public
}

func checksum(payload []byte) []byte {
	firstSha := sha256.Sum256(payload)
	secondSha := sha256.Sum256(firstSha[:])

	return secondSha[:util.ADDRESS_CHECKSUM_LEN]
}
