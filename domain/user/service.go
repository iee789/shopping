package user

import "shopping/utils/hash"

// 除了再次封装数据库外，主要是为了满足业务逻辑
// 流程逻辑：UI -> controller -> service -> repository

// 用户service结构体
type Service struct {
	r Repository
}

// 实例化service
func NewUserService(r Repository) *Service {
	r.Migration()
	r.InsertSampleData()
	return &Service{
		r: r,
	}
}

// 创建用户
func (c *Service) Create(user *User) error {
	// 先验证用户密码
	if user.Password != user.Password2 {
		return ErrMismatchedPasswords
	}
	// 用户名是否存在
	_, err := c.r.GetByName(user.Username)
	if err == nil {
		return ErrUserExistWithName
	}
	// 无效用户名
	if ValidateUserName(user.Username) {
		return ErrInvalidUsername
	}
	// 无效密码
	if ValidatePassword(user.Password) {
		return ErrInvalidPassword
	}
	// 创建用户
	err = c.r.Create(user)
	return err
}

// 查询用户
func (c *Service) GetUser(username string, password string) (User, error) {
	user, err := c.r.GetByName(username)
	if err != nil {
		return User{}, ErrUserNotFound
	}
	match := hash.CheckPasswordHash(password+user.Salt, user.Password)
	if !match {
		return User{}, ErrUserNotFound
	}
	return user, nil
}

// 更新用户
func (c *Service) UpdateUser(user *User) error {
	return c.r.Update(user)
}
