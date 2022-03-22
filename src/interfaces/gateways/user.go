package gateways

import (
	"sp/src/domains/entities"
	"sp/src/usecases/port"

	"golang.org/x/crypto/bcrypt"

	"gorm.io/gorm"
)

type userRepository struct {
	Conn *gorm.DB
}

func NewUserRepository(conn *gorm.DB) port.UserRepository {
	return &userRepository{
		Conn: conn,
	}
}

func (ur *userRepository) FindByID(id string) (*entities.User, error) {
	var user = &entities.User{}
	user.ID = id
	err := ur.Conn.First(&user).Error
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (ur *userRepository) FindByAddress(address string) (*entities.User, error) {
	var user = &entities.User{}
	err := ur.Conn.Where("address = ?", address).First(&user).Error
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (ur *userRepository) Create(u *entities.User) (*entities.User, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(u.Password), 12)
	if err != nil {
		return nil, err
	}
	u.Password = string(hash)
	err = ur.Conn.Create(u).Error
	if err != nil {
		return nil, err
	}
	return u, nil
}

func (ur *userRepository) Update(u *entities.User) (user *entities.User, err error) {
	err = ur.Conn.Save(u).Error
	if err != nil {
		return nil, err
	}
	return user, nil
}
