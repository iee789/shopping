package database_handler

import (
	"fmt"
	"time"

	"gorm.io/gorm/schema"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func NewMySQLDB(conString string) *gorm.DB {
	db, err := gorm.Open(mysql.Open(conString), &gorm.Config{
		PrepareStmt: true, // 启用预处理允许数据库重用相同sql语句
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true, // 表名单数形式
			NoLowerCase:   true, // 表名、列名不会转换为小写
		},
		NowFunc: func() time.Time {
			return time.Now().UTC()
		},
	})
	if err != nil {
		panic(fmt.Sprintf("不能连接到数据库 : %s", err.Error()))
	}

	return db
}
