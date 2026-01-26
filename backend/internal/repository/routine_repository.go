package repository

import (
	"errors"

	"github.com/prodrebound/MyWorkoutTracker/backend/internal/core/domain"
	"gorm.io/gorm"
)

type RoutineRepository struct {
	db *gorm.DB
}

func NewRoutineRepository(db *gorm.DB) *RoutineRepository {
	return &RoutineRepository{db: db}
}

func (r *RoutineRepository) Create(routine *domain.Routine) error {
	return r.db.Create(routine).Error
}

func (r *RoutineRepository) GetAll() ([]domain.Routine, error) {
	var routines []domain.Routine
	err := r.db.Preload("Exercises.Exercise").Find(&routines).Error
	return routines, err
}

func (r *RoutineRepository) GetByID(id uint) (*domain.Routine, error) {
	var routine domain.Routine
	err := r.db.Preload("Exercises.Exercise").First(&routine, id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("routine not found")
		}
		return nil, err
	}
	return &routine, nil
}

func (r *RoutineRepository) Update(routine *domain.Routine) error {
	return r.db.Transaction(func(tx *gorm.DB) error {

		if err := tx.Model(routine).Omit("Exercises").Updates(routine).Error; err != nil {
			return err
		}

		if err := tx.Where("routine_id = ?", routine.ID).Delete(&domain.RoutineExercise{}).Error; err != nil {
			return err
		}

		if len(routine.Exercises) > 0 {
			for i := range routine.Exercises {
				routine.Exercises[i].RoutineID = routine.ID
			}

			if err := tx.Create(&routine.Exercises).Error; err != nil {
				return err
			}
		}

		return nil
	})
}

func (r *RoutineRepository) Delete(id uint) error {
	return r.db.Delete(&domain.Routine{}, id).Error
}
