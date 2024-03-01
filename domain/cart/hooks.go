package cart

import (
	"gorm.io/gorm"
)

// 如果计数为零，则删除商品
func (item *Item) AfterUpdate(tx *gorm.DB) (err error) {
	if item.Count <= 0 {
		//库存数量为0或负数会执行非关联删除（不跟踪软删除）
		return tx.Unscoped().Delete(&item).Error
	}
	return
}
