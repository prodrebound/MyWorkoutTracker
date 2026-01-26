package ports

import "github.com/prodrebound/MyWorkoutTracker/backend/internal/core/domain"

type RoutineRepository interface {
	Create(routine *domain.Routine) error
	GetAll() ([]domain.Routine, error)
	GetByID(id uint) (*domain.Routine, error)
	Update(routine *domain.Routine) error
	Delete(id uint) error
}
