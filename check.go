package helper

import (
	"fmt"
	"regexp"
)

// 验证手机号
func CheckPhone(code string) (err error) {
	if len(code) != 11 {
		err = fmt.Errorf("手机号应该为11位")
		return
	}

	reg, err := regexp.Compile(`^(1[3-9]\d{9})$`)
	if err != nil {
		return
	}

	if !reg.Match([]byte(code)) {
		err = fmt.Errorf("手机号格式错误")
		return
	}
	return
}
