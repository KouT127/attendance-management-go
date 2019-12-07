package database

import (
	"github.com/KouT127/Attendance-management/backend/config"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
)

var (
	engine *xorm.Engine
	err    error
)

func Init(c *config.Config) {
	USER := c.Database.User
	PASS := c.Database.Pass
	PROTOCOL := "tcp(" + c.Database.Host + ":" + c.Database.Port + ")"
	DBNAME := c.Database.DbName
	OPTION := c.Database.Option

	CONNECT := USER + ":" + PASS + "@" + PROTOCOL + "/" + DBNAME + "?" + OPTION
	engine, err = xorm.NewEngine("mysql", CONNECT)
	if err != nil {
		panic(err)
	}
	engine.Logger()
}

func NewDB() *xorm.Engine {
	return engine
}
