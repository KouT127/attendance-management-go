package sqlstore

import (
	"context"
	"github.com/KouT127/attendance-management/domain/models"
	"time"
)

type WorkingHour interface {
	GetWorkingHours(ctx context.Context, now time.Time) (*models.WorkingHour, error)
	CreateWorkingHour(ctx context.Context, hour *models.WorkingHour) error
}

func (sqlStore) GetWorkingHours(ctx context.Context, now time.Time) (*models.WorkingHour, error) {
	sess, err := getDBSession(ctx)
	if err != nil {
		return nil, err
	}

	wh := &models.WorkingHour{}
	has, err := sess.
		Where("started_at < ?", now).
		And("finished_at > ?", now).
		Get(wh)
	if err != nil {
		return nil, err
	}
	if !has {
		return nil, nil
	}

	return wh, nil
}

func (sqlStore) CreateWorkingHour(ctx context.Context, hour *models.WorkingHour) error {
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
