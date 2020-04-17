package database

import (
	"fmt"
	"github.com/KouT127/attendance-management/modules/directory"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"os"
	"xorm.io/xorm"
)

func getMigrationsPath() string {
	return "file://" + directory.RootDir() + "/database/migrations"
}

func loadTestEnv() string {
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

func deleteData() error {
	tables := []string{
		AttendanceTimeTable,
		AttendanceTable,
		UserTable,
		"schema_migrations",
	}
	for _, table := range tables {
		sql := fmt.Sprintf("delete from %s", table)
		if _, err := engine.Exec(sql); err != nil {
			panic(err)
		}
	}
	return nil
}

func CreateTestTable() error {
	if err := migrateUp(); err != nil {
		return err
	}
	return nil
}

func DeleteTestData() error {
	if err := deleteData(); err != nil {
		return err
	}
	return nil
}

func InitTestConnection() error {
	var (
		err       error
		dbUser    = mustGetenv("DB_USER")
		dbPwd     = mustGetenv("DB_PASS")
		dbTCPHost = mustGetenv("DB_TCP_HOST")
		dbName    = mustGetenv("DB_NAME")
	)

	uri := fmt.Sprintf("%s:%s@tcp(%s)/%s", dbUser, dbPwd, dbTCPHost, dbName)
	engine, err = xorm.NewEngine("mysql", uri)
	if err != nil {
		return fmt.Errorf("xorm.NewEngine: %v", err)
	}

	// configure settings
	configureConnectionPool(engine)
	configureLogger(engine)
	configureTimezone(engine)
	return nil
}
