package user

import (
	"shopping/utils/hash"

	"gorm.io/gorm"
)

// User模型数据保存到数据库之前自动执行的钩子函数
func (u *User) BeforeSave(tx *gorm.DB) (err error) {
	// 为空说明是新用户或者盐值需要重新生成
	if u.Salt == "" {
		salt := hash.CreateSalt()                                 // 为salt创建一个随机字符串用于给密码结合
		hashPassword, err := hash.HashPassword(u.Password + salt) // 创建hash加密密码
		if err != nil {
			return nil
		}
		u.Password = hashPassword //更新密码与盐值
		u.Salt = salt
	}
	// 确保了用户密码在保存到数据库之前是经过哈希处理的，并且每次保存时都会检查盐值是否需要更新
	return
}
