package main

import (
	"github.com/KouT127/attendance-management/infrastructure/routes"
	"github.com/KouT127/attendance-management/infrastructure/sqlstore"
	"github.com/KouT127/attendance-management/infrastructure/uploader"
	"github.com/KouT127/attendance-management/utilities/logger"
	"github.com/KouT127/attendance-management/utilities/timezone"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

func main() {
	logger.SetUp()
	timezone.Set("Asia/Tokyo")
	upl, err := uploader.NewStorageUploader()
	if err != nil {
		log.Fatalf("%v", err)
	}
	routes.InitRouter(sqlstore.InitDatabase(), upl)
}
