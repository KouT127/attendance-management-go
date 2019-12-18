package main

import (
	"github.com/KouT127/attendance-management/backend/configs"
	"github.com/KouT127/attendance-management/backend/database"
	"github.com/KouT127/attendance-management/backend/routes"
)

func main() {
	configs.Init(configs.Development)
	c := configs.NewConfig()
	database.Init(c)
	routes.Init()
}
