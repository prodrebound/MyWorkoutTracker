package domain

import (
	"time"

	"gorm.io/gorm"
)

type Exercise struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`

	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`

	Name        string `json:"name"`
	MuscleGroup string `json:"muscle_group"`
	DefaultSets int    `json:"default_sets"`
	DefaultReps int    `json:"default_reps"`
}
