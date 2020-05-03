package sqlstore

import (
	"context"
	"errors"
	"golang.org/x/xerrors"
	"xorm.io/xorm"
)

type Transaction interface {
	InTransaction(ctx context.Context, fn func(ctx context.Context) (interface{}, error)) (interface{}, error)
	Begin(ctx context.Context) (context.Context, error)
	Commit(ctx context.Context) error
	Close(ctx context.Context)
}

type ContextSessionKey struct{}

type DBSession struct {
	*xorm.Session
}

type dbTransactionFunc func(sess *DBSession) (interface{}, error)

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

func inTransactionCtx(ctx context.Context, eng *xorm.Engine, callback dbTransactionFunc) (interface{}, error) {
	var v interface{}
	sess, err := startSession(ctx, eng, true)
	if err != nil {
		return nil, err
	}

	defer sess.Close()

	v, err = callback(sess)

	if err != nil {
		if rollErr := sess.Rollback(); rollErr != nil {
			return nil, xerrors.Errorf("Rolling back transaction due to error failed: %s", rollErr)
		}
		return nil, err
	}
	if err := sess.Commit(); err != nil {
		return nil, err
	}

	return v, nil
}

func (ss *sqlStore) InTransaction(ctx context.Context, fn func(ctx context.Context) (interface{}, error)) (interface{}, error) {
	return inTransactionCtx(ctx, ss.engine, func(sess *DBSession) (interface{}, error) {
		withValue := context.WithValue(ctx, ContextSessionKey{}, sess)
		return fn(withValue)
	})
}

func getDBSession(ctx context.Context) (*DBSession, error) {
	return startSession(ctx, eng, false)
}

func (ss *sqlStore) Begin(ctx context.Context) (context.Context, error) {
	sess, err := startSession(ctx, eng, true)
	if err != nil {
		return nil, err
	}
	c := context.WithValue(ctx, ContextSessionKey{}, sess)
	return c, nil
}

func (ss *sqlStore) Commit(ctx context.Context) error {
	sess, err := startSession(ctx, eng, false)
	if err != nil {
		return err
	}
	if err := sess.Commit(); err != nil {
		return err
	}
	return nil
}

func (ss *sqlStore) Close(ctx context.Context) {
	sess, _ := startSession(ctx, eng, false)
	sess.Close()
}
