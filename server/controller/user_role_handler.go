package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"graduation-project/server/common/utils"
	bind "graduation-project/server/controller/bind_param"
	"graduation-project/server/models"
	"net/http"
	"strconv"
)

// 工具变量
var r = new(models.Role)

// 修改角色
func PutUpdateRole(c *gin.Context) {
	var roleInfo bind.RoleInfoParam2
	c.Bind(&roleInfo)
	role, err := r.UpdateRole(roleInfo.RoleId, roleInfo.RoleName, roleInfo.Desc)
	utils.AssertErr(err)

	c.JSON(http.StatusOK, gin.H{
		"succ":       true,
		"updateTime": role.UpdateTime,
	})
}

// 删除角色
func DeleteRole(c *gin.Context) {
	var roleInfo bind.RoleInfoParam2
	c.Bind(&roleInfo)

	err := r.DeleteRole(roleInfo.RoleId)
	utils.AssertErr(err)
	c.JSON(http.StatusOK, gin.H{
		"succ": true,
	})
}

// 添加角色
func PostAddRole(c *gin.Context) {
	var roleInfo bind.RoleInfoParam
	c.Bind(&roleInfo)
	role, err := r.AddRole(roleInfo.RoleName, roleInfo.Desc)
	utils.AssertErr(err)

	c.JSON(http.StatusOK, gin.H{
		"succ":       true,
		"roleId":     role.RoleId,
		"createTime": role.CreateTime,
	})
}

// 更新角色权限
func PutUpdatePerm(c *gin.Context) {
	var rolePerms bind.RoleInfoParam
	c.Bind(&rolePerms)

	id, _ := strconv.Atoi(rolePerms.RoleId)
	err := r.UpdatePerms(id, rolePerms.PermType, rolePerms.Values)
	utils.AssertErr(err)

	fmt.Println()
	c.JSON(http.StatusOK, gin.H{
		"succ": true,
	})
}

// 获取角色权限
func GetRolePerms(c *gin.Context) {
	roleId, _ := strconv.Atoi(c.Param("roleId"))

	permValList, role, err := r.GetRolePerms(roleId)
	utils.AssertErr(err, -1)

	c.JSON(http.StatusOK, gin.H{
		"succ":      true,
		"menuPvals": permValList,
		"role":      role,
	})
}

// 查询角色列表[查询条件可选]
func PostRoleQueryToPage(c *gin.Context) {
	var roleInfo bind.RoleInfoParam
	c.Bind(&roleInfo)

	roleName := ""
	current := 1
	size := 10

	if roleInfo.RoleName != "" {
		roleName = roleInfo.RoleName
	}
	if roleInfo.Current != 0 {
		current = roleInfo.Current
	}
	if roleInfo.Size != 0 {
		size = roleInfo.Size
	}

	roleList, total := r.FindRolePageByCondition(roleName, current, size)
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

// 获取用户所有角色
func GetRole(c *gin.Context) {
	roleList, err := r.GetAllRoles()
	utils.AssertErr(err, -1)

	c.JSON(http.StatusOK, gin.H{
		"succ":    true,
		"options": roleList,
	})
}
