package database

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	"github.com/volatiletech/sqlboiler/boil"
	"os"
)

const (
	userTable           = "users"
	attendanceTable     = "attendances"
	attendanceTimeTable = "attendances_time"
)

var (
	db         *sql.DB
	err        error
	CONNECTION string
	USER       string
	PASS       string
	TABLE      string
	SOURCE     string
)

func Init() {
	SOURCE = loadEnv()
	if CONNECTION == "" {
		SOURCE = loadLocalEnv()
		boil.DebugMode = true
	}
	db, err = sql.Open("mysql", SOURCE)

	if err != nil {
		panic(err)
	}
}

func NewDB() *sql.DB {
	return db
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
	err := godotenv.Load(fmt.Sprintf("./backend/.env.local"))
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
