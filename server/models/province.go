package models

import (
	"fmt"
	orm "graduation-project/server/database"
	"time"
)

// 发行省份
type Province struct {
	ProvinceId   int    `gorm:"column:id" json:"provinceId"`
	ProvinceName string `gorm:"column:province_name" json:"provinceName"`
	Ban          int    `gorm:"column:ban" json:"ban"`
	UserCount    int	`gorm:"default:'0'" json:"userCount"`
	CreateTime   time.Time `gorm:"column:create_time" json:"createTime"`
	UpdateTime   time.Time `gorm:"column:update_time" json:"updateTime"`
}

func (p *Province) TableName() string { return "province" }

// 添加发行省份
func (p *Province) AddProvince(name string) (pro Province, err error) {
	pro.ProvinceName = name
	pro.Ban = 0
	pro.CreateTime = time.Now()
	pro.UpdateTime = time.Now()

	err = orm.DB.Create(&pro).Error
	return
}

// 分页查询发行省份
func (p *Province) FindProvincePageByCondition(provinceName string, current int, size int) ([]Province, int) {
	proDB := orm.DB.Model(p).Where("province_name like ?", fmt.Sprintf("%%%s%%", provinceName))
	var total int
	proDB.Count(&total)

	var proList []Province
	proDB.Offset((current - 1) * size).Limit(size).Find(&proList)

	for index, pro := range proList {
		var count int
		orm.DB.Model(&User{}).Where("province_id = ?", pro.ProvinceId).Count(&count)
		proList[index].UserCount = count
		//(pro).UserCount = count
	}

	return proList, total
}

// 获取所有发行省份
func (p *Province) GetAllProvice() (list []Province, err error) {
	err = orm.DB.Find(&list).Error
	return
}
