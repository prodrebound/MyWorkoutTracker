package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/prodrebound/MyWorkoutTracker/backend/internal/service"
)

type CreateExerciseRequest struct {
	Name        string `json:"name" binding:"required"`
	MuscleGroup string `json:"muscle_group"`
	DefaultSets int    `json:"default_sets"`
	DefaultReps int    `json:"default_reps"`
}

type ExerciseHandler struct {
	service *service.ExerciseService
}

func NewExerciseHandler(service *service.ExerciseService) *ExerciseHandler {
	return &ExerciseHandler{service: service}
}

func (h *ExerciseHandler) CreateExercise(c *gin.Context) {
	var req CreateExerciseRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	createdExercise, err := h.service.CreateExercise(req.Name, req.MuscleGroup, req.DefaultSets, req.DefaultReps)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create exercise"})
		return
	}

	c.JSON(http.StatusCreated, createdExercise)
}

func (h *ExerciseHandler) GetExercises(c *gin.Context) {
	muscleGroup := c.Query("muscle_group")

	if muscleGroup != "" {
		exercises, err := h.service.GetExerciseByMuscleGroup(muscleGroup)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, exercises)
	} else {
		exercises, err := h.service.GetAllExercises()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve exercises"})
			return
		}
		c.JSON(http.StatusOK, exercises)
	}
}

func (h *ExerciseHandler) GetExerciseByID(c *gin.Context) {
	id, err := h.parseID(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	exercise, err := h.service.GetExerciseByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Exercise not found"})
		return
	}

	c.JSON(http.StatusOK, exercise)
}

func (h *ExerciseHandler) UpdateExercise(c *gin.Context) {
	id, err := h.parseID(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	var req CreateExerciseRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updatedExercise, err := h.service.UpdateExercise(id, req.Name, req.MuscleGroup, req.DefaultSets, req.DefaultReps)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update: " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, updatedExercise)
}

func (h *ExerciseHandler) DeleteExercise(c *gin.Context) {
	id, err := h.parseID(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	err = h.service.DeleteExercise(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete exercise"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Exercise deleted successfully"})
}

func (h *ExerciseHandler) parseID(c *gin.Context) (uint, error) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		return 0, err
	}
	return uint(id), nil
}
