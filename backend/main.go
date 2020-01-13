package main

import (
	"github.com/KouT127/attendance-management/database"
	"github.com/KouT127/attendance-management/routes"
	"github.com/KouT127/attendance-management/utils/timezone"
)

func main() {
	timezone.Init("Asia/Tokyo")
	database.Init()
	routes.Init()
}
