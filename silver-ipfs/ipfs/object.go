package ipfs

import "github.com/boltdb/bolt"

type IPFSObject struct {
	FileHash  []byte
	Name      []byte
	Size      int64
	Timestamp int64
	Db        *bolt.DB
}
