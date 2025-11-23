package ports

import "github.com/prodrebound/MyWorkoutTracker/backend/internal/core/domain"

type ExerciseRepository interface {
	Create(exercise *domain.Exercise) error
	GetAllExercises() ([]domain.Exercise, error)
	GetExerciseByID(id uint) (*domain.Exercise, error)
	GetExerciseByMuscleGroup(muscleGroup string) ([]domain.Exercise, error)
	UpdateExercise(exercise *domain.Exercise) error
	DeleteExercise(id uint) error
}
