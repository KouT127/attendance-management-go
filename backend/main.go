package main

import (
	"github.com/KouT127/attendance-management/database"
	"github.com/KouT127/attendance-management/models"
	"github.com/KouT127/attendance-management/routes"
	"github.com/KouT127/attendance-management/utils/logger"
	"github.com/KouT127/attendance-management/utils/timezone"
)

func main() {
	logger.Init()
	timezone.Init("Asia/Tokyo")
	database.Init()
	models.SetDatabase()
	routes.Init()
}
