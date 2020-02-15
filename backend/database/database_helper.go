package database

import (
	"database/sql"
	"fmt"
	"github.com/KouT127/attendance-management/utils/directory"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/joho/godotenv"
	"github.com/volatiletech/sqlboiler/boil"
	"os"
)

func ConnectDatabase() {
	conn := loadTestEnv()

	db, err = sql.Open("mysql", conn)
	if err != nil {
		panic(err)
	}
	boil.DebugMode = true
}

func getMigrationsPath() string {
	return "file://" + directory.RootDir() + "/database/migrations"
}

func loadTestEnv() string {
	r := directory.RootDir()
	err := godotenv.Load(fmt.Sprintf(r + "/.env.test"))
	if err != nil {
		panic(err)
	}
	CONNECTION = os.Getenv("DB_CONNECTION_NAME")
	USER = os.Getenv("DB_USER")
	PASS = os.Getenv("DB_PASSWORD")
	TABLE = os.Getenv("DB_TABLE")
	return fmt.Sprintf("%s:%s@%s/%s?charset=utf8&parseTime=true", USER, PASS, CONNECTION, TABLE)
}

func migrateUp() error {
	p := getMigrationsPath()
	conn := loadTestEnv()
	m, err := migrate.New(p, "mysql://"+conn)
	if err != nil {
		return err
	}
	err = m.Up()
	return nil
}

func dropTable() error {
	tables := []string{
		attendanceTimeTable,
		attendanceTable,
		userTable,
		"schema_migrations",
	}
	for _, table := range tables {

		_, err := db.Exec("drop database ?", table)
		if err != nil {
			panic(err)
		}
	}
	return nil
}

func CreateTestTable() {
	if err := migrateUp(); err != nil {
		panic(err)
	}
}

func DropTestTable() {
	if err := dropTable(); err != nil {
		panic(err)
	}
}

func PrepareTestDatabase() func() {
	CreateTestTable()
	return func() {
		DropTestTable()
	}
}
