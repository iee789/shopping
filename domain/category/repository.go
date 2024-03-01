package category

import (
	"log"

	"gorm.io/gorm"
)

// 商品分类的数据库操作封装

type Repository struct {
	db *gorm.DB
}

// 创建商品分类
func NewCategoryRepository(db *gorm.DB) *Repository {
	return &Repository{
		db: db,
	}
}

// 生成商品分类表
func (r *Repository) Migration() {
	err := r.db.AutoMigrate(&Category{})
	if err != nil {
		log.Print(err)
	}
}

// 生成商品分类测试数据
func (r *Repository) InsertSampleData() {
	categories := []Category{
		{Name: "CAT1", Desc: "Category 1"},
		{Name: "CAT2", Desc: "Category 2"},
	}

	for _, c := range categories {
		// 如果数据库中已经存在一个具有相同Name的记录，那么只有Name字段会被更新
		r.db.Where(Category{Name: c.Name}).Attrs(Category{Name: c.Name}).FirstOrCreate(&c)
	}
}

// 创建商品分类
func (r *Repository) Create(c *Category) error {
	result := r.db.Create(c)

	if result.Error != nil {
		return result.Error
	}

	return nil
}

// 通过名称查询商品分类
func (r *Repository) GetByName(name string) []Category {
	var categories []Category
	r.db.Where("Name = ?", name).Find(&categories)

	return categories
}

// 批量创建商品分类
func (r *Repository) BulkCreate(categories []*Category) (int, error) {
	var count int64
	err := r.db.Create(&categories).Count(&count).Error
	return int(count), err
}

// 提供一个分页获取数据的接口，使得上层逻辑可以根据需要获取特定页的数据，同时知道数据库中总共有多少记录
func (r *Repository) GetAll(pageIndex, pageSize int) ([]Category, int) {
	var categories []Category
	var count int64
	// 1.计算偏移量：pageIndex从1开始，所以需要减去1；pageSize是每页显示的记录数
	// 2.设置每页显示的记录数
	// 3.执行查询，结果存储在categories切片中
	r.db.Offset((pageIndex - 1) * pageSize).Limit(pageSize).Find(&categories).Count(&count)

	return categories, int(count)
}
