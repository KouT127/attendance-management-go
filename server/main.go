package main

import (
	"github.com/KouT127/attendance-management/database"
	"github.com/KouT127/attendance-management/models"
	"github.com/KouT127/attendance-management/modules/logger"
	"github.com/KouT127/attendance-management/modules/timezone"
	"github.com/KouT127/attendance-management/routes"
)

func main() {
	logger.SetUp()
	timezone.Set("Asia/Tokyo")
	database.InitDatabase()
	models.SetDatabase()
	routes.Init()
}
