package api_helper

import (
	"shopping/utils/pagination"

	"github.com/gin-gonic/gin"
)

var userIdText = "userId"

// 从gin请求上下文中提取用户ID将其转换为uint类型
func GetUserId(g *gin.Context) uint {
	return uint(pagination.ParseInt(g.GetString(userIdText), -1))
}
