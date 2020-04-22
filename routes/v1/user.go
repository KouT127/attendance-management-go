package v1

import (
	"github.com/KouT127/attendance-management/application/facades"
	"github.com/KouT127/attendance-management/handlers/middlewares"
	"github.com/KouT127/attendance-management/handlers/v1/user"
	"github.com/KouT127/attendance-management/infrastructure/sqlstore"
	. "github.com/gin-gonic/gin"
)

func UsersRouter(v1 *RouterGroup) {
	funcs := []HandlerFunc{
		middlewares.AuthRequired(),
	}

	store := sqlstore.InitDatabase()
	facade := facades.NewUserFacade(&store)
	handler := user.NewUserHandler(facade)

	users := v1.Group("/users", funcs...)
	users.GET("/mine", handler.MineHandler)
	users.PUT("/:id", handler.UpdateHandler)
}
