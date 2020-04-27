package models

import (
	"fmt"
	orm "graduation-project/server/database"
	"time"
)

// 权限表
type Permission struct {
	PermId     int       `gorm:"column:id" json:"permId"`
	PermName   string    `gorm:"column:perm_name" json:"name"`
	Value      string    `gorm:"column:value" json:"val"`
	PermType   int       `gorm:"column:type" json:"permType"`
	CreateTime time.Time `gorm:"column:create_time" json:"createTime"`
	UpdateTime time.Time `gorm:"column:update_time" json:"updateTime"`
}

// 角色权限关联表
type RolePerm struct {
	RoleId       int `gorm:"column:role_id" json:"roleId"`
	PermissionId int `gorm:"column:permission_id" json:"permissionId"`
}

// 角色表
type Role struct {
	RoleId     int          `gorm:"column:id;primary_key" json:"roleId"`
	RoleName   string       `gorm:"column:role_name" json:"roleName"`
	Desc       string       `gorm:"column:desc" json:"desc"`
	Priority   int          `gorm:"column:priority" json:"priority"`
	Perms      []Permission `gorm:"many2many:role_perm" json:"perms"` // 多对多关系
	CreateTime time.Time    `gorm:"column:create_time" json:"createTime"`
	UpdateTime time.Time    `gorm:"column:update_time" json:"updateTime"`
}

// 设置表名
func (r *Permission) TableName() string { return "permission" }
func (r *RolePerm) TableName() string   { return "role_perm" }
func (r *Role) TableName() string       { return "role" }

// 删除角色
func (r *Role) DeleteRole(roleId int) error {
	return orm.DB.Where("id = ?", roleId).Delete(r).Error
}

// 更新角色
func (r *Role) UpdateRole(roleId int, roleName string, desc string) (role Role, err error) {
	role.RoleId = roleId
	role.RoleName = roleName
	role.Desc = desc
	role.UpdateTime = time.Now()

	err = orm.DB.Model(r).Where("id = ?", roleId).Update(&role).Error
	return
}

// 添加角色
func (r *Role) AddRole(roleName string, desc string) (role Role, err error) {
	role.RoleName = roleName
	role.Desc = desc
	role.Priority = 0
	role.CreateTime = time.Now()
	role.UpdateTime = time.Now()
	err = orm.DB.Create(&role).Error
	return
}

// 更新用户角色
func (r *Role) UpdatePerms(roleId int, permType int, vals []string) error {
	err := orm.DB.Where("role_id = ?", roleId).Delete(&RolePerm{}).Error

	var allPerms []Permission
	var permsMap = make(map[string]int)
	err = orm.DB.Find(&allPerms).Error
	for _, perm := range allPerms {
		permsMap[perm.Value] = perm.PermId
	}

	var newPerms []RolePerm
	for _, perm := range vals {
		permId, ok := permsMap[perm]
		if ok {
			newPerms = append(newPerms, RolePerm{RoleId: roleId, PermissionId:permId})
		}
	}

	for _, rp := range newPerms{
		err = orm.DB.Create(&rp).Error
	}

	return err
}

// 获取角色权限
func (r *Role) GetRolePerms(roleId int) (permValList []string, role Role, err error) {
	role.RoleId = roleId
	err = orm.DB.Model(&role).Take(&role).Error
	err = orm.DB.Model(&role).Related(&role.Perms, "perms").Error

	for _, val := range role.Perms {
		permValList = append(permValList, val.Value)
	}
	return
}

// 查询角色列表
func (r *Role) FindRolePageByCondition(roleName string, current int, size int) ([]Role, int) {
	roleDB := orm.DB.Model(&Role{}).Where("role_name like ?", fmt.Sprintf("%%%s%%", roleName))
	var total int
	roleDB.Count(&total)

	var roleList []Role
	roleDB.Offset((current - 1) * size).Limit(size).Find(&roleList)

	var retRoleList []Role
	for _, role := range retRoleList {
		orm.DB.Model(&role).Related(&role.Perms)
		retRoleList = append(retRoleList, role)
	}

	return roleList, total
}

// 查询所有角色
func (r *Role) GetAllRoles() (roleList []Role, err error) {
	err = orm.DB.Find(&roleList).Error
	return
}
