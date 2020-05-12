package sqlstore

import (
	"fmt"
)

func deleteData() error {
	tables := []string{
		WorkingHourTable,
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

func DeleteTestData() error {
	if err := deleteData(); err != nil {
		return err
	}
	return nil
}
