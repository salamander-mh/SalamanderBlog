package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/salamander-mh/SalamanderBlog/models"
	"strconv"
)

func GetCategories(c *gin.Context) {
	page, _ := strconv.Atoi(c.Query("page"))
	if page <= 0 {
		page = 1
	}
	categories, count, err := models.GetCategories(page)
	if err != nil {
		respond(c, DB_ERR_CODE, "数据库错误："+err.Error(), nil)
	} else {
		respond(c, 0, "success", gin.H{
			"categories": categories,
			"total":      count,
		})
	}
}

func GetTags(c *gin.Context) {
	page, _ := strconv.Atoi(c.Query("page"))
	if page <= 0 {
		page = 1
	}
	tags, count, err := models.GetTags(page)
	if err != nil {
		respond(c, DB_ERR_CODE, "数据库错误："+err.Error(), nil)
	} else {
		respond(c, 0, "success", gin.H{
			"tags":  tags,
			"total": count,
		})
	}
}
