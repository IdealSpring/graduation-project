package middleware

import (
	"errors"
	"github.com/gin-gonic/gin"
	"graduation-project/server/common/auth"
	"graduation-project/server/common/utils"
	"graduation-project/server/models"
)

func AuthMiddleware(c *gin.Context) {
	token := c.GetHeader("Admin-Token")

	if token == "" {
		utils.AssertErr(errors.New("从 Header 中获取 token 为空，您无权访问"), -1)
	}

	userStdClaims, err := auth.JwtParseForToken(token)
	utils.AssertErr(err, 500)

	var user = new(models.User)
	err = user.FindUserByUsername(userStdClaims.Username)
	utils.AssertErr(err, -1, "无权限，请重新登录")

	if userStdClaims.UserId != user.UserId || userStdClaims.RoleId != user.RoleId {
		utils.AssertErr(errors.New("无权限，请重新登录"), -1)
	}

	// 存储用户信息
	c.Set("current_user", *user)

	// 放行
	c.Next()
}
