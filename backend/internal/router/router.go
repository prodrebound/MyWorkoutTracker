package router

import (
	"github.com/gin-gonic/gin"
	"github.com/prodrebound/MyWorkoutTracker/backend/internal/handler"
)

func SetupRoutes(r *gin.Engine, exerciseHandler *handler.ExerciseHandler, routineHandler *handler.RoutineHandler) {

	api := r.Group("/api/v1")
	{
		exercises := api.Group("/exercises")
		{
			exercises.GET("", exerciseHandler.GetExercises)

			exercises.POST("", exerciseHandler.CreateExercise)

			exercises.GET("/:id", exerciseHandler.GetExerciseByID)

			exercises.PUT("/:id", exerciseHandler.UpdateExercise)

			exercises.DELETE("/:id", exerciseHandler.DeleteExercise)
		}

		routines := api.Group("/routines")
		{
			routines.POST("", routineHandler.CreateRoutine)
			routines.GET("", routineHandler.GetAllRoutines)
			routines.GET("/:id", routineHandler.GetRoutineByID)
			routines.PUT("/:id", routineHandler.UpdateRoutine)
			routines.DELETE("/:id", routineHandler.DeleteRoutine)
		}
	}
}
