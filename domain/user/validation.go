package user

import "regexp"

// 用户名密码验证工具类

// 用户名正则表达式，字母开头，最少 7-29 个字母、数字或者下划线
var usernameRegex = regexp.MustCompile("^[A-Za-z][A-Za-z0-9_]{7,29}$")

// 密码正则表达式，字母开头，最少 7-29 个字母、数字或者下划线，字符串至少包含一个数字
var passwordRegex = regexp.MustCompile(`^[A-Za-z][A-Za-z0-9_]{7,29}$`)

func ValidateUserName(name string) bool {
	return usernameRegex.MatchString(name)
}

func ValidatePassword(password string) bool {
	return passwordRegex.MatchString(password)

}
