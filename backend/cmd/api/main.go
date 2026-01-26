package main

import (
	"log"

	"github.com/prodrebound/MyWorkoutTracker/backend/internal/config"
	"github.com/prodrebound/MyWorkoutTracker/backend/internal/core/domain"
	"github.com/prodrebound/MyWorkoutTracker/backend/internal/handler"
	"github.com/prodrebound/MyWorkoutTracker/backend/internal/repository"
	"github.com/prodrebound/MyWorkoutTracker/backend/internal/router"
	"github.com/prodrebound/MyWorkoutTracker/backend/internal/service"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	cfg := config.LoadConfig()

	db, err := gorm.Open(sqlite.Open(cfg.DBName), &gorm.Config{})
	if err != nil {
		log.Fatal("DB Error: ", err)
		return
	}

	err = db.AutoMigrate(&domain.Exercise{}, &domain.Routine{}, &domain.RoutineExercise{})
	if err != nil {
		log.Fatal("Migration Error:", err)
	}

	exerciseRepo := repository.NewExerciseRepository(db)
	exerciseService := service.NewExerciseService(exerciseRepo)
	exerciseHandler := handler.NewExerciseHandler(exerciseService)

	routineRepo := repository.NewRoutineRepository(db)
	routineService := service.NewRoutineService(routineRepo)
	routineHandler := handler.NewRoutineHandler(routineService)

	if cfg.Env == "production" {
		gin.SetMode(gin.ReleaseMode)
	}

	r := gin.Default()
	router.SetupRoutes(r, exerciseHandler, routineHandler)

	log.Printf("Server runs in %s mode on port %s", cfg.Env, cfg.ServerPort)

	if err := r.Run(cfg.ServerPort); err != nil {
		log.Fatal("Could not start server:", err)
	}
}
