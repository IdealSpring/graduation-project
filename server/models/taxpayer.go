package models

import (
	"bytes"
	"fmt"
	"github.com/jinzhu/gorm"
	orm "graduation-project/server/database"
	"time"
)

type Taxpayer struct {
	TaxpayerId       int       `gorm:"column:id" json:"taxpayerId"`
	IndustryCode     int       `gorm:"column:industry_code" json:"industryCode"`
	RegistrationType int       `gorm:"column:registration_type" json:"registrationType"`
	OpenTime         time.Time `gorm:"column:open_time" json:"openTime"`
	CreateTime       time.Time `gorm:"column:create_time" json:"createTime"`
	UpdateTime       time.Time `gorm:"column:update_time" json:"updateTime"`
}

func (t *Taxpayer) TableName() string { return "taxpayer" }

// 批量导入
func (t *Taxpayer) BatchImport(taxpayerList []Taxpayer) error {
	// 开启事务
	db := orm.DB
	tx := db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	var buf bytes.Buffer
	sql := "insert into taxpayer values"
	if _, err := buf.WriteString(sql); err != nil {
		return err
	}
	for i, val := range taxpayerList {
		if i == len(taxpayerList) - 1 {
			buf.WriteString(fmt.Sprintf("(%d, %d, %d, '%s', now(), now());",
				val.TaxpayerId, val.IndustryCode, val.RegistrationType, val.OpenTime.Format("2006-01-02 15:04:05")))
		} else {
			buf.WriteString(fmt.Sprintf("(%d, %d, %d, '%s', now(), now()),",
				val.TaxpayerId, val.IndustryCode, val.RegistrationType, val.OpenTime.Format("2006-01-02 15:04:05")))
		}
	}

	err := tx.Exec(buf.String()).Error
	fmt.Println("err -->", err)
	// 提交
	return tx.Commit().Error
}

// 删除
func (t *Taxpayer) DeleteTaxpayer(taxpayerId int) error {
	return orm.DB.Where("id = ?", taxpayerId).Delete(t).Error
}

// 查询
func (t *Taxpayer) QueryTaxpayerToPage(taxpayerId int, current int, size int) ([]Taxpayer, int) {
	var taxpayerDB *gorm.DB
	if taxpayerId == 0 {
		taxpayerDB = orm.DB.Model(t).Order("create_time desc")
	} else {
		taxpayerDB = orm.DB.Model(t).Where("id like ?", fmt.Sprintf("%%%d%%", taxpayerId)).Order("create_time desc")
	}
	var total int
	taxpayerDB.Count(&total)

	var list []Taxpayer
	taxpayerDB.Offset((current - 1) * size).Limit(size).Find(&list)

	return list, total
}
