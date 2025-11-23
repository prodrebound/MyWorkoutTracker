package router

import (
	"github.com/gin-gonic/gin"
	"github.com/prodrebound/MyWorkoutTracker/backend/internal/handler"
)

func SetupRoutes(r *gin.Engine, exerciseHandler *handler.ExerciseHandler) {

	api := r.Group("/api/v1")
	{
		exercises := api.Group("/exercises")
		{
			// GET /api/v1/exercises (Hole alle)
			// GET /api/v1/exercises?muscle_group=Chest (Hole nur Brust)
			exercises.GET("", exerciseHandler.GetExercises)

			// POST /api/v1/exercises (Erstellen)
			exercises.POST("", exerciseHandler.CreateExercise)

			// GET /api/v1/exercises/1 (Hole ID 1)
			exercises.GET("/:id", exerciseHandler.GetExerciseByID)

			// PUT /api/v1/exercises/1 (Update ID 1)
			exercises.PUT("/:id", exerciseHandler.UpdateExercise)

			// DELETE /api/v1/exercises/1 (Lösche ID 1)
			exercises.DELETE("/:id", exerciseHandler.DeleteExercise)
		}
	}
}
