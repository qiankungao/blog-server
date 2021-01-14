package model

import (
	"blog-server/global"
	"blog-server/pkg/setting"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"time"
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
	db.Callback().Create().Replace("gorm:update_time_stamp", updateTimeStampForCreateCallback)
	db.Callback().Update().Replace("gorm:update_time_stamp", updateTimeStampForUpdateCallBack)
	db.Callback().Delete().Replace("gorm:delete", deleteCallback)
	db.DB().SetMaxIdleConns(dataBaseSetting.MaxIdleConns)
	db.DB().SetMaxOpenConns(dataBaseSetting.MaxOpenConns)
	if err = db.DB().Ping(); err != nil {
		return nil, err
	}
	return db, nil
}

func updateTimeStampForCreateCallback(scope *gorm.Scope) {
	if !scope.HasError() {
		nowTime := time.Now().Unix()
		if createTimeField, ok := scope.FieldByName("CreatedOn"); ok {
			if createTimeField.IsBlank {
				_ = createTimeField.Set(nowTime)
			}
		}
		if modifyTimeField, ok := scope.FieldByName("ModifiedOn"); ok {
			if modifyTimeField.IsBlank {
				_ = modifyTimeField.Set(nowTime)
			}
		}
	}
}

//更新回调
func updateTimeStampForUpdateCallBack(scope *gorm.Scope) {
	if _, ok := scope.Get("gorm:update_column"); !ok {
		_ = scope.SetColumn("ModifiedOn", time.Now().Unix())
	}
}

//删除回调
func deleteCallback(scope *gorm.Scope) {
	if !scope.HasError() {
		var extraOption string
		if str, ok := scope.Get("gorm:delete_column"); ok {
			extraOption = fmt.Sprintf("%s", str)
		}

		deleteOnField, hasDeleteOnField := scope.FieldByName("DeleteOn")
		isDelField, hasIsDelField := scope.FieldByName("IsDel")
		if !scope.Search.Unscoped && hasDeleteOnField && hasIsDelField {
			nowTime := time.Now().Unix()
			scope.Raw(fmt.Sprintf(
				"UPDATE $v SET %v=%v,%v=%v%v%v",
				scope.QuotedTableName(),
				scope.Quote(deleteOnField.DBName),
				scope.AddToVars(nowTime),
				scope.Quote(isDelField.DBName),
				scope.AddToVars(1),
				addExtraSpaceIfExist(scope.CombinedConditionSql()),
				addExtraSpaceIfExist(extraOption))).Exec()
		} else {
			scope.Raw(fmt.Sprintf(
				"DELETE FROM %V%V%V",
				scope.QuotedTableName(),
				addExtraSpaceIfExist(scope.CombinedConditionSql()),
				addExtraSpaceIfExist(extraOption))).Exec()

		}
	}
}

func addExtraSpaceIfExist(str string) string {
	if str != "" {
		return " " + str
	}
	return ""
}
