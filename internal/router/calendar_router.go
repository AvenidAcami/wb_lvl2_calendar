package router

import (
	"wb_lvl2_calendar/internal/controller"
	"wb_lvl2_calendar/internal/repository"
	"wb_lvl2_calendar/internal/service"

	"github.com/gin-gonic/gin"
)

func InitCalendarRoutes(r *gin.Engine) {
	calendarRepo := *repository.NewCalendarRepository()
	calendarServ := *service.NewCalendarService(&calendarRepo)
	calendarContr := controller.NewCalendarController(&calendarServ)

	r.POST("/create_event", calendarContr.CreateEvent)
	r.PATCH("/update_event", calendarContr.UpdateEvent)
	r.DELETE("/delete_event", calendarContr.DeleteEvent)
	r.GET("/events_for_day", calendarContr.GetEventsForDay)
	r.GET("/events_for_week", calendarContr.GetEventsForWeek)
	r.GET("/events_for_month", calendarContr.GetEventsForMonth)
}
