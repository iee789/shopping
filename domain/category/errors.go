package category

import (
	"errors"
)

// 商品分类自定义错误
var (
	ErrCategoryExistWithName = errors.New("商品分类已经存在")
)
