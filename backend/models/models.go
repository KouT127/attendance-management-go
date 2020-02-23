package models

import (
	"database/sql"
	"github.com/KouT127/attendance-management/database"
	"github.com/go-xorm/xorm"
)

var (
	engine *xorm.Engine
)

func SetDatabase() {
	engine = database.NewDB()
}

type Engine interface {
	Table(tableNameOrBean interface{}) *xorm.Session
	UseBool(columns ...string) *xorm.Session
	Count(...interface{}) (int64, error)
	Decr(column string, arg ...interface{}) *xorm.Session
	Delete(interface{}) (int64, error)
	Exec(...interface{}) (sql.Result, error)
	Find(interface{}, ...interface{}) error
	Get(interface{}) (bool, error)
	ID(interface{}) *xorm.Session
	Select(string) *xorm.Session
	In(string, ...interface{}) *xorm.Session
	Incr(column string, arg ...interface{}) *xorm.Session
	Insert(...interface{}) (int64, error)
	InsertOne(interface{}) (int64, error)
	Iterate(interface{}, xorm.IterFunc) error
	Join(joinOperator string, tablename interface{}, condition string, args ...interface{}) *xorm.Session
	SQL(interface{}, ...interface{}) *xorm.Session
	Where(interface{}, ...interface{}) *xorm.Session
	Asc(colNames ...string) *xorm.Session
	Limit(limit int, start ...int) *xorm.Session
	SumInt(bean interface{}, columnName string) (res int64, err error)
}
