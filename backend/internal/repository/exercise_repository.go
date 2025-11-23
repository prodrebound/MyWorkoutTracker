package repository

import (
	"errors"

	"github.com/prodrebound/MyWorkoutTracker/backend/internal/core/domain"
	"gorm.io/gorm"
)

type ExerciseRepository struct {
	db *gorm.DB
}

func NewExerciseRepository(db *gorm.DB) *ExerciseRepository {
	return &ExerciseRepository{db: db}
}

func (r *ExerciseRepository) Create(exercise *domain.Exercise) error {
	return r.db.Create(exercise).Error
}

func (r *ExerciseRepository) GetAllExercises() ([]domain.Exercise, error) {
	var exercises []domain.Exercise
	err := r.db.Find(&exercises).Error
	return exercises, err
}

func (r *ExerciseRepository) GetExerciseByID(id uint) (*domain.Exercise, error) {
	var exercise domain.Exercise
	result := r.db.First(&exercise, id)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, errors.New("exercise not found")
		}
		return nil, result.Error
	}
	return &exercise, nil
}

func (r *ExerciseRepository) GetExerciseByMuscleGroup(muscleGroup string) ([]domain.Exercise, error) {
	var exercises []domain.Exercise
	result := r.db.Where("muscle_group = ?", muscleGroup).Find(&exercises)
	if result.RowsAffected == 0 {
		return nil, errors.New("no exercises found for the specified muscle group")
	}
	return exercises, result.Error
}

func (r *ExerciseRepository) UpdateExercise(exercise *domain.Exercise) error {
	return r.db.Save(exercise).Error
}

func (r *ExerciseRepository) DeleteExercise(id uint) error {
	return r.db.Delete(&domain.Exercise{}, id).Error
}
