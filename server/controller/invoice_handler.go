package controller

import (
	"bufio"
	"fmt"
	"github.com/gin-gonic/gin"
	"graduation-project/server/common/utils"
	bind "graduation-project/server/controller/bind_param"
	"graduation-project/server/models"
	"io"
	"net/http"
	"strconv"
	"strings"
	"time"
)

var invoice = new(models.Invoice)

// 删除
func DeleteInvoice(c *gin.Context) {
	id := c.Param("invoiceId")
	invoiceId, _ := strconv.Atoi(id)

	err := invoice.DeleteInvoice(invoiceId)
	utils.AssertErr(err)

	c.JSON(http.StatusOK, gin.H{
		"succ": true,
	})
}

// 文件导入
func PostUploadInvoiceFile(c *gin.Context) {
	var invoiceList []models.Invoice

	file, _ := c.FormFile("file")

	open, _ := file.Open()
	reader := bufio.NewReader(open)

	// 将数据解析到切片
	for {
		line, err := reader.ReadString('\n')
		line = strings.TrimSpace(line)
		if (err == io.EOF || err != nil) && len(line) == 0 {
			fmt.Println("文件读读取 完毕/失败......")
			break
		}

		var inv models.Invoice
		split := strings.Split(line, ",")
		inv.InvoiceId,err = strconv.Atoi(strings.TrimSpace(split[0]))
		if err != nil {	// 如果是第一行，跳过
			continue
		}

		inv.SaleId,_ = strconv.Atoi(strings.TrimSpace(split[1]))
		inv.PurchaseId,_ = strconv.Atoi(strings.TrimSpace(split[2]))

		inv.AmountMoney,_ = strconv.ParseFloat(strings.TrimSpace(split[3]), 64)
		inv.TaxMoney,_ = strconv.ParseFloat(strings.TrimSpace(split[4]), 64)
		inv.TotalMoney,_ = strconv.ParseFloat(strings.TrimSpace(split[5]), 64)

		inv.BillingTime,_ = time.Parse("2006-1-2", strings.TrimSpace(split[6]))

		inv.VoidMark = strings.TrimSpace(split[7])

		invoiceList = append(invoiceList, inv)
	}

	err := invoice.BatchImport(invoiceList)
	utils.AssertErr(err)

	c.JSON(http.StatusOK, gin.H{
		"succ": true,
	})
}

// 查询
func PostQueryInvoiceToPage(c *gin.Context) {
	var invoiceParam bind.InvoiceParam
	c.Bind(&invoiceParam)

	current := 1
	size := 10

	if invoiceParam.Current != 0 {
		current = invoiceParam.Current
	}
	if invoiceParam.Size != 0 {
		size = invoiceParam.Size
	}

	atoi, _ := strconv.Atoi(invoiceParam.InvoiceId)
	invoiceList, total := invoice.QueryInvoiceToPage(atoi, current, size)
	pages := total / size
	if total%size != 0 {
		pages++
	}

	var page = utils.Page{
		Records: invoiceList,
		Current: current,
		Pages:   pages,
		Size:    size,
		Total:   total,
	}

	c.JSON(http.StatusOK, gin.H{
		"succ": true,
		"page": page,
	})
}
