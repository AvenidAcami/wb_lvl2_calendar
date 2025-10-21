package app

import (
	"wb_lvl2_calendar/internal/router"

	"github.com/gin-gonic/gin"
)

func Run() {
	r := gin.Default()
	router.InitCalendarRoutes(r)
	r.Run(":8080")
}
