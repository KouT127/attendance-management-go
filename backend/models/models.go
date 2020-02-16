package models

import (
	"github.com/KouT127/attendance-management/database"
	"github.com/go-xorm/xorm"
)

var (
	engine *xorm.Engine
)

func Init() {
	engine = database.NewDB()
}
