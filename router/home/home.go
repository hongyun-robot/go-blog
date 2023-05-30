package home

import "github.com/gin-gonic/gin"

func InitHome(r *gin.Engine) {
	r.Group("/home")
}