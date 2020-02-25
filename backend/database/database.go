package database

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"github.com/joho/godotenv"
	"os"
	"time"
	"xorm.io/core"
)

const (
	UserTable           = "users"
	AttendanceTable     = "attendances"
	AttendanceTimeTable = "attendances_time"
)

var (
	engine     *xorm.Engine
	err        error
	CONNECTION string
	USER       string
	PASS       string
	TABLE      string
	SOURCE     string
)

func SetUp() {
	SOURCE = loadEnv()
	if CONNECTION == "" {
		SOURCE = loadLocalEnv()
	}

	engine, err = xorm.NewEngine("mysql", SOURCE)
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
	engine.ShowExecTime(true)
}

func NewDB() *xorm.Engine {
	return engine
}

func loadEnv() string {
	CONNECTION = os.Getenv("CLOUDSQL_CONNECTION_NAME")
	USER = os.Getenv("CLOUDSQL_USER")
	PASS = os.Getenv("CLOUDSQL_PASSWORD")
	TABLE = os.Getenv("DB_TABLE")
	SOURCE = fmt.Sprintf("%s:%s@%s/%s?charset=utf8&parseTime=true", USER, PASS, CONNECTION, TABLE)
	return SOURCE
}

func loadLocalEnv() string {
	err := godotenv.Load(fmt.Sprintf("./backend/configs/.env.local"))
	if err != nil {
		panic(err)
	}
	CONNECTION = os.Getenv("DB_CONNECTION_NAME")
	USER = os.Getenv("DB_USER")
	PASS = os.Getenv("DB_PASSWORD")
	TABLE = os.Getenv("DB_TABLE")
	SOURCE = fmt.Sprintf("%s:%s@%s/%s?charset=utf8&parseTime=true", USER, PASS, CONNECTION, TABLE)
	return SOURCE
}
