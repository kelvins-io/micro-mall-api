package mysql_callback

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"time"
)

// UpdateTimeStampForCreateCallback sets `CreateTime`, `ModifyTime` when creating.
func UpdateTimeStampForCreateCallback(scope *gorm.Scope) {
	if !scope.HasError() {
		now := time.Now().Unix()

		if createTimeField, ok := scope.FieldByName("CreateTime"); ok {
			if createTimeField.IsBlank {
				createTimeField.Set(now)
			}
		}

		if modifyTimeField, ok := scope.FieldByName("ModifyTime"); ok {
			if modifyTimeField.IsBlank {
				modifyTimeField.Set(now)
			}
		}
	}
}

// UpdateTimeStampForUpdateCallback sets `ModifyTime` when updating.
func UpdateTimeStampForUpdateCallback(scope *gorm.Scope) {
	if _, ok := scope.Get("gorm:update_column"); !ok {
		now := time.Now().Unix()
		scope.SetColumn("ModifyTime", now)
	}
}

// DeleteCallback used to delete data from database or set DeletedTime to
// current time and is_del = 1(when using with soft delete).
func DeleteCallback(scope *gorm.Scope) {
	if !scope.HasError() {
		var extraOption string
		if str, ok := scope.Get("gorm:delete_option"); ok {
			extraOption = fmt.Sprint(str)
		}

		deletedTimeField, hasDeletedAtField := scope.FieldByName("DeletedTime")
		isDelField, hasIsDelField := scope.FieldByName("IsDel")

		if !scope.Search.Unscoped && hasDeletedAtField && hasIsDelField {
			now := time.Now().Unix()
			scope.Raw(fmt.Sprintf(
				"UPDATE %v SET %v=%v,%v=%v%v%v",
				scope.QuotedTableName(),
				scope.Quote(deletedTimeField.DBName),
				scope.AddToVars(now),
				scope.Quote(isDelField.DBName),
				scope.AddToVars(1),
				addExtraSpaceIfExist(scope.CombinedConditionSql()),
				addExtraSpaceIfExist(extraOption),
			)).Exec()
		} else {
			scope.Raw(fmt.Sprintf(
				"DELETE FROM %v%v%v",
				scope.QuotedTableName(),
				addExtraSpaceIfExist(scope.CombinedConditionSql()),
				addExtraSpaceIfExist(extraOption),
			)).Exec()
		}
	}
}

// addExtraSpaceIfExist used to add extra space if exists str.
func addExtraSpaceIfExist(str string) string {
	if str != "" {
		return " " + str
	}
	return ""
}