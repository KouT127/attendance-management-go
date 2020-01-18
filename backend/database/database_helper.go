package database

import (
	"fmt"
	"github.com/KouT127/attendance-management/utils/directory"
	"github.com/go-xorm/xorm"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/joho/godotenv"
	"os"
	"time"
	"xorm.io/core"
)

func ConnectDatabase() {
	conn := loadTestEnv()
	engine, err = xorm.NewEngine("mysql", conn)
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
		"attendances_time",
		"attendances",
		"users",
		"schema_migrations",
	}
	for _, table := range tables {
		if err := engine.DropTables(table); err != nil {
			panic(err)
		}
	}
	return nil
}

func CreateTestTable() {
	if err := migrateUp(); err != nil {
		panic(err)
	}
	print("migrate \n")
}

func DropTestTable() {
	if err := dropTable(); err != nil {
		panic(err)
	}
	print("drop \n")
}

func PrepareTestDatabase() func() {
	print("preparing \n")
	CreateTestTable()
	return func() {
		print("teardown\n")
		DropTestTable()
	}
}

func SetupTest() {

}
