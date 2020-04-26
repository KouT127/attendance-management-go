package main

import (
	"github.com/KouT127/attendance-management/infrastructure/routes"
	"github.com/KouT127/attendance-management/modules/logger"
	"github.com/KouT127/attendance-management/modules/timezone"
)

func main() {
	logger.SetUp()
	timezone.Set("Asia/Tokyo")
	routes.Init()
}
