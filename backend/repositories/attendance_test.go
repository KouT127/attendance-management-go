package repositories

import (
	"github.com/KouT127/attendance-management/database"
	"github.com/stretchr/testify/assert"
	"testing"
)

func prepareTestDatabase() func() {
	print("preparing \n")
	database.CreateTestTable()
	return func() {
		print("teardown\n")
		database.DropTestTable()
	}
}
func PreparingTest() {
	tearDown := prepareTestDatabase()
	defer tearDown()
}

func TestAttendanceRepository_CreateAttendance(t *testing.T) {
	PreparingTest()
	assert.Equal(t, "", "")
}
