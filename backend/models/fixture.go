package models

import (
	"github.com/go-testfixtures/testfixtures/v3"
	"github.com/go-xorm/xorm"
)

var fixtures *testfixtures.Loader

func InitFixtures(eng *xorm.Engine) (err error) {
	fixtures, err = testfixtures.New(
		testfixtures.Database(eng.DB().DB),
		testfixtures.Dialect("mysql"),
		testfixtures.Files(
			"fixtures/users.yml",
			"fixtures/attendance.yml",
		),
	)
	return err
}

func PrepareTestDatabase() error {
	if err := fixtures.Load(); err != nil {
		return err
	}
	return nil
}
