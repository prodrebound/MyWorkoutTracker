package domain

import (
	"time"

	"gorm.io/gorm"
)

type WorkoutSession struct {
	ID        uint           `gorm:"primaryKey" json:"id"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`

	Date      time.Time `json:"date"`
	Completed bool      `json:"completed"`
	Duration  int       `json:"duration"` // in Minutes

	RoutineID uint     `json:"routine_id"`
	Routine   *Routine `json:"routine,omitempty"`

	Exercises []SessionExercise `json:"exercises"`
}

type SessionExercise struct {
	ID uint `gorm:"primaryKey" json:"id"`

	WorkoutSessionID uint `json:"-"`

	ExerciseID uint      `json:"exercise_id"`
	Exercise   *Exercise `json:"exercise,omitempty"`

	Sets   int     `json:"sets"`
	Reps   int     `json:"reps"`
	Weight float64 `json:"weight"`
	Order  int     `json:"order"`
}
