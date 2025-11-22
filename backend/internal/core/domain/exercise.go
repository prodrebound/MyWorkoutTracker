package domain

import "gorm.io/gorm"

type Exercise struct {
	gorm.Model
	Name        string `json:"name"`
	MuscleGroup string `json:"muscle_group"`
	DefaultSets int    `json:"default_sets"`
	DefaultReps int    `json:"default_reps"`
}
