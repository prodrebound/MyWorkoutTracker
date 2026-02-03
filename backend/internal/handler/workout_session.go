package handler

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/prodrebound/MyWorkoutTracker/backend/internal/service"
)

type CreateSessionRequest struct {
	Date      string `json:"date" binding:"required"`
	RoutineID uint   `json:"routine_id" binding:"required"`
}

type WorkoutSessionHandler struct {
	service *service.WorkoutSessionService
}

func NewWorkoutSessionHandler(service *service.WorkoutSessionService) *WorkoutSessionHandler {
	return &WorkoutSessionHandler{service: service}
}

// POST /sessions
func (h *WorkoutSessionHandler) ScheduleSession(c *gin.Context) {
	var req CreateSessionRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	parsedDate, err := time.Parse("2006-01-02", req.Date)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid date format. Use YYYY-MM-DD"})
		return
	}
	session, err := h.service.ScheduleSession(parsedDate, req.RoutineID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, session)
}

// GET /sessions?start_date=2023-10-01&end_date=2023-10-31
func (h *WorkoutSessionHandler) GetHistory(c *gin.Context) {
	// 1. Query Parameter auslesen
	startStr := c.Query("start_date")
	endStr := c.Query("end_date")

	var start, end time.Time
	var err error

	if startStr == "" || endStr == "" {
		end = time.Now().Add(24 * time.Hour)
		start = time.Now().Add(-30 * 24 * time.Hour)
	} else {
		start, err = time.Parse("2006-01-02", startStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid start_date format (YYYY-MM-DD)"})
			return
		}

		end, err = time.Parse("2006-01-02", endStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid end_date format (YYYY-MM-DD)"})
			return
		}
		end = end.Add(24 * time.Hour)
	}

	sessions, err := h.service.GetHistory(start, end)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, sessions)
}
