package database

import (
	"github.com/KouT127/attendance-management/configs"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"os"
	"time"
	"xorm.io/core"
)

var (
	engine *xorm.Engine
	err    error
)

func Init(c *configs.Config) {
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
	logger := xorm.NewSimpleLogger(os.Stdout)
	logger.ShowSQL(true)
	logger.SetLevel(core.LOG_INFO)
	loc, err := time.LoadLocation("UTC")
	if err != nil {
		panic(err)
	}
	engine.SetTZLocation(loc)
	engine.SetTZDatabase(loc)
	engine.SetLogger(logger)
}

func NewDB() *xorm.Engine {
	return engine
}
