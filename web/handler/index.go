package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// 获取首页
func GetIndex(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", nil)
}