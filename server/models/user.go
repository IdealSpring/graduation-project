package models

import (
	"errors"
	"fmt"
	"graduation-project/server/common/auth"
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
	Value      string       `gorm:"column:value" json:"val"`
	Perms      []Permission `gorm:"many2many:role_perm" json:"perms"` // 多对多关系
	CreateTime time.Time    `gorm:"column:create_time" json:"createTime"`
	UpdateTime time.Time    `gorm:"column:update_time" json:"updateTime"`
}

// 用户表
type User struct {
	UserId   int    `gorm:"column:id" json:"userId"`
	Username string `gorm:"column:username" json:"username"`
	Nick     string `gorm:"column:nick" json:"nick"`
	Password string `gorm:"column:password" json:"password"`
	RoleId   int    `gorm:"column:role_id" json:"roleId"`

	Role Role `gorm:"foreignkey:RoleId"` //指定关联外键

	CreateTime time.Time `gorm:"column:create_time" json:"createTime"`
	UpdateTime time.Time `gorm:"column:update_time" json:"updateTime"`
}

// 设置表名
func (r *Permission) TableName() string {
	return "permission"
}
func (r *RolePerm) TableName() string {
	return "role_perm"
}
func (r *Role) TableName() string {
	return "role"
}
func (u *User) TableName() string {
	return "user"
}

// 添加用户
func (u *User) AddUser(username string, nick string, roleID int, pwd string) (int, time.Time, error) {
	user := User{
		Username:   username,
		Nick:       nick,
		RoleId:     roleID,
		Password:   pwd,
		CreateTime: time.Now(),
		UpdateTime: time.Now(),
	}

	err := orm.DB.Create(&user).Error
	id := user.UserId
	created := user.CreateTime

	fmt.Println("id->", id)
	fmt.Println("created->", created)

	return id, created, err
}

func (u *User) UpdateUser(userId int, nick string, pwd string) error {
	return orm.DB.Model(User{}).Where("id = ?", userId).Update(User{Nick:nick, Password:pwd}).Error
}

// 删除用户
func (u *User) DeleteById(id int) error {
	return orm.DB.Where("id = ?", id).Delete(&User{}).Error
}

// 根据条件查询页面信息
func (u *User) FindUserPageByCondition(nick string, current int, size int) ([]User, int) {
	//userDB := orm.DB.Model(&User{}).Where(&User{Nick: nick})
	userDB := orm.DB.Model(&User{}).Where("nick like ?", fmt.Sprintf("%%%s%%", nick))
	var total int
	userDB.Count(&total)

	var userList []User
	userDB.Offset((current - 1) * size).Limit(size).Find(&userList)

	var retUserList []User
	for _, user := range userList {
		orm.DB.Model(&user).Related(&user.Role)
		retUserList = append(retUserList, user)
	}

	return retUserList, total
}

// 查询所有角色
func (r *Role) GetAllRoles() (roleList []Role, err error) {
	err = orm.DB.Find(&roleList).Error
	return
}

// 用户关联角色查询
func (u *User) FindUserAllInfoById(id int) (user User, err error) {
	err = orm.DB.Where("id = ?", id).Take(&user).Error
	if err != nil {
		return user, err
	}

	err = orm.DB.Model(&user).Related(&user.Role).Error
	if err != nil {
		return user, err
	}

	err = orm.DB.Model(&user.Role).Related(&user.Role.Perms, "perms").Error
	if err != nil {
		return user, err
	}

	return
}

// 登录，发放Token
func (u *User) Login() (token string, err error) {
	err = orm.DB.Where("username = ? and password = ?", u.Username, u.Password).Take(u).Error

	if err != nil {
		return "", errors.New("用户名或者密码错误")
	}

	// 清空密码
	u.Password = ""

	token, err = auth.JwtCreateToken(u.UserId, u.Username, u.RoleId)
	return
}

// 根据用户名寻找用户
func (u *User) FindUserByUsername(username string) error {
	err := orm.DB.Where("username = ?", username).Take(u).Error
	return err
}

// 根据id更新用户密码
func (u *User) UpdatePassowordById(id int, pwd string) error {
	u2 := User{Password: pwd, UpdateTime: time.Now()}

	return orm.DB.Model(User{}).Where("id = ?", id).Updates(u2).Error
}
