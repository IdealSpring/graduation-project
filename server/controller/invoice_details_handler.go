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
)

var invoiceDetails = new(models.InvoiceDetails)


// 删除
func DeleteInvoiceDetails(c *gin.Context) {
	idd := c.Param("id")
	id, _ := strconv.Atoi(idd)

	err := invoiceDetails.DeleteInvoiceDetails(id)
	utils.AssertErr(err)

	c.JSON(http.StatusOK, gin.H{
		"succ": true,
	})
}

// 文件导入
func PostUploadInvoiceDetailsFile(c *gin.Context) {
	var invoiceDetailsList []models.InvoiceDetails

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

		var det models.InvoiceDetails
		split := strings.Split(line, ",")
		det.InvoiceId,err = strconv.Atoi(strings.TrimSpace(split[0]))
		if err != nil {	// 如果是第一行，跳过
			continue
		}

		det.GoodsName = strings.TrimSpace(split[1])
		det.Specifications = strings.TrimSpace(split[2])
		det.Unit = strings.TrimSpace(split[3])

		det.Number,_ = strconv.ParseFloat(strings.TrimSpace(split[4]), 64)
		det.Price,_ = strconv.ParseFloat(strings.TrimSpace(split[5]), 64)
		det.AmountMoney,_ = strconv.ParseFloat(strings.TrimSpace(split[6]), 64)
		det.TaxAmount,_ = strconv.ParseFloat(strings.TrimSpace(split[7]), 64)
		det.ProductCode = strings.TrimSpace(split[7])

		invoiceDetailsList = append(invoiceDetailsList, det)
	}

	err := invoiceDetails.BatchImport(invoiceDetailsList)
	utils.AssertErr(err)

	c.JSON(http.StatusOK, gin.H{
		"succ": true,
	})
}

// 查询
func PostQueryInvoiceDetailsToPage(c *gin.Context) {
	var detailsParam bind.InvoiceDetailsParam
	c.Bind(&detailsParam)

	current := 1
	size := 10

	if detailsParam.Current != 0 {
		current = detailsParam.Current
	}
	if detailsParam.Size != 0 {
		size = detailsParam.Size
	}

	atoi, _ := strconv.Atoi(detailsParam.DetailsId)
	invoiceDetailsList, total := invoiceDetails.QueryInvoiceDetailsToPage(atoi, current, size)
	pages := total / size
	if total%size != 0 {
		pages++
	}

	var page = utils.Page{
		Records: invoiceDetailsList,
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
