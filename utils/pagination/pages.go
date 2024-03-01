package pagination

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// 封装分页工具类

// TODO: refactor methods
var (
	// 默认页数
	DefaultPageSize = 100
	// 最大页数
	MaxPageSize = 1000
	// 查询参数名称
	PageVar = "page"
	// 页数查询参数名称
	PageSizeVar = "pageSize"
)

// 分页结构体
type Pages struct {
	// 当前页码：用户请求的哪一页数据
	Page       int         `json:"page"`
	// 每页显示的记录数：指定每页应该显示多少条记录
	PageSize   int         `json:"pageSize"`
	// 总页数：根据PageSize和TotalCount计算的总页数，知道用户可以浏览多少页
	PageCount  int         `json:"pageCount"`
	// 总记录数：符合条件的总记录数，计算PageCount和确定是否有更多的数据可以加载
	TotalCount int         `json:"totalCount"`
	// 当前页的记录列表
	Items      interface{} `json:"items"`
}

// 实例化分页结构体
func New(page, pageSize, total int) *Pages {
	if pageSize <= 0 {
		pageSize = DefaultPageSize
	}
	if pageSize > MaxPageSize {
		pageSize = MaxPageSize
	}
	pageCount := -1
	if total >= 0 {
		pageCount = (total + pageSize - 1) / pageSize
		if page > pageCount {
			page = pageCount
		}
	}
	if page <= 0 {
		page = 1
	}

	return &Pages{
		Page:       page,
		PageSize:   pageSize,
		TotalCount: total,
		PageCount:  pageCount,
	}
}

// 根据http请求创建分页结构体
func NewFromRequest(req *http.Request, count int) *Pages {
	page := ParseInt(req.URL.Query().Get(PageVar), 1)
	pageSize := ParseInt(req.URL.Query().Get(PageSizeVar), DefaultPageSize)
	return New(page, pageSize, count)
}

// 根据gin请求创建分页结构体
func NewFromGinRequest(g *gin.Context, count int) *Pages {
	page := ParseInt(g.Query(PageVar), 1)
	pageSize := ParseInt(g.Query(PageSizeVar), DefaultPageSize)
	return New(page, pageSize, count)
}

// 解析字符串为整数
func ParseInt(value string, defaultValue int) int {
	if value == "" {
		return defaultValue
	}
	if result, err := strconv.Atoi(value); err == nil {
		return result
	}
	return defaultValue
}

// 分页的偏移量
func (p *Pages) Offset() int {
	return (p.Page - 1) * p.PageSize
}

// 每页显示的记录数
func (p *Pages) Limit() int {
	return p.PageSize
}
