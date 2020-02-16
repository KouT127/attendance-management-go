package repositories

import (
	"context"
	"github.com/go-xorm/xorm"
)

const (
	engineKey = "ENGINEKEY"
	txKey     = "TXKEY"
)

func NewTransaction() *transaction {
	return &transaction{}
}

type Transaction interface {
	NewSession(engine *xorm.Engine) *xorm.Session
	Begin(session *xorm.Session) error
	Commit(session *xorm.Session) error
	Close(session *xorm.Session)
}

type transaction struct{}

func (r *transaction) NewSession(engine *xorm.Engine) *xorm.Session {
	return engine.NewSession()
}
func (r *transaction) Begin(session *xorm.Session) error {
	return session.Begin()
}
func (r *transaction) Commit(session *xorm.Session) error {
	return session.Commit()
}
func (r *transaction) Close(session *xorm.Session) {
	session.Close()
}

// TODO: use context
func GetEngine(ctx context.Context) (*xorm.Engine, bool) {
	engine, ok := ctx.Value(engineKey).(*xorm.Engine)
	return engine, ok
}

// TODO: use context
func GetTx(ctx context.Context) (*xorm.Session, bool) {
	tx, ok := ctx.Value(txKey).(*xorm.Session)
	return tx, ok
}
