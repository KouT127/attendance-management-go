package main

import (
	"github.com/KouT127/Attendance-management/backend/config"
	"github.com/KouT127/Attendance-management/backend/database"
	"github.com/KouT127/Attendance-management/backend/routes"
)

func main() {
	configs.Init(configs.Development)
	c := configs.NewConfig()
	database.Init(c)
	routes.Init()
}
