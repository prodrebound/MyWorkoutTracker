package service

import (
	"errors"

	"github.com/prodrebound/MyWorkoutTracker/backend/internal/core/domain"
	"github.com/prodrebound/MyWorkoutTracker/backend/internal/core/ports"
)

type ExerciseService struct {
	repo ports.ExerciseRepository
}

func NewExerciseService(repo ports.ExerciseRepository) *ExerciseService {
	return &ExerciseService{
		repo: repo,
	}
}

func (s *ExerciseService) CreateExercise(name, muscleGroup string, sets, reps int) (*domain.Exercise, error) {
	if name == "" || muscleGroup == "" {
		return nil, errors.New("exercise name or muscle group cannot be empty")
	}

	if sets < 1 || reps < 1 {
		return nil, errors.New("sets and reps must be greater than 0")
	}

	newExercise := &domain.Exercise{Name: name, MuscleGroup: muscleGroup, DefaultSets: sets, DefaultReps: reps}

	err := s.repo.Create(newExercise)
	if err != nil {
		return nil, err
	}

	return newExercise, nil
}

func (s *ExerciseService) GetAllExercises() ([]domain.Exercise, error) {
	return s.repo.GetAllExercises()
}

func (s *ExerciseService) GetExerciseByID(id uint) (*domain.Exercise, error) {
	return s.repo.GetExerciseByID(id)
}

func (s *ExerciseService) GetExerciseByMuscleGroup(group string) ([]domain.Exercise, error) {
	if group == "" {
		return nil, errors.New("muscle group cannot be empty")
	}
	return s.repo.GetExerciseByMuscleGroup(group)
}

func (s *ExerciseService) UpdateExercise(id uint, name, muscleGroup string, sets, reps int) (*domain.Exercise, error) {
	existingExercise, err := s.repo.GetExerciseByID(id)
	if err != nil {
		return nil, err
	}

	if name != "" {
		existingExercise.Name = name
	}
	existingExercise.MuscleGroup = muscleGroup
	existingExercise.DefaultSets = sets
	existingExercise.DefaultReps = reps

	err = s.repo.UpdateExercise(existingExercise)
	if err != nil {
		return nil, err
	}

	return existingExercise, nil
}

func (s *ExerciseService) DeleteExercise(id uint) error {
	return s.repo.DeleteExercise(id)
}
