package models

import (
	"fmt"
	orm "graduation-project/server/database"
	"time"
)

type OperationLog struct {
	LogId         int       `gorm:"column:id" json:"logId"`
	ModuleName    string    `gorm:"column:module_name" json:"moduleName"`
	OperationType string    `gorm:"column:operation_type" json:"operationType"`
	Method        string    `gorm:"column:method" json:"method"`
	IP            string    `gorm:"column:ip" json:"ip"`
	URI           string    `gorm:"column:uri" json:"uri"`
	Status        string    `gorm:"column:status" json:"status"`
	UserId        int       `gorm:"column:user_id" json:"userId"`
	CreateTime    time.Time `gorm:"column:create_time" json:"createTime"`
}
type OperationLogPlus struct {
	*OperationLog
	Username string `gorm:"column:username" json:"username"`
	Nick     string `gorm:"column:nick" json:"nick"`
	Total    int    `gorm:"column:count(*)" json:"total"`
}
// 设置表名
func (l *OperationLog) TableName() string { return "operation_log" }

// 删除所有操作日志
func (l *OperationLog) DeleteAllOperationLog() error {
	return orm.DB.Delete(l).Error
}

// 删除所有登录日志
func (l *OperationLog) DeleteAllLoginLog() error {
	return orm.DB.Where("operation_type = ? or operation_type = ?", "登录", "登出").Delete(l).Error
}

// 删除单条操作日志
func (l *OperationLog) DeleteOperationLog(logId int) error {
	return orm.DB.Where("id = ?", logId).Delete(l).Error
}

// 获取日志 to 列表
func (l *OperationLog) FindLogPageByCondition(isLogin bool, moduleName string, username string, nick string, current int, size int) ([]OperationLogPlus, int) {
	loginStr := ""
	if isLogin {
		loginStr = "登"
	}
	sql := "select l.id,l.module_name,l.operation_type,l.method,l.ip,l.uri,l.status,l.user_id,l.create_time,u.username,u.nick" +
		" from operation_log l join user u on l.user_id = u.id" +
		" where l.operation_type like ? and l.module_name like ?" +
		" and u.username like ? and u.nick like ?"

	logDB := orm.DB.Raw(sql, fmt.Sprintf("%%%s%%", loginStr), fmt.Sprintf("%%%s%%", moduleName),
		fmt.Sprintf("%%%s%%", username), fmt.Sprintf("%%%s%%", nick))
	var list []OperationLogPlus
	logDB.Offset((current - 1) * size).Limit(size).Find(&list)

	var countNew  OperationLogPlus
	sql2 := "select count(*)" +
		" from operation_log l join user u on l.user_id = u.id" +
		" where l.operation_type like ? and l.module_name like ?" +
		" and u.username like ? and u.nick like ?"
	orm.DB.Raw(sql2, fmt.Sprintf("%%%s%%", loginStr), fmt.Sprintf("%%%s%%", moduleName),
		fmt.Sprintf("%%%s%%", username), fmt.Sprintf("%%%s%%", nick)).Find(&countNew)

	return list, countNew.Total
}

// 创建操作日志
func (l *OperationLog) AddOperationLog(ol *OperationLog) error {
	return orm.DB.Create(ol).Error
}
