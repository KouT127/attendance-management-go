package routes

import (
	"github.com/KouT127/attendance-management/api/handler/middlewares"
	"github.com/KouT127/attendance-management/api/handler/v1/attendance"
	"github.com/KouT127/attendance-management/application/services"
	"github.com/KouT127/attendance-management/infrastructure/sqlstore"
	"github.com/gin-gonic/gin"
)

func configureAttendancesRouter(v1 *gin.RouterGroup, store sqlstore.SQLStore) {
	attendanceService := services.NewAttendanceService(store)
	handler := attendance.NewAttendanceHandler(attendanceService)

	funcs := []gin.HandlerFunc{
		middlewares.AuthRequired(),
	}

	attendances := v1.Group("/attendances", funcs...)
	attendances.GET("", handler.ListHandler)
	attendances.POST("", handler.CreateHandler)
}
