package models

import (
	orm "graduation-project/server/database"
	"time"
)

// 发行省份
type Province struct {
	ProvinceId   int       `gorm:"column:id" json:"provinceId"`
	ProvinceName string    `gorm:"column:province_name" json:"provinceName"`
	Ban          int       `gorm:"column:ban" json:"ban"`
	CreateTime   time.Time `gorm:"column:create_time" json:"createTime"`
	UpdateTime   time.Time `gorm:"column:update_time" json:"updateTime"`
}

func (p *Province) TableName() string { return "province" }

// 获取所有发行省份
func (p *Province) GetAllProvice() (list []Province, err error) {
	err = orm.DB.Find(&list).Error
	return
}
