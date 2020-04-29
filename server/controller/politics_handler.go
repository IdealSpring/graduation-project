package controller

import (
	"github.com/gin-gonic/gin"
	"graduation-project/server/common/utils"
	bind "graduation-project/server/controller/bind_param"
	"graduation-project/server/models"
	"net/http"
	"strconv"
)

var politics = new(models.Politics)

// 通知用户
func PostPoliticsNotify(c *gin.Context) {
	var politicsParam bind.PoliticsParam
	c.Bind(&politicsParam)

	politicsList, err := politics.QueryPoliticsNotify(politicsParam.UserId, politicsParam.PoliticsType)
	utils.AssertErr(err)

	c.JSON(http.StatusOK, gin.H{
		"succ":         true,
		"politicsList": politicsList,
	})
}

// 删除
func DeletePolitics(c *gin.Context) {
	id := c.Param("politicsId")
	politicsId, _ := strconv.Atoi(id)

	err := politics.DeletePolitics(politicsId)
	utils.AssertErr(err)

	c.JSON(http.StatusOK, gin.H{
		"succ": true,
	})
}

// 添加
func PostAddPolitics(c *gin.Context) {
	var politicsParam bind.PoliticsParam
	c.Bind(&politicsParam)
	retPolitics, err := politics.AddPolitics(politicsParam.Title, politicsParam.Abstract, politicsParam.Content, politicsParam.PoliticsType)
	utils.AssertErr(err)

	c.JSON(http.StatusOK, gin.H{
		"succ":       true,
		"politicsId": retPolitics.PoliticsId,
		"createTime": retPolitics.CreateTime,
	})
}

// 查询
func PostQueryPoliticsToPage(c *gin.Context) {
	var politicsParam bind.PoliticsParam
	c.Bind(&politicsParam)

	title := ""
	current := 1
	size := 10

	if politicsParam.Title != "" {
		title = politicsParam.Title
	}
	if politicsParam.Current != 0 {
		current = politicsParam.Current
	}
	if politicsParam.Size != 0 {
		size = politicsParam.Size
	}

	roleList, total := politics.QueryPoliticsToPage(politicsParam.PoliticsType, title, current, size)
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
