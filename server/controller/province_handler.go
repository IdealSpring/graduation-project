package controller

import (
	"github.com/gin-gonic/gin"
	"graduation-project/server/common/utils"
	bind "graduation-project/server/controller/bind_param"
	"graduation-project/server/models"
	"net/http"
)

var p = new(models.Province)

// 添加发行省份
func PostAddProvince(c *gin.Context) {
	var provincePara bind.ProvinceParam
	c.Bind(&provincePara)

	province, err := p.AddProvince(provincePara.ProvinceName)
	utils.AssertErr(err)

	c.JSON(http.StatusOK, gin.H{
		"succ":       true,
		"provinceId": province.ProvinceId,
		"createTime": province.CreateTime,
	})
}

// 获取所有发行省份
func GetAllProvice(c *gin.Context) {
	proviceList, err := p.GetAllProvice()
	utils.AssertErr(err, -1)

	c.JSON(http.StatusOK, gin.H{
		"succ":    true,
		"options": proviceList,
	})
}

// 发行省份列表
func PostQueryProvince(c *gin.Context) {
	var province bind.ProvinceParam
	c.Bind(&province)

	provinceName := ""
	current := 1
	size := 10

	if province.ProvinceName != "" {
		provinceName = province.ProvinceName
	}
	if province.Current != 0 {
		current = province.Current
	}
	if province.Size != 0 {
		size = province.Size
	}

	roleList, total := p.FindProvincePageByCondition(provinceName, current, size)
	pages := total / size
	if total%size != 0 {
		pages++
	}

	var page = utils.Page{
		Records: roleList,
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
