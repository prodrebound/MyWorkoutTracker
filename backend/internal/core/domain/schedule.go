package domain

import "gorm.io/gorm"

type Schedule struct {
	gorm.Model
	DayOfWeek int     `json:"day_of_week"`
	RoutineID uint    `json:"routine_id"`
	Routine   Routine `json:"routine"`
}
