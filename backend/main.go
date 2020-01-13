package main

import (
	"github.com/KouT127/attendance-management/configs"
	"github.com/KouT127/attendance-management/database"
	"github.com/KouT127/attendance-management/routes"
	"github.com/KouT127/attendance-management/utils/timezone"
)

func main() {
	configs.Init(configs.Development)
	c := configs.NewConfig()
	timezone.Init("Asia/Tokyo")
	database.Init(c)
	routes.Init()
}
