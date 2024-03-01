package category

import (
	"gorm.io/gorm"
)

// 商品分类模型
type Category struct {
	gorm.Model
	Name     string `gorm:"unique"` // 商品分类名
	Desc     string // 商品分类描述
	IsActive bool   // 是否被激活
}

// 新建商品分类
func NewCategory(name string, desc string) *Category {
	return &Category{
		Name:     name,
		Desc:     desc,
		IsActive: true,
	}
}
