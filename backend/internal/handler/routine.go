package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/prodrebound/MyWorkoutTracker/backend/internal/core/domain"
	"github.com/prodrebound/MyWorkoutTracker/backend/internal/service"
)

// --- Request DTOs (Data Transfer Objects) ---
// Diese Structs definieren, wie das JSON aussieht, das der User schickt.

type RoutineExerciseRequest struct {
	ExerciseID uint `json:"exercise_id" binding:"required"`
	Sets       int  `json:"sets" binding:"required"`
	Reps       int  `json:"reps" binding:"required"`
	Order      int  `json:"order"`
}

type CreateRoutineRequest struct {
	Name        string                   `json:"name" binding:"required"`
	Description string                   `json:"description"`
	Exercises   []RoutineExerciseRequest `json:"exercises" binding:"required,dive"`
}

type RoutineHandler struct {
	service *service.RoutineService
}

func NewRoutineHandler(service *service.RoutineService) *RoutineHandler {
	return &RoutineHandler{service: service}
}

func (h *RoutineHandler) CreateRoutine(c *gin.Context) {
	var req CreateRoutineRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var domainExercises []domain.RoutineExercise

	for i, exReq := range req.Exercises {
		order := exReq.Order
		if order == 0 {
			order = i + 1
		}

		domainExercises = append(domainExercises, domain.RoutineExercise{
			ExerciseID: exReq.ExerciseID,
			Sets:       exReq.Sets,
			Reps:       exReq.Reps,
			Order:      order,
		})
	}

	newRoutine := &domain.Routine{
		Name:        req.Name,
		Description: req.Description,
		Exercises:   domainExercises,
	}

	err := h.service.CreateRoutine(newRoutine)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, newRoutine)
}

func (h *RoutineHandler) GetAllRoutines(c *gin.Context) {
	routines, err := h.service.GetAllRoutines()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch routines"})
		return
	}
	c.JSON(http.StatusOK, routines)
}

func (h *RoutineHandler) GetRoutineByID(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	routine, err := h.service.GetRoutineByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Routine not found"})
		return
	}

	c.JSON(http.StatusOK, routine)
}

func (h *RoutineHandler) UpdateRoutine(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	var req CreateRoutineRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var domainExercises []domain.RoutineExercise
	for i, exReq := range req.Exercises {
		order := exReq.Order
		if order == 0 {
			order = i + 1
		}
		domainExercises = append(domainExercises, domain.RoutineExercise{
			ExerciseID: exReq.ExerciseID,
			Sets:       exReq.Sets,
			Reps:       exReq.Reps,
			Order:      order,
		})
	}

	updatedRoutine, err := h.service.UpdateRoutine(uint(id), req.Name, req.Description, domainExercises)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, updatedRoutine)
}

func (h *RoutineHandler) DeleteRoutine(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
	}
	err = h.service.DeleteRoutine(uint(id))
}
