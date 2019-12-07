package routes

import (
	"github.com/KouT127/Attendance-management/backend/middlewares"
	. "github.com/gin-gonic/gin"
	"net/http"
)

func renderIndex(c *Context) {
	c.HTML(http.StatusOK, "index.html", H{})
}

func defaultRouter(r *Engine) {
	r.LoadHTMLFiles("frontend/build/index.html")
	r.NoRoute(renderIndex)
}

func v1UsersRouter(v1 *RouterGroup) {
	handlers := []HandlerFunc{
		middlewares.AuthRequired(),
		middlewares.FetchAuthorizedUser(),
	}

	users := v1.Group("/users", handlers...)
	users.GET("", func(ctx *Context) {
		ctx.JSON(200, H{
			"message": "hello",
		})
	})
}

func v1Router(r *Engine) {
	v1Group := r.Group("/v1")
	v1UsersRouter(v1Group)
}

func Init() {
	r := New()
	r.Use(Logger())
	r.Use(Recovery())

	r.StaticFS("/static", http.Dir("frontend/build/static"))

	v1Router(r)
	defaultRouter(r)
	panic(r.Run(":9000"))
}
