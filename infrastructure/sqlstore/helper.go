package sqlstore

import (
	"fmt"
	"github.com/KouT127/attendance-management/modules/directory"
	_ "github.com/go-sql-driver/mysql"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"xorm.io/xorm"
)

func getMigrationsPath() string {
	return "file://" + directory.RootDir() + "/sqlstore/migrations"
}

func loadTestEnv() string {
	var (
		dbUser    = mustGetenv("DB_USER")
		dbPwd     = mustGetenv("DB_PASS")
		dbTCPHost = mustGetenv("DB_TCP_HOST")
		dbName    = mustGetenv("TEST_DB_NAME")
	)

	uri := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=true", dbUser, dbPwd, dbTCPHost, dbName)
	return uri
}

func migrateUp() error {
	p := getMigrationsPath()
	conn := loadTestEnv()
	m, err := migrate.New(p, "mysql://"+conn)
	if err != nil {
		return err
	}
	err = m.Up()
	if err != nil {
		return err
	}
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
		if _, err := eng.Exec(sql); err != nil {
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
	eng, err = xorm.NewEngine("mysql", uri)
	if err != nil {
		return fmt.Errorf("xorm.NewEngine: %v", err)
	}

	// configure settings
	configureConnectionPool(eng)
	configureLogger(eng)
	configureTimezone(eng)
	return nil
}
