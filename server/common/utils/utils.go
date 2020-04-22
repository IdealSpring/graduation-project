package utils

import (
	"strconv"
)

// AssertErr 错误断言
// 当 error 不为 nil 时触发 panic
// 对于当前请求不会再执行接下来的代码，并且返回指定格式的错误信息和错误码
// 若 msg 为空，则默认为 error 中的内容
func AssertErr(err error, v ...interface{}) {
	if err != nil {
		statusCode := 200
		msg := err.Error()

		for _, val := range v {
			switch value := val.(type) {
			case int:
				statusCode = value
			case string:
				msg = value
			}
		}

		panic("CustomError#" + strconv.Itoa(statusCode) + "#" + msg)
	}
}
