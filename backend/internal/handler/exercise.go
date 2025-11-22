package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// ExerciseHandler hält die Abhängigkeiten (hier erstmal nur DB)
type ExerciseHandler struct {
	DB *gorm.DB
}

// NewExerciseHandler ist der Konstruktor
func NewExerciseHandler(db *gorm.DB) *ExerciseHandler {
	return &ExerciseHandler{DB: db}
}

// GetExercises ist die Methode, die wir gleich aufrufen
func (h *ExerciseHandler) GetExercises(c *gin.Context) {
	// Hier kommt später die Logik rein (Service aufrufen)
	c.JSON(http.StatusOK, gin.H{
		"message": "Hier kommen alle Übungen",
	})
}
