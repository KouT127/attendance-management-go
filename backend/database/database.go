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

var (
	engine        *xorm.Engine
	err           error
	CONNECTION    string
	USER          string
	PASS          string
	DB_CONNECTION string
)

func Init() {
	DB_CONNECTION = loadEnv()
	if CONNECTION == "" {
		DB_CONNECTION = loadLocalEnv()
	}

	engine, err = xorm.NewEngine("mysql", DB_CONNECTION)
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

func loadEnv() string {
	CONNECTION := os.Getenv("CLOUDSQL_CONNECTION_NAME")
	USER := os.Getenv("CLOUDSQL_USER")
	PASS := os.Getenv("CLOUDSQL_PASSWORD")
	DB_CONNECTION = fmt.Sprintf("%s:%s@%s/attendance_management?charset=utf8&parseTime=true", USER, PASS, CONNECTION)
	return DB_CONNECTION
}

func loadLocalEnv() string {
	err := godotenv.Load(fmt.Sprintf("./backend/.env.local"))
	if err != nil {
		panic(err)
	}
	CONNECTION := os.Getenv("DB_CONNECTION_NAME")
	USER := os.Getenv("DB_USER")
	PASS := os.Getenv("PASSWORD")
	DB_CONNECTION = fmt.Sprintf("%s:%s@%s/attendance_management?charset=utf8&parseTime=true", USER, PASS, CONNECTION)
	return DB_CONNECTION
}
