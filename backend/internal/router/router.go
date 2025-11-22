package router

import (
	"github.com/gin-gonic/gin"
	"github.com/prodrebound/MyWorkoutTracker/backend/internal/handler"
)

// SetupRoutes nimmt die Engine UND die Handler entgegen
func SetupRoutes(r *gin.Engine, exerciseHandler *handler.ExerciseHandler) {

	// API Versionierung ist immer gut
	api := r.Group("/api/v1")
	{
		// Exercise Gruppe
		exercises := api.Group("/exercises")
		{
			// GET /api/v1/exercises
			exercises.GET("", exerciseHandler.GetExercises)
		}

		// Hier kommen später weitere Gruppen (z.B. /workouts)
	}
}
