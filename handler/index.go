package handler

import (
	"github.com/gin-gonic/gin"
)

func Index(ctx *gin.Context) (interface{}, error) {
	return "================   Welcome to qa Index Page!    https://github.com/Hui4401/github.com/Hui4401/qa   ================", nil
}
