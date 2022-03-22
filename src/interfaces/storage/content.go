package storage

import (
	"path/filepath"
	"sp/src/domains/entities"
	"sp/src/drivers/s3"
	"sp/src/usecases/port"

	"bytes"
	"encoding/binary"
	"fmt"
	"os"
)

type ContentStorage struct {
}

func NewContentStorage() port.ContentStorage {
	return &ContentStorage{}
}
func (pr *ContentStorage) Create(c *entities.Content) (*entities.Content, error) {
	storagePath := "storage/" + c.Id + "/content"
	WriteBinaryFile(storagePath, binary.BigEndian, c.Content)
	for i:=0; i<c.SplitCount; i++ {
		storagePath := "storage/" + c.Id + fmt.Sprint(i)
		WriteBinaryFile(storagePath, binary.BigEndian, c.MetaData[i])
	}
	return c, nil
}

func (pr *ContentStorage) Get(id string) (*entities.Content, error) {
	return &entities.Content{
		Content:     []byte{},
		MetaData:    [][]byte{},
		HashedData:  [][]byte{},
		ContentName: "",
		SplitCount:  0,
		Owner:       "",
		Id:          id,
		UserId:      id,
		ContentId:   id,
	}, nil
}

func (pr *ContentStorage) GetPreSignedURL(key string) (string, error) {
	return s3.GetPreSignedURL(key)
}

/*
バイナリデータに変換してファイルに出力する
*/
func WriteBinaryFile(filename string, order binary.ByteOrder, val interface{}) {
	// バイナリデータの格納用
	buf := new(bytes.Buffer)
	// valの値をバイナリデータに変換してbufに格納する
	err := binary.Write(buf, order, val)
	if err != nil {
		fmt.Println("err:", err)
		return
	}
	if err = os.MkdirAll(filepath.Dir(filename), 0770); err != nil {
		return
	}
	// ファイル作成
	file, err2 := os.Create(filename)
	if err2 != nil {
		fmt.Println("file create err:", err2)
		return
	}
	// バイナリデータをファイルに書き込み
	_, err3 := file.Write(buf.Bytes())
	if err3 != nil {
		fmt.Println("file write err:", err3)
		return
	}
	fmt.Println("file write ok.")
}
