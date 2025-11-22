package domain

import "gorm.io/gorm"

type Routine struct {
	gorm.Model
	Name        string            `json:"name"`
	Description string            `json:"description"`
	Exercises   []RoutineExercise `json:"exercises"`
}

type RoutineExercise struct {
	gorm.Model
	RoutineID  uint     `json:"routine_id"`
	ExerciseID uint     `json:"exercise_id"`
	Exercise   Exercise `json:"exercise"`
	Order      int      `json:"order"`
	Sets       int      `json:"sets"`
	Reps       int      `json:"reps"`
}
