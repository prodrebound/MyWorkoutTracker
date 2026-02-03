package service

import (
	"errors"
	"time"

	"github.com/prodrebound/MyWorkoutTracker/backend/internal/core/domain"
	"github.com/prodrebound/MyWorkoutTracker/backend/internal/core/ports"
)

type WorkoutSessionService struct {
	repo        ports.WorkoutSessionRepository
	routineRepo ports.RoutineRepository
}

func NewWorkoutSessionService(repo ports.WorkoutSessionRepository, routineRepo ports.RoutineRepository) *WorkoutSessionService {
	return &WorkoutSessionService{repo: repo, routineRepo: routineRepo}
}

func (s *WorkoutSessionService) ScheduleSession(date time.Time, routineID uint) (*domain.WorkoutSession, error) {
	routine, err := s.routineRepo.GetByID(routineID)
	if err != nil {
		return nil, errors.New("routine not found")
	}
	session := &domain.WorkoutSession{
		Date:      date,
		RoutineID: routineID,
		Completed: false,
		Exercises: []domain.SessionExercise{},
	}
	for _, routineEx := range routine.Exercises {
		session.Exercises = append(session.Exercises, domain.SessionExercise{
			ExerciseID: routineEx.ExerciseID,
			Sets:       routineEx.Sets,
			Reps:       routineEx.Reps,
			Weight:     0,
			Order:      routineEx.Order,
		})
	}
	return session, s.repo.Create(session)
}

func (s *WorkoutSessionService) GetHistory(start, end time.Time) ([]domain.WorkoutSession, error) {
	if start.After(end) {
		return nil, errors.New("start date cannot be after end date")
	}

	return s.repo.GetInTimeRange(start, end)
}
