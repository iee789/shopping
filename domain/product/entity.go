package product

import (
	"shopping/domain/category"

	"gorm.io/gorm"
)

// 商品模型
type Product struct {
	gorm.Model
	Name       string            // 商品名
	SKU        string            // 商品的唯一描述
	Desc       string            // 商品描述
	StockCount int               // 商品库存
	Price      float32           // 商品价格
	CategoryID uint              // 分类id
	Category   category.Category `json:"-"` // 分类
	IsDeleted  bool              // 软删除字段
}

// 商品结构体实例
func NewProduct(name string, desc string, stockCount int, price float32, cid uint) *Product {
	return &Product{
		Name:       name,
		Desc:       desc,
		StockCount: stockCount,
		Price:      price,
		CategoryID: cid,
		IsDeleted:  false,
	}
}
