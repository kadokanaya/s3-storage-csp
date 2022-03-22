package contracts

import (
	"sp/src/domains/entities"
	"sp/src/usecases/port"
)

type Param struct {
	Pairing string
	G      []byte
	U      []byte
}

type ContentContract struct {}

func NewContentContracts() port.ContentContract {
	return &ContentContract{
	}
}

func (cc *ContentContract) Register(content *entities.Content) error {
	return nil
}