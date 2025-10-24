package app

import (
	"log"
	"os"
	"wb_lvl2_calendar/config"
	"wb_lvl2_calendar/internal/router"

	"github.com/gin-gonic/gin"
)

func Run() {
	r := gin.Default()
	router.InitCalendarRoutes(r)
	config.InitENV()
	port, err := config.GetEnv("PORT")
	if err != nil {
		log.Fatal("env error")
		os.Exit(1)
	}
	r.Run(":" + port)
}
