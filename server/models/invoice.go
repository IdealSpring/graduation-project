package models

import (
	"bytes"
	"fmt"
	"github.com/jinzhu/gorm"
	orm "graduation-project/server/database"
	"time"
)

type Invoice struct {
	InvoiceId   int       `gorm:"column:id" json:"invoiceId"`
	SaleId      int       `gorm:"column:sale_id" json:"saleId"`
	PurchaseId  int       `gorm:"column:purchase_id" json:"purchaseId"`
	AmountMoney float64   `gorm:"column:amount_money" json:"amountMoney"`
	TaxMoney    float64   `gorm:"column:tax_money" json:"taxMoney"`
	TotalMoney  float64   `gorm:"column:total_money" json:"totalMoney"`
	BillingTime time.Time `gorm:"column:billing_time" json:"billingTime"`
	VoidMark    string    `gorm:"column:void_mark" json:"voidMark"`
	CreateTime    time.Time `gorm:"column:create_time" json:"createTime"`
}
func (i *Invoice) TableName() string { return "invoice" }

// 删除
func (i *Invoice) DeleteInvoice(invoiceId int) error {
	return orm.DB.Where("id = ?", invoiceId).Delete(i).Error
}

// 批量导入
func (i *Invoice) BatchImport(invoiceList []Invoice) error {
	// 开启事务
	db := orm.DB
	tx := db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	var buf bytes.Buffer
	sql := "insert into invoice values"
	if _, err := buf.WriteString(sql); err != nil {
		return err
	}
	for i, val := range invoiceList {
		if i == len(invoiceList) - 1 {
			buf.WriteString(fmt.Sprintf("(%d, %d, %d, '%f', '%f', '%f', '%s', '%s', now());",
				val.InvoiceId, val.SaleId, val.PurchaseId, val.AmountMoney, val.TaxMoney, val.TotalMoney,
				val.BillingTime.Format("2006-01-02"),val.VoidMark))
		} else {
			buf.WriteString(fmt.Sprintf("(%d, %d, %d, '%f', '%f', '%f', '%s', '%s', now()),",
				val.InvoiceId, val.SaleId, val.PurchaseId, val.AmountMoney, val.TaxMoney, val.TotalMoney,
				val.BillingTime.Format("2006-01-02"),val.VoidMark,))
		}
	}

	err := tx.Exec(buf.String()).Error
	fmt.Println("err -->", err)
	// 提交
	return tx.Commit().Error
}

// 查询
func (i *Invoice) QueryInvoiceToPage(invoiceId int, current int, size int) ([]Invoice, int) {
	var invoiceDB *gorm.DB
	if invoiceId == 0 {
		invoiceDB = orm.DB.Model(i).Order("create_time desc")
	} else {
		invoiceDB = orm.DB.Model(i).Where("id like ?", fmt.Sprintf("%%%d%%", invoiceId)).Order("create_time desc")
	}
	var total int
	invoiceDB.Count(&total)

	var list []Invoice
	invoiceDB.Offset((current - 1) * size).Limit(size).Find(&list)

	return list, total
}
