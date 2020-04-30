package models

import (
	"bytes"
	"fmt"
	"github.com/jinzhu/gorm"
	orm "graduation-project/server/database"
	"time"
)

type InvoiceDetails struct {
	Id             int       `gorm:"column:id" json:"id"`
	InvoiceId      int       `gorm:"column:invoice_id" json:"invoiceId"`
	GoodsName      string    `gorm:"column:goods_name" json:"goodsName"`
	Specifications string    `gorm:"column:specifications" json:"specifications"`
	Unit           string    `gorm:"column:unit" json:"unit"`
	Number         float64   `gorm:"column:number" json:"number"`
	Price          float64   `gorm:"column:price" json:"price"`
	AmountMoney    float64   `gorm:"column:amount_money" json:"amountMoney"`
	TaxAmount      float64   `gorm:"column:tax_amount" json:"taxAmount"`
	ProductCode    string    `gorm:"column:product_code" json:"productCode"`
	CreateTime     time.Time `gorm:"column:create_time" json:"createTime"`
}

func (d *InvoiceDetails) TableName() string { return "details" }


// 删除
func (d *InvoiceDetails) DeleteInvoiceDetails(id int) error {
	return orm.DB.Where("id = ?", id).Delete(d).Error
}

// 批量导入
func (d *InvoiceDetails) BatchImport(detailsList []InvoiceDetails) error {
	// 开启事务
	db := orm.DB
	tx := db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	var buf bytes.Buffer
	sql := "insert into details values"
	if _, err := buf.WriteString(sql); err != nil {
		return err
	}
	for i, val := range detailsList {
		if i == len(detailsList)-1 {
			buf.WriteString(fmt.Sprintf("(null,%d,'%s','%s','%s','%f','%f','%f','%f','%s',now());",
				val.InvoiceId, val.GoodsName, val.Specifications, val.Unit, val.Number, val.Price, val.AmountMoney, val.TaxAmount, val.ProductCode))
		} else {
			buf.WriteString(fmt.Sprintf("(null,%d,'%s','%s','%s','%f','%f','%f','%f','%s',now()),",
				val.InvoiceId, val.GoodsName, val.Specifications, val.Unit, val.Number, val.Price, val.AmountMoney, val.TaxAmount, val.ProductCode))
		}
	}

	err := tx.Exec(buf.String()).Error
	fmt.Println("err -->", err)
	// 提交
	return tx.Commit().Error
}

// 查询
func (d *InvoiceDetails) QueryInvoiceDetailsToPage(detailsId int, current int, size int) ([]InvoiceDetails, int) {
	var detailsDB *gorm.DB
	if detailsId == 0 {
		detailsDB = orm.DB.Model(d).Order("create_time desc")
	} else {
		detailsDB = orm.DB.Model(d).Where("id like ?", fmt.Sprintf("%%%d%%", detailsId)).Order("create_time desc")
	}
	var total int
	detailsDB.Count(&total)

	var list []InvoiceDetails
	detailsDB.Offset((current - 1) * size).Limit(size).Find(&list)

	return list, total
}
