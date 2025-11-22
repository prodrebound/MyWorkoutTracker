package main

import (
	"log"

	"github.com/prodrebound/MyWorkoutTracker/backend/internal/config"
	"github.com/prodrebound/MyWorkoutTracker/backend/internal/core/domain"
	"github.com/prodrebound/MyWorkoutTracker/backend/internal/handler"
	"github.com/prodrebound/MyWorkoutTracker/backend/internal/router"

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

	db.AutoMigrate(
		&domain.Exercise{},
		&domain.Routine{},
		&domain.RoutineExercise{},
		&domain.Schedule{},
	)

	exerciseHandler := handler.NewExerciseHandler(db)

	if cfg.Env == "development" {
		gin.SetMode(gin.ReleaseMode)
	}

	r := gin.Default()
	router.SetupRoutes(r, exerciseHandler)

	log.Printf("Server runs in %s mode on port %s", cfg.Env, cfg.ServerPort)

	// Wir nutzen hier fmt.Sprintf, um den Port korrekt zu übergeben, falls nötig
	if err := r.Run(cfg.ServerPort); err != nil {
		log.Fatal("Could not start server:", err)
	}
}
