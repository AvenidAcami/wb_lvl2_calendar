package controller

import (
	"errors"
	"net/http"
	"wb_lvl2_calendar/internal/model"
	"wb_lvl2_calendar/internal/service"

	"github.com/gin-gonic/gin"
)

type CalendarController struct {
	serv service.ICalendarService
}

func NewCalendarController(serv service.ICalendarService) CalendarController {
	return CalendarController{serv: serv}
}

func (cc *CalendarController) CreateEvent(c *gin.Context) {
	var event model.EventInput
	err := c.ShouldBindJSON(event)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	id, err := cc.serv.CreateEvent(event)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"result": id,
	})
}

func (cc *CalendarController) UpdateEvent(c *gin.Context) {
	var event model.EventInput
	err := c.ShouldBindJSON(event)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	err = cc.serv.UpdateEvent(event)
	if err != nil {
		if errors.Is(err, errors.New("event not found")) {
			c.JSON(http.StatusServiceUnavailable, gin.H{
				"error": err.Error(),
			})
			return
		}
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"result": "ok",
	})
}

func (cc *CalendarController) DeleteEvent(c *gin.Context) {
	id := c.Param("id")
	err := cc.serv.DeleteEvent(id)
	if err != nil {
		if errors.Is(err, errors.New("event not found")) {
			c.JSON(http.StatusServiceUnavailable, gin.H{
				"error": err.Error(),
			})
			return
		}
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"result": "ok",
	})
}

func (cc *CalendarController) GetEventsForDay(c *gin.Context) {
	userId := c.Param("user_id")
	date := c.Param("date")
	events, err := cc.serv.GetEventsForDay(userId, date)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"result": events,
	})
}

func (cc *CalendarController) GetEventsForWeek(c *gin.Context) {
	userId := c.Param("user_id")
	date := c.Param("date")
	events, err := cc.serv.GetEventsForWeek(userId, date)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"result": events,
	})
}

func (cc *CalendarController) GetEventsForMonth(c *gin.Context) {
	userId := c.Param("user_id")
	date := c.Param("date")
	events, err := cc.serv.GetEventsForMonth(userId, date)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"result": events,
	})
}
