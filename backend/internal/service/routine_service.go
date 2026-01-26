package service

import (
	"errors"

	"github.com/prodrebound/MyWorkoutTracker/backend/internal/core/domain"
	"github.com/prodrebound/MyWorkoutTracker/backend/internal/core/ports"
)

type RoutineService struct {
	repo ports.RoutineRepository
}

func NewRoutineService(repo ports.RoutineRepository) *RoutineService {
	return &RoutineService{repo: repo}
}

func (s *RoutineService) CreateRoutine(routine *domain.Routine) error {
	if routine.Name == "" {
		return errors.New("routine name is required")
	}

	if len(routine.Exercises) == 0 {
		return errors.New("a routine must contain at least one exercise")
	}

	return s.repo.Create(routine)
}

func (s *RoutineService) GetAllRoutines() ([]domain.Routine, error) {
	return s.repo.GetAll()
}

func (s *RoutineService) GetRoutineByID(id uint) (*domain.Routine, error) {
	return s.repo.GetByID(id)
}

func (s *RoutineService) UpdateRoutine(id uint, name, description string, exercises []domain.RoutineExercise) (*domain.Routine, error) {
	existingRoutine, err := s.repo.GetByID(id)
	if err != nil {
		return nil, err
	}

	if name == "" {
		return nil, errors.New("routine name cannot be empty")
	}
	if len(exercises) == 0 {
		return nil, errors.New("routine must have at least one exercise")
	}

	updatedRoutine := &domain.Routine{
		ID:          existingRoutine.ID,
		CreatedAt:   existingRoutine.CreatedAt,
		UpdatedAt:   existingRoutine.UpdatedAt,
		Name:        name,
		Description: description,
		Exercises:   exercises,
	}

	err = s.repo.Update(updatedRoutine)
	if err != nil {
		return nil, err
	}

	return s.repo.GetByID(id)
}

func (s *RoutineService) DeleteRoutine(id uint) error {
	return s.repo.Delete(id)
}
