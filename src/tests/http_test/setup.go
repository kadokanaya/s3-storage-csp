package http

import (
	"sp/src/domains/entities"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// テスト用のDBを取得
func LoadTestDB() (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	db.AutoMigrate(&entities.User{})
	db.AutoMigrate(&entities.Receipt{})
	return db, nil
}

// テスト用のParamを取得
func LoadTestParam() (*entities.Param, error) {
	return &entities.Param{}, nil
}
