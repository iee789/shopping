package product

import (
	"errors"
)

// 自定义商品错误
var (
	ErrProductNotFound         = errors.New("商品没有找到")
	ErrProductStockIsNotEnough = errors.New("商品库存不足")
)
