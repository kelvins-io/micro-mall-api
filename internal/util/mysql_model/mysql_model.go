package mysql_model

// ColumnCreateModifyDeleteTime is a basic model for mysql.
type ColumnCreateModifyDeleteTime struct {
	ID          int64 `gorm:"primary_key;AUTO_INCREMENT" json:"id" db:"id"`
	CreateTime  int64 `json:"create_time" db:"create_time"`
	ModifyTime  int64 `json:"modify_time" db:"modify_time"`
	DeletedTime int64 `json:"deleted_time" db:"deleted_time"`
	IsDel       int32 `json:"is_del" db:"is_del"`
}
