package app

import "github.com/gin-gonic/gin"

func Run() {
	r := gin.Default()
	r.Run(":8080")
}
