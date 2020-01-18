package database

import (
	"fmt"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/joho/godotenv"
	"os"
	"path"
	"path/filepath"
	"runtime"
)

func rootDir() string {
	_, b, _, _ := runtime.Caller(0)
	d := path.Join(path.Dir(b))
	return filepath.Dir(d)
}

func loadTestEnv() string {
	r := rootDir()
	err := godotenv.Load(fmt.Sprintf(r + "/.env.test"))
	if err != nil {
		panic(err)
	}
	CONNECTION = os.Getenv("DB_CONNECTION_NAME")
	USER = os.Getenv("DB_USER")
	PASS = os.Getenv("PASSWORD")
	return fmt.Sprintf("%s:%s@%s/attendance_management?charset=utf8&parseTime=true", USER, PASS, CONNECTION)
}

func migrateUp() error {
	p := "file://" + rootDir() + "/database/migrations"
	conn := loadTestEnv()
	m, err := migrate.New(p, "mysql://"+conn)
	if err != nil {
		return err
	}
	err = m.Up()
	return nil
}

func migrateDown() error {
	p := "file://" + rootDir() + "/database/migrations"
	conn := loadTestEnv()
	m, err := migrate.New(p, "mysql://"+conn)
	if err != nil {
		return err
	}
	err = m.Down()
	return nil
}

func CreateTestTable() {
	if err := migrateUp(); err != nil {
		panic(err)
	}
	print("migrate \n")
}

func DropTestTable() {
	if err := migrateDown(); err != nil {
		panic(err)
	}
	print("drop \n")
}
