package model

import (
	"blog-server/global"
	"blog-server/pkg/setting"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type Model struct {
	ID         uint32 `gorm:"primary_key" json:"id"`
	CreatedBy  string `json:"created_by"`
	ModifiedBy string `json:"modified_by"`
	CreatedOn  uint32 `json:"created_on"`
	ModifiedOn uint32 `json:"modified_on"`
	DeletedOn  uint32 `json:"deleted_on"`
	IsDel      uint8  `json:"is_del"`
}

func NewDBEngine(dataBaseSetting *setting.DataBaseSettings) (*gorm.DB, error) {
	db, err := gorm.Open(dataBaseSetting.DBType,
		fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=%s&parseTime=%t&loc=Local",
			dataBaseSetting.UserName,
			dataBaseSetting.Password,
			dataBaseSetting.Host,
			dataBaseSetting.DBName,
			dataBaseSetting.Charset,
			dataBaseSetting.ParseTime))
	if err != nil {
		return nil, err
	}
	if global.ServerSetting.RunMode == "debug" {
		db.LogMode(true)
	}
	db.SingularTable(true)
	db.DB().SetMaxIdleConns(dataBaseSetting.MaxIdleConns)
	db.DB().SetMaxOpenConns(dataBaseSetting.MaxOpenConns)
	if err = db.DB().Ping(); err != nil {
		return nil, err
	}

	return db, nil

}
