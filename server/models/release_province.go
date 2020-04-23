package models

import "time"

// 发行省份
type ReleaseProvince struct {
	ProvinceId   int       `gorm:"column:id" json:"provinceId"`
	ProvinceName string    `gorm:"column:province_name" json:"provinceName"`
	Ban          int       `gorm:"column:ban" json:"ban"`
	CreateTime   time.Time `gorm:"column:create_time" json:"createTime"`
	UpdateTime   time.Time `gorm:"column:update_time" json:"updateTime"`
}

func (rp *ReleaseProvince) TableName() string { return "release_province" }
