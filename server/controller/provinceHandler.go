package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"graduation-project/server/common/utils"
	"graduation-project/server/models"
	"net/http"
	"time"
)

var p = new(models.Province)

// 获取所有发行省份
func GetAllProvice(c *gin.Context) {
	proviceList, err := p.GetAllProvice()
	utils.AssertErr(err, -1)

	fmt.Println("--------------------------------", proviceList)

	c.JSON(http.StatusOK, gin.H{
		"succ":    true,
		"options": proviceList,
	})
}

// 发行省份列表
func PostQueryProvince(c *gin.Context) {
	var province = models.Province{
		ProvinceId:   1,
		ProvinceName: "吉林省",
		CreateTime:   time.Now(),
		UpdateTime:   time.Now(),
		Ban:          0,
	}
	var province2 = models.Province{
		ProvinceId:   2,
		ProvinceName: "黑龙江省",
		CreateTime:   time.Now(),
		UpdateTime:   time.Now(),
		Ban:          0,
	}
	var province3 = models.Province{
		ProvinceId:   3,
		ProvinceName: "辽宁省",
		CreateTime:   time.Now(),
		UpdateTime:   time.Now(),
		Ban:          0,
	}

	var page = utils.Page{
		Records: [...]models.Province{province, province2, province3},
		Current: 1,
		Pages:   1,
		Size:    10,
		Total:   1,
	}

	c.JSON(http.StatusOK, gin.H{
		"succ": true,
		"page": page,
	})
}
