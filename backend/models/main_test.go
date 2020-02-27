package models

import (
	"fmt"
	"github.com/KouT127/attendance-management/database"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	MainTest()
	code := m.Run()
	TearDown()
	os.Exit(code)
}

func fatalTestError(fmtStr string, args ...interface{}) {
	_, _ = fmt.Fprintf(os.Stderr, fmtStr, args...)
	os.Exit(1)
}

func MainTest() {
	engine := database.CreateTestEngine()
	if err := InitFixtures(engine); err != nil {
		fatalTestError("initial fixture error %s\n", err)
	}
}

func TearDown() {
	
}
