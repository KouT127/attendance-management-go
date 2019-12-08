package routes

import (
	"github.com/KouT127/Attendance-management/backend/controllers"
	"github.com/KouT127/Attendance-management/backend/middlewares"
	"github.com/gin-contrib/cors"
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
		//middlewares.FetchAuthorizedUser(),
	}
	uc := controllers.UserController{}
	users := v1.Group("/users", handlers...)
	users.GET("", uc.UserListController)
	users.GET("/mine", uc.UserMineController)
	users.PUT("/:id", uc.UserUpdateController)
}

func v1Router(r *Engine) {
	v1Group := r.Group("/v1")
	v1UsersRouter(v1Group)
}

func Init() {
	r := New()
	r.Use(Logger())
	r.Use(Recovery())
	config := cors.DefaultConfig()
	config.AllowMethods = []string{"OPTION", "GET", "POST", "PUT", "DELETE"}
	config.AllowOrigins = []string{"*"}
	config.AllowHeaders = []string{"*"}
	r.Use(cors.New(config))

	r.StaticFS("/static", http.Dir("frontend/build/static"))

	v1Router(r)
	defaultRouter(r)
	panic(r.Run(":8080"))
}
