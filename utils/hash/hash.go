package hash

import (
	"math/rand"
	"time"

	"golang.org/x/crypto/bcrypt"
)

// 随机字符串生成，包含了大小写字母和数字
const charset = "abcdefghijklmnopqrstuvwxyz" +
	"ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

// 使用当前的纳秒级时间创建seed
var seededRand *rand.Rand = rand.New(
	rand.NewSource(time.Now().UnixNano()))

// 创建salt
func CreateSalt() string {
	b := make([]byte, bcrypt.MaxCost) //最大成本因子，创建一个足够大的字节切片
	for i := range b {
		b[i] = charset[seededRand.Intn(len(charset))] //随机选择哈希字符插入其中最后转换为字符串返回
	}
	return string(b)
}

// 接受一个明文密码作为输入，并使用bcrypt算法对其进行哈希
func HashPassword(password string) (string, error) {
	// 函数接受两个参数：密码的字节切片和成本因子
	// 成本因子决定了哈希的复杂度；bcrypt.DefaultCost 是bcrypt算法的默认成本因子
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}

// 验证一个用户输入的明文密码和一个存储在数据库中已知的哈希密码是否匹配
func CheckPasswordHash(password string, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
