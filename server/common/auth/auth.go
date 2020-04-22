package auth

import (
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"graduation-project/server/common/logger"
	"time"
)

var secret = "Token Secret" // 加密盐值
var issuer = "IdealSpring"  // 发布者

// 自定义 payload 结构体
type UserStdClaims struct {
	jwt.StandardClaims
	UserId   int
	Username string
	RoleId   int
}

func (c *UserStdClaims) Valid() (err error) {
	if c.VerifyExpiresAt(time.Now().Unix(), true) == false {
		//utils.AssertErr(errors.New("token is expired"), 4433)
		return errors.New("CustomError#4433#token is expired")
	}
	if !c.VerifyIssuer(issuer, true) {
		return errors.New("token's issuer is wrong")
	}
	return
}

// 创建Token
func JwtCreateToken(userId int, username string, roleId int) (string, error) {
	expireTime := time.Now().Add(time.Hour)
	stdClaims := jwt.StandardClaims{
		ExpiresAt: expireTime.Unix(),
		IssuedAt:  time.Now().Unix(),
		Issuer:    issuer,
	}
	uClaims := &UserStdClaims{
		StandardClaims: stdClaims,
		UserId:         userId,
		Username:       username,
		RoleId:         roleId,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, uClaims)
	tokenStr, err := token.SignedString([]byte(secret))
	if err != nil {
		logger.Logln("SignedString 错误")
	}
	return tokenStr, err
}

//JwtParseUser 解析payload的内容,得到用户信息
//gin-middleware 会使用这个方法
func JwtParseForToken(tokenString string) (*UserStdClaims, error) {
	if tokenString == "" {
		return nil, errors.New("no token is found in Authorization Bearer")
	}
	claims := UserStdClaims{}
	_, err := jwt.ParseWithClaims(tokenString, &claims, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v\n", token.Header["alg"])
		}
		return []byte(secret), nil
	})
	if err != nil {
		return nil, err
	}
	return &claims, err
}
