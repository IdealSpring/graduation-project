package test

import (
	"fmt"
	"graduation-project/server/models"
	"testing"
)

func TestFindUserPageByCondition(t *testing.T) {
	var user = &models.User{}
	user.FindUserAllInfoById(1)
	fmt.Println(user.FindUserPageByCondition("", 0, 10))
}
