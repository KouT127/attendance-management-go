package v1

import (
	"github.com/KouT127/attendance-management/handlers/v1/user"
	"github.com/KouT127/attendance-management/middlewares"
	. "github.com/gin-gonic/gin"
)

func UsersRouter(v1 *RouterGroup) {
	funcs := []HandlerFunc{
		middlewares.AuthRequired(),
	}

	users := v1.Group("/users", funcs...)
	users.GET("/mine", user.V1MineHandler)
	users.PUT("/:id", user.V1UpdateHandler)
}
