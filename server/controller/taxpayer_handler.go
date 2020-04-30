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

var taxpayer = new(models.Taxpayer)

// 删除
func DeleteTaxpayer(c *gin.Context) {
	id := c.Param("taxpayerId")
	taxpayerId, _ := strconv.Atoi(id)

	err := taxpayer.DeleteTaxpayer(taxpayerId)
	utils.AssertErr(err)

	c.JSON(http.StatusOK, gin.H{
		"succ": true,
	})
}

// 查询
func PostUploadFile(c *gin.Context) {
	var taxpayerList []models.Taxpayer

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

		var t models.Taxpayer
		split := strings.Split(line, ",")
		t.TaxpayerId,err = strconv.Atoi(strings.TrimSpace(split[0]))
		if err != nil {	// 如果是第一行，跳过
			continue
		}

		t.IndustryCode,_ = strconv.Atoi(strings.TrimSpace(split[1]))
		t.RegistrationType,_ = strconv.Atoi(strings.TrimSpace(split[2]))
		t.OpenTime,_ = time.Parse("2006/1/2 3:04", strings.TrimSpace(split[3]))

		taxpayerList = append(taxpayerList, t)
	}

	err := taxpayer.BatchImport(taxpayerList)
	utils.AssertErr(err)

	c.JSON(http.StatusOK, gin.H{
		"succ": true,
	})
}

// 查询
func PostQueryTaxpayerToPage(c *gin.Context) {
	var taxpayerParam bind.TaxpayerParam
	c.Bind(&taxpayerParam)

	current := 1
	size := 10

	if taxpayerParam.Current != 0 {
		current = taxpayerParam.Current
	}
	if taxpayerParam.Size != 0 {
		size = taxpayerParam.Size
	}

	atoi, _ := strconv.Atoi(taxpayerParam.TaxpayerId)
	taxpayerList, total := taxpayer.QueryTaxpayerToPage(atoi, current, size)
	pages := total / size
	if total%size != 0 {
		pages++
	}

	var page = utils.Page{
		Records: taxpayerList,
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
