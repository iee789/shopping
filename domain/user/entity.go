package user

import (
	"gorm.io/gorm"
)

// 用户模型
type User struct {
	gorm.Model
	Username  string `gorm:"type:varchar(30)"`  //用户名
	Password  string `gorm:"type:varchar(100)"` //密码
	Password2 string `gorm:"-"`                 //临时存储密码
	Salt      string `gorm:"type:varchar(100)"` //密码盐值
	Token     string `gorm:"type:varchar(500)"` //jwt认证令牌
	IsDeleted bool   //软删除标识
	IsAdmin   bool   //该用户是否具有管理员权限
}

// 新建用户实例
func NewUser(username, password, password2 string) *User {
	return &User{
		Username:  username,
		Password:  password,
		Password2: password2,
		IsDeleted: false,
		IsAdmin:   false,
	}
}
