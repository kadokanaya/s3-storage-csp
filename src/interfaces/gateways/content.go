package gateways

import (
	"sp/src/domains/entities"
	"sp/src/usecases/port"

	"gorm.io/gorm"
)

type ContentSQLHandler interface {
	Find(interface{}, ...interface{}) (*entities.Receipt, error)
	First(interface{}, ...interface{}) (*entities.Receipt, error)
	Create(interface{}) (*entities.Receipt, error)
	Save(interface{}) error
	Delete(interface{}) *entities.Content
	Where(interface{}, ...interface{}) *entities.Content
}

type ContentRepository struct {
	Conn *gorm.DB
	ContentSQLHandler
}

func NewContentRepository(conn *gorm.DB) port.ContentRepository {
	return &ContentRepository{
		Conn: conn,
	}
}

func (ur *ContentRepository) Find(id string) (*entities.Receipt, error) {
	var receipt = &entities.Receipt{}
	receipt.Id = id
	err := ur.Conn.First(&receipt).Error
	if err != nil {
		return nil, err
	}
	return receipt, nil
}

func (ur *ContentRepository) Create(c *entities.Content) (receipt *entities.Receipt, err error) {
	receipt = &entities.Receipt{
		Id:           c.Id,
		UserId:       c.UserId,
		ContentLogId: c.Id,
		ContentURL:   "localhost:4001/api/content/" + c.Id,
		FileName:     c.ContentName,
	}
	err = ur.Conn.Create(receipt).Error
	if err != nil {
		return nil, err
	}
	return receipt, nil
}
