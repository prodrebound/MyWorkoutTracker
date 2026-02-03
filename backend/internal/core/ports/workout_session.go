package ports

import (
	"time"

	"github.com/prodrebound/MyWorkoutTracker/backend/internal/core/domain"
)

type WorkoutSessionRepository interface {
	Create(session *domain.WorkoutSession) error
	GetInTimeRange(start, end time.Time) ([]domain.WorkoutSession, error)
	GetByID(id uint) (*domain.WorkoutSession, error)
	Update(session *domain.WorkoutSession) error
}
