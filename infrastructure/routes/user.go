package routes

import (
	"github.com/KouT127/attendance-management/api/handler/middlewares"
	"github.com/KouT127/attendance-management/api/handler/v1/user"
	"github.com/KouT127/attendance-management/application/services"
	"github.com/KouT127/attendance-management/infrastructure/sqlstore"
	"github.com/gin-gonic/gin"
)

func configureUsersRouter(v1 *gin.RouterGroup, store sqlstore.SqlStore) {
	userService := services.NewUserService(store)
	handler := user.NewUserHandler(userService)

	funcs := []gin.HandlerFunc{
		middlewares.AuthRequired(),
	}

	users := v1.Group("/users", funcs...)
	users.GET("/mine", handler.MineHandler)
	users.PUT("/:id", handler.UpdateHandler)
}
