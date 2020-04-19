package sqlstore

import (
	"context"
	"errors"
	"golang.org/x/xerrors"
	"xorm.io/xorm"
)

type ContextSessionKey struct{}

type DBSession struct {
	*xorm.Session
}

type dbTransactionFunc func(sess *DBSession) error

func startSession(ctx context.Context, engine *xorm.Engine, beginTran bool) (*DBSession, error) {
	value := ctx.Value(ContextSessionKey{})
	var sess *DBSession
	sess, ok := value.(*DBSession)

	if ok {
		return sess, nil
	}

	if engine == nil {
		return nil, errors.New("xorm.eng is nil")
	}

	newSess := &DBSession{Session: engine.NewSession()}
	if beginTran {
		err := newSess.Begin()
		if err != nil {
			return nil, err
		}
	}
	return newSess, nil
}

func inTransactionCtx(ctx context.Context, eng *xorm.Engine, callback dbTransactionFunc) error {
	sess, err := startSession(ctx, eng, true)
	if err != nil {
		return err
	}

	defer sess.Close()

	err = callback(sess)

	if err != nil {
		if rollErr := sess.Rollback(); rollErr != nil {
			return xerrors.Errorf("Rolling back transaction due to error failed: %s", rollErr)
		}
		return err
	}
	if err := sess.Commit(); err != nil {
		return err
	}

	return nil
}

func (ss *SQLStore) InTransaction(ctx context.Context, fn func(ctx context.Context) error) error {
	return inTransactionCtx(ctx, ss.engine, func(sess *DBSession) error {
		withValue := context.WithValue(ctx, ContextSessionKey{}, sess)
		return fn(withValue)
	})
}

func withDBSession(ctx context.Context, callback dbTransactionFunc) error {
	sess, err := startSession(ctx, eng, false)
	if err != nil {
		return err
	}

	return callback(sess)
}
