package sqlstore

import (
	"context"
	"github.com/KouT127/attendance-management/domain/models"
	"time"
)

type WorkingHour interface {
	GetWorkingHours(ctx context.Context, now time.Time) (*models.WorkingHour, error)
	CreateWorkingHours(ctx context.Context, hour *WorkingHour) error
}

func (sqlStore) GetWorkingHours(ctx context.Context, now time.Time) (*models.WorkingHour, error) {
	sess, err := getDBSession(ctx)
	if err != nil {
		return nil, err
	}

	WorkingHours := &models.WorkingHour{}
	_, err = sess.
		Where("started_at < ?", now).
		And("finished_at > ?", now).
		Get(WorkingHours)
	if err != nil {
		return nil, err
	}

	return WorkingHours, nil
}

func (sqlStore) CreateWorkingHours(ctx context.Context, hour *WorkingHour) error {
	sess, err := getDBSession(ctx)
	if err != nil {
		return err
	}
	_, err = sess.
		Insert(hour)

	if err != nil {
		return err
	}

	return nil
}
