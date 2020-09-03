package mysql

import "time"

type AppVersion struct {
	Id          uint32    `xorm:"'id' "`
	ClientType  uint8     `xorm:"'client_type' "`
	BuildCode   string    `xorm:"'build_code' "`
	DownloadUrl string    `xorm:"'download_url' "`
	ForceUpdate uint8     `xorm:"'force_update' "`
	VersionName string    `xorm:"'version_name' "`
	Title       string    `xorm:"'title' "`
	Content     string    `xorm:"'content' "`
	Remark      string    `xorm:"'remark' "`
	Status      int8      `xorm:"'status' "`
	IsDelete    uint8     `xorm:"'is_delete' "`
	CreateTime  time.Time `xorm:"'create_time' "`
	UpdateTime  time.Time `xorm:"'update_time' "`
}
