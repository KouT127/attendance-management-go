package main

import (
	"github.com/KouT127/attendance-management/backend/configs"
	"github.com/KouT127/attendance-management/backend/database"
	"github.com/KouT127/attendance-management/backend/routes"
	"github.com/KouT127/attendance-management/backend/utils/timezone"
)

func main() {
	configs.Init(configs.Development)
	c := configs.NewConfig()
	timezone.Init("Asia/Tokyo")
	database.Init(c)
	routes.Init()
}
