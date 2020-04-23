package controller

import (
	"github.com/gin-gonic/gin"
	"graduation-project/server/common/utils"
	"graduation-project/server/models"
	"net/http"
	"time"
)

// 发行省份列表
func PostQueryProvince(c *gin.Context) {
	var province = models.ReleaseProvince{
		ProvinceId:   1,
		ProvinceName: "吉林省",
		CreateTime:   time.Now(),
		UpdateTime:   time.Now(),
		Ban:          0,
	}
	var province2 = models.ReleaseProvince{
		ProvinceId:   2,
		ProvinceName: "黑龙江省",
		CreateTime:   time.Now(),
		UpdateTime:   time.Now(),
		Ban:          0,
	}
	var province3 = models.ReleaseProvince{
		ProvinceId:   3,
		ProvinceName: "辽宁省",
		CreateTime:   time.Now(),
		UpdateTime:   time.Now(),
		Ban:          0,
	}

	var page = utils.Page{
		Records: [...]models.ReleaseProvince{province, province2, province3},
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
