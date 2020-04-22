package test

import (
	"errors"
	"graduation-project/server/common/utils"
	"testing"
)

func TestAssertErr(t *testing.T) {
	err := errors.New("test err")
	utils.AssertErr(err)
	utils.AssertErr(err, -1)
	utils.AssertErr(err, -1, "这里是错误消息")
}
