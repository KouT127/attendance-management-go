package main

import (
	"github.com/KouT127/attendance-management/infrastructure/routes"
	"github.com/KouT127/attendance-management/infrastructure/sqlstore"
	"github.com/KouT127/attendance-management/utilities/logger"
	"github.com/KouT127/attendance-management/utilities/timezone"
)

func main() {
	store := sqlstore.InitDatabase()
	logger.SetUp()
	timezone.Set("Asia/Tokyo")
	routes.InitRouter(&store)
}
