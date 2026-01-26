package domain

import (
	"time"

	"gorm.io/gorm"
)

type Routine struct {
	ID        uint           `gorm:"primaryKey" json:"id"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`

	Name        string            `json:"name"`
	Description string            `json:"description"`
	Exercises   []RoutineExercise `json:"exercises"`
}

type RoutineExercise struct {
	ID        uint `gorm:"primaryKey" json:"-"`
	RoutineID uint `json:"-"`

	ExerciseID uint      `json:"exercise_id"`
	Exercise   *Exercise `json:"exercise,omitempty"`

	Sets  int `json:"sets"`
	Reps  int `json:"reps"`
	Order int `json:"order"`
}
