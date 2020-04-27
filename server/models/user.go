package models

import (
	"errors"
	"fmt"
	"graduation-project/server/common/auth"
	orm "graduation-project/server/database"
	"time"
)

// 用户表
type User struct {
	UserId   int    `gorm:"column:id" json:"userId"`
	Username string `gorm:"column:username" json:"username"`
	Nick     string `gorm:"column:nick" json:"nick"`
	Password string `gorm:"column:password" json:"password"`

	RoleId int  `gorm:"column:role_id" json:"roleId"`
	Role   Role `gorm:"foreignkey:RoleId"` //指定关联外键

	ProvinceId int      `gorm:"column:province_id" json:"provinceId"`
	Province   Province `gorm:"foreignkey:ProvinceId"` //指定关联外键

	CreateTime time.Time `gorm:"column:create_time" json:"createTime"`
	UpdateTime time.Time `gorm:"column:update_time" json:"updateTime"`
}

// 设置表名
func (u *User) TableName() string {
	return "user"
}

func (u *User) UpdateRole(userId int, roleId int) error {
	return orm.DB.Model(User{}).Where("id = ?", userId).Update(User{RoleId: roleId}).Error
}

// 添加用户
func (u *User) AddUser(username string, nick string, roleID int, provinceId int, pwd string) (int, time.Time, error) {
	user := User{
		Username:   username,
		Nick:       nick,
		RoleId:     roleID,
		ProvinceId: provinceId,
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

// 更新用户 昵称 和 密码
func (u *User) UpdateUser(userId int, nick string, pwd string, provinceId int) error {
	return orm.DB.Model(User{}).Where("id = ?", userId).Update(User{Nick: nick, Password: pwd, ProvinceId: provinceId}).Error
}

// 删除用户
func (u *User) DeleteById(id int) error {
	return orm.DB.Where("id = ?", id).Delete(&User{}).Error
}

// 根据条件查询页面信息
func (u *User) FindUserPageByCondition(nick string, current int, size int) ([]User, int) {
	userDB := orm.DB.Model(&User{}).Where("nick like ?", fmt.Sprintf("%%%s%%", nick))
	var total int
	userDB.Count(&total)

	var userList []User
	userDB.Offset((current - 1) * size).Limit(size).Find(&userList)

	var retUserList []User
	for _, user := range userList {
		orm.DB.Model(&user).Related(&user.Role)
		orm.DB.Model(&user).Related(&user.Province)
		retUserList = append(retUserList, user)
	}

	return retUserList, total
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
