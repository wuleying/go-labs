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

// 获取IPFS对象
func GetObject(fileHash string) (*IPFSObject, error) {
	db, err := bolt.Open(util.DB_FILE_PATH, util.DB_FILE_MODE, nil)
	if err != nil {
		clog.Fatal(util.CLOG_SKIP_DISPLAY_INFO, err.Error())
	}
	defer db.Close()

	var object *IPFSObject

	err = db.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket(util.Str2bytes(util.BLOCK_BUCKET_NAME))
		object = DeserializeBlock(bucket.Get(util.Str2bytes(fileHash)))

		return nil
	})

	if err != nil {
		return nil, err
	}

	return object, nil
}

// 保存数据
func AddObject(filePath string) (string, error) {
	fileHash, err := commands.AddFile(filePath)
	if err != nil {
		return "", err
	}

	fileName := util.FileGetName(filePath)
	fileSize := util.FileGetSize(filePath)

	object := IPFSObject{util.Str2bytes(fileHash), util.Str2bytes(fileName), fileSize, time.Now().Unix()}
	db, err := bolt.Open(util.DB_FILE_PATH, util.DB_FILE_MODE, nil)
	if err != nil {
		clog.Fatal(util.CLOG_SKIP_DISPLAY_INFO, err.Error())
	}
	defer db.Close()

	err = db.Update(func(tx *bolt.Tx) error {
		bucket, err := tx.CreateBucketIfNotExists(util.Str2bytes(util.BLOCK_BUCKET_NAME))
		if err != nil {
			clog.Fatal(util.CLOG_SKIP_DISPLAY_INFO, err.Error())
		}

		err = bucket.Put(object.FileHash, object.Serialize())
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
func (object *IPFSObject) Serialize() []byte {
	var result bytes.Buffer

	encoder := gob.NewEncoder(&result)
	err := encoder.Encode(object)
	if err != nil {
		clog.Fatal(util.CLOG_SKIP_DISPLAY_INFO, err.Error())
	}

	return result.Bytes()
}

// 反序列化对象
func DeserializeBlock(data []byte) *IPFSObject {
	var object IPFSObject

	decoder := gob.NewDecoder(bytes.NewReader(data))
	err := decoder.Decode(&object)
	if err != nil {
		clog.Fatal(util.CLOG_SKIP_DISPLAY_INFO, err.Error())
	}

	return &object
}

// 检查数据库文件是否存在
func dbExists(dbFile string) bool {
	if _, err := os.Stat(dbFile); os.IsNotExist(err) {
		return false
	}

	return true
}
