package controller

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"graduation-project/server/common/utils"
	bind "graduation-project/server/controller/bind_param"
	"graduation-project/server/models"
	"net/http"
	"strconv"
	"time"
)

var u = new(models.User)

// 登录，发放Token
func LoginHandler(c *gin.Context) {
	var user models.User

	err := c.ShouldBind(&user)
	utils.AssertErr(err, -1)

	token, err := user.Login()
	utils.AssertErr(err, -1)

	// 设置 current_user
	c.Set("current_user", models.User{UserId: user.UserId})

	c.JSON(http.StatusOK, gin.H{
		"succ":   true,
		"code":   http.StatusOK,
		"token":  token,
		"expire": time.Now().Add(time.Hour).Format(time.RFC3339),
	})
}

// 修改用户密码
func PutUserPwd(c *gin.Context) {
	data, ok := c.Get("current_user")
	if !ok {
		utils.AssertErr(errors.New("current_user 未获取到"), -1)
	}

	var Pwd = struct {
		Pwd  string `form:"pwd" binding:"required"`
		Pwd2 string `form:"pwd2" binding:"required"`
	}{}

	err := c.Bind(&Pwd)
	utils.AssertErr(err, -1)

	if Pwd.Pwd != Pwd.Pwd2 {
		utils.AssertErr(errors.New("两次密码不相同"), -1)
	}

	user := data.(models.User)

	err = new(models.User).UpdatePassowordById(user.UserId, Pwd.Pwd)
	utils.AssertErr(err, -1)

	c.JSON(http.StatusOK, gin.H{
		"succ": true,
		"code": 1,
		"msg":  "更新成功",
	})
}

// 用户退出
func PostLogout(c *gin.Context) {
	// 使用的是Token作为身份验证，所以后端什么都需要做，前端自己清楚存储的Token就可以了
	c.JSON(http.StatusOK, gin.H{
		"succ": true,
	})
}

// 获取用户信息
func GetUserInfo(c *gin.Context) {
	currentUserData, ok := c.Get("current_user")
	if !ok {
		utils.AssertErr(errors.New("用户无权限"), -1)
	}

	currentUser := currentUserData.(models.User)
	fmt.Println(currentUser)

	user, err := new(models.User).FindUserAllInfoById(currentUser.UserId)
	utils.AssertErr(err, -1)

	c.JSON(http.StatusOK, gin.H{
		"succ":  true,
		"code":  200,
		"role":  user.Role,
		"perms": user.Role.Perms,
		"nick":  user.Nick,
		"name":  user.Username,
		"user":  user,
	})
}

// 获取用户列表[条件可有可无]
func PostQuery(c *gin.Context) {
	var params = bind.PostQueryBindParam{}
	err := c.Bind(&params)
	utils.AssertErr(err)

	nick := ""
	current := 1
	size := 10

	if params.Nick != "" {
		nick = params.Nick
	}
	if params.Current != 0 {
		current = params.Current
	}
	if params.Size != 0 {
		size = params.Size
	}
	userList, total := u.FindUserPageByCondition(nick, current, size)
	pages := total / size
	if total%size != 0 {
		pages++
	}

	var page = utils.Page{
		Records: userList,
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

// 删除用户
func DeleteUserById(c *gin.Context) {
	var user models.User
	uid := c.Param("uid")
	userId, err := strconv.Atoi(uid)

	err = user.DeleteById(userId)
	utils.AssertErr(err, -1)

	c.JSON(http.StatusOK, gin.H{
		"succ": true,
		"msg":  "删除成功",
	})
}

// 添加用户
func PostAddUser(c *gin.Context) {
	var param = bind.PostAddUserBindParam{}
	c.Bind(&param)

	uid, created, err := u.AddUser(param.Username, param.Nick, param.RoleId, param.ProvinceId, param.Pwd)
	utils.AssertErr(err)

	c.JSON(http.StatusOK, gin.H{
		"succ":    true,
		"uid":     uid,
		"created": created,
	})
}

// 更新用户信息
func PutUpdateUser(c *gin.Context) {
	var param = bind.PostAddUserBindParam{}
	c.Bind(&param)

	err := u.UpdateUser(param.UserId, param.Nick, param.Pwd, param.ProvinceId)

	utils.AssertErr(err)

	c.JSON(http.StatusOK, gin.H{
		"succ":       true,
		"updateTime": time.Now(),
	})
}

// 更新用户角色信息
func UpdateRole(c *gin.Context) {
	var userInfo = bind.UserInfoParam{}
	c.Bind(&userInfo)

	err := u.UpdateRole(userInfo.UserId, userInfo.RoleId)
	utils.AssertErr(err)

	c.JSON(http.StatusOK, gin.H{
		"succ": true,
	})
}
