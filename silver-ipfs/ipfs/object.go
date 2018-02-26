package ipfs

import (
	"bytes"
	"encoding/gob"
	"github.com/boltdb/bolt"
	"github.com/go-clog/clog"
	"github.com/wuleying/go-labs/silver-ipfs/commands"
	"github.com/wuleying/go-labs/silver-ipfs/util"
	"os"
	"time"
)

type IPFSObject struct {
	FileHash  []byte
	Name      []byte
	Size      int64
	Timestamp int64
}

// IPFS对象
func NewObject() *IPFSObject {
	return &IPFSObject{[]byte{}, []byte{}, 0, 0}
}

// 保存数据
func (o *IPFSObject) Save(filePath string) (string, error) {
	fileHash, err := commands.AddFile(filePath)
	if err != nil {
		return "", err
	}

	o.FileHash = []byte(fileHash)
	o.Name = []byte("test")
	o.Size = 12
	o.Timestamp = time.Now().Unix()

	db, err := bolt.Open(util.DB_FILE_PATH, 0600, nil)
	if err != nil {
		clog.Fatal(util.CLOG_SKIP_DISPLAY_INFO, err.Error())
	}

	err = db.Update(func(tx *bolt.Tx) error {
		b, err := tx.CreateBucketIfNotExists([]byte(util.BLOCK_BUCKET_NAME))
		if err != nil {
			clog.Fatal(util.CLOG_SKIP_DISPLAY_INFO, err.Error())
		}

		err = b.Put(o.FileHash, o.Serialize())
		if err != nil {
			clog.Fatal(util.CLOG_SKIP_DISPLAY_INFO, err.Error())
		}

		return nil
	})

	if err != nil {
		return "", err
	}

	return fileHash, nil
}

// 序列化对象
func (o *IPFSObject) Serialize() []byte {
	var result bytes.Buffer

	encoder := gob.NewEncoder(&result)
	err := encoder.Encode(o)
	if err != nil {
		clog.Fatal(util.CLOG_SKIP_DISPLAY_INFO, err.Error())
	}

	return result.Bytes()
}

// 反序列化对象
func DeserializeBlock(d []byte) *IPFSObject {
	var o IPFSObject

	decoder := gob.NewDecoder(bytes.NewReader(d))
	err := decoder.Decode(&o)
	if err != nil {
		clog.Fatal(util.CLOG_SKIP_DISPLAY_INFO, err.Error())
	}

	return &o
}

// 检查数据库文件是否存在
func dbExists(dbFile string) bool {
	if _, err := os.Stat(dbFile); os.IsNotExist(err) {
		return false
	}

	return true
}
