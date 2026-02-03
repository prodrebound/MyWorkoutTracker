package repository

import (
	"errors"
	"time"

	"github.com/prodrebound/MyWorkoutTracker/backend/internal/core/domain"
	"gorm.io/gorm"
)

type WorkoutSessionRepository struct {
	db *gorm.DB
}

func NewWorkoutSessionRepository(db *gorm.DB) *WorkoutSessionRepository {
	return &WorkoutSessionRepository{db: db}
}

func (r *WorkoutSessionRepository) Create(session *domain.WorkoutSession) error {
	return r.db.Create(session).Error
}

func (r *WorkoutSessionRepository) GetInTimeRange(start, end time.Time) ([]domain.WorkoutSession, error) {
	var sessions []domain.WorkoutSession
	err := r.db.
		Preload("Routine").
		Preload("Exercises.Exercise").
		Where("date >= ? AND date <= ?", start, end).
		Order("date asc").
		Find(&sessions).Error
	return sessions, err
}

func (r *WorkoutSessionRepository) GetByID(id uint) (*domain.WorkoutSession, error) {
	var session domain.WorkoutSession
	err := r.db.
		Preload("Routine").
		Preload("Exercises.Exercise").
		First(&session, id).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("session not found")
		}
		return nil, err
	}
	return &session, nil
}

func (r *WorkoutSessionRepository) Update(session *domain.WorkoutSession) error {
	return r.db.Session(&gorm.Session{FullSaveAssociations: true}).Save(session).Error
}
